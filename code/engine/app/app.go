package app

import (
    "flag"
    "fmt"
    "os"
)

var (
    ServiceConfig *YamlConfig
    ConfigFile string
)

type App struct {
    comps []interface{}
}

func (a *App) Init() {
    initFromArgs()
    initFromConfFile()
}

func initFromArgs() {
    if ConfigFile == "" {
        flag.StringVar(&ConfigFile, "conf", "", "support config file.")
    }
    flag.Parse()
}

func initFromConfFile() {
    var err error
    cfg, err := NewYamlFromFile(ConfigFile)
    if err != nil {
        fmt.Fprintf(os.Stderr, "can not parse config file %s\n", ConfigFile)
        os.Exit(-1)
    }

    ServiceConfig = GetConfigItem(cfg, "Develop")
}

func (a *App) BootPrepare() {
}

func (a *App) RegisterComponent(comp Component) {
    // TODO: 1. 单例  2. 检查重注册
    a.comps = append(a.comps, comp)
}