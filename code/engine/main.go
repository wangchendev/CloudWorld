package main

import (
    "errors"
    "fmt"
    "net"
    "os"
    "os/signal"
    "strings"
    "sync"
    "syscall"

    "engine/gate"
    "engine/logs"
)

func waitSignal(errCh chan error) error {
    signals := make(chan os.Signal, 1)
    signal.Notify(signals, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM)

    defer logs.Flush()

    for {
        select {
        case sig := <-signals:
            switch sig {
            case syscall.SIGTERM:
                return errors.New(sig.String()) // force exit
            case syscall.SIGHUP, syscall.SIGINT:
                return nil // graceful shutdown
            }
        case err := <-errCh:
            return err
        }
    }
}

func main() {
    gateApp := gate.Gate{}
    gateApp.Init()
    gateApp.BootPrepare()

    // errCh可以接收来自一个关键子协程的错误，例如Server loop
    errCh := make(chan error, 1)
    if err := waitSignal(errCh); err != nil {
        logs.Error("%s", err)
        return
    }
}


func doGetIntranetIP() (string, error) {
    ifaces, err := net.Interfaces()
    if err != nil {
        return "", err
    }

    for _, iface := range ifaces {
        if iface.Flags & net.FlagUp == 0 {
            continue
        }

        if iface.Flags & net.FlagLoopback != 0 {
            continue
        }

        addrs, err := iface.Addrs()
        if err != nil {
            return "", err
        }

        for _, addr := range addrs {
            var ip net.IP

            switch v := addr.(type) {
            case *net.IPNet:
                ip = v.IP
            case *net.IPAddr:
                ip = v.IP
            }

            if ip == nil || ip.IsLoopback() {
                continue
            }

            ip = ip.To4()
            if ip == nil {
                continue
            }

            //只取内网IP
            //10.0.0.0/8 172.16.0.0/12 192.168.0.0/16
            sip := ip.String()
            if strings.HasPrefix(sip, "10.") || strings.HasPrefix(sip, "192.168.") {
                return sip, nil
            }
            if strings.HasPrefix(sip, "172.") {
                remain := sip[4:]
                for i := 16; i < 32; i++ {
                    prx := fmt.Sprintf("%d.", i)
                    if strings.HasPrefix(remain, prx) {
                        return sip, nil
                    }
                }
            }
        }
    }
    return "", errors.New("no intranet ip found")
}

type intranetIpCache struct {
    lock sync.RWMutex
    ip string
}

func (c *intranetIpCache) getIntranetIP() (string, error) {
    c.lock.RLock()
    ret := c.ip
    c.lock.RUnlock()
    if len(ret) > 0 {
        return ret, nil
    }

    c.lock.Lock()

    var err error
    if len(c.ip) == 0 {
        c.ip, err = doGetIntranetIP()
    }

    c.lock.Unlock()

    return c.ip, err
}

var gIntranetIpCache intranetIpCache

func GetIntranetIP() (string, error) {
    return gIntranetIpCache.getIntranetIP()
}
