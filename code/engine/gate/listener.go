package gate

import (
    "net"
    "os"
    "sync/atomic"

    "engine/logs"
)

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
}

func newTcpListener(address string) *TcpListener {
    return &TcpListener{
        stopCh: make(chan bool, 1),
        addr: address,
        closeFlag: 0,
    }
}

func (l *TcpListener) Init() {
    socket, err := net.Listen("tcp", l.addr)
    if err != nil {
        logs.Error("tcp listen: %s", err)
        logs.Flush()
        os.Exit(1)
    }

    l.Listener = socket
}

func (l *TcpListener) Start() {
    go func() {
        l.acceptAndServe()
    }()
}

func (l *TcpListener) Stop() {
    atomic.StoreInt32(&l.closeFlag, 1)
    if err := l.Close(); err != nil {
        logs.Error("tcp listen close: %s", err)
    }
    <- l.stopCh
}

func (l *TcpListener) acceptAndServe() {
    defer func() {
        if atomic.LoadInt32(&l.closeFlag) == 0 {
            if err := l.Close(); err != nil {
                logs.Error("tcp listen close: %s", err)
            }
        }
        l.stopCh <- true
    }()

    logs.Info("tcp listener started")

    for {
        conn, err := l.Accept()
        if err != nil {
            if atomic.LoadInt32(&l.closeFlag) == 0 {
                logs.Error("tcp listener accept: %s", err)
                continue
            } else {
                break
            }
        }
        go func() {
            logs.Info("session start %v", conn)
            //sess := newTcpSession(conn)
            //sess.Start()
        }()
    }
}
