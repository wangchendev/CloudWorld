package app

type App struct {
    comps []interface{}
}

func (a *App) BootPrepare() {
}

func (a *App) RegisterComponent(comp Component) {
    // TODO: 1. 单例  2. 检查重注册
    a.comps = append(a.comps, comp)
}