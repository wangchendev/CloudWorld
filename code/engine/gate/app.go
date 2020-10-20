package gate

import "engine/app"

type Gate struct {
    app.App
}

func (g *Gate) BootPrepare() {
    g.RegisterComponent(SessionMgr{})
}