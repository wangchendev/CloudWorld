package gate

import "engine/app"

type Gate struct {
    app.App
}

func (g *Gate) BootPrepare() {
    mgr := SessionMgr{}
    g.RegisterComponent(mgr)
    mgr.BootPrepare()
}