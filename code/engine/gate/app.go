package gate

import (
    "engine/app"
)

type Gate struct {
    *app.App
}

func NewGate() *Gate {
    gate := Gate{
        App: app.NewApp(),
    }
    gate.App.BootUnit.IBootUnit = &gate
    return &gate
}

func (g *Gate) BootPrepare() {
    g.RegisterComponent(NewSessionMgr())
}