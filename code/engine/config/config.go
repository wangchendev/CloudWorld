package config

import (
    "engine/logs"
    "flag"
    "os"
)

var (
    ServiceConfig *YamlConfig
    ConfigFile string
)

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
        logs.Error("can not parse config file %s", ConfigFile)
        logs.Flush()
        os.Exit(-1)
    }

    ServiceConfig = GetConfigItem(cfg, "Develop")
}

func Init() {
    initFromArgs()
    initFromConfFile()
}
