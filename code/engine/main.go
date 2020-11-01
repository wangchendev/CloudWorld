package main

import (
    "engine/config"
    "errors"
    "os"
    "os/signal"
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
    config.Init()

    gateApp := gate.NewGate()
    gateApp.Init()

    // errCh可以接收来自一个关键子协程的错误，例如Server loop
    errCh := make(chan error, 1)
    if err := waitSignal(errCh); err != nil {
        logs.Error("%s", err)
        return
    }
}