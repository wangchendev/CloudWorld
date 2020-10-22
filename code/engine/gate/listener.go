package gate

import "net"

type IListener interface {
    Init()
    Start()
    Stop()
}

type TcpListener struct {
    net.Listener
    closeFlag int32
    stopCh chan bool
    addr string
    needAuth bool
}

func newTcpListener(address string, needAuth bool) *TcpListener {
    return &TcpListener{
        stopCh: make(chan bool, 1),
        addr: address,
        closeFlag: 0,
        needAuth: needAuth,
    }
}
