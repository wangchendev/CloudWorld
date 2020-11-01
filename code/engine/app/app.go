package app

type App struct {
    *BootUnit
    comps []IComponent
}

func NewApp() *App {
    app := App {
        BootUnit: NewBootUnit(),
    }
    return &app
}

func (a *App) Init() {
    a.onAppEnterStageBegin()
}

func (a *App) onAppEnterStageBegin() {
    a.OnBootStage()
    for _, comp := range a.comps {
        comp.OnBootStage()
    }

    // TODO: 注册服务发现
}

func (a *App) RegisterComponent(comp IComponent) {
    // TODO: 1. 单例  2. 检查重注册
    a.comps = append(a.comps, comp)
}