package app

import (
    "os"

    "engine/logs"
)

type BootStage int

const (
    BootStagePrepare BootStage = 1 + iota
    BootStageMutualConnect
    BootStageDataLoad
    BootStagePostDataLoad
    BootStageBootFinished
    MaxBootStage
)

func newStage(stage BootStage) BootStage {
    if stage < MaxBootStage {
        return stage + 1
    }
    return stage
}

type IBootUnit interface {
    BootPrepare()
    OnBootStage()
}

type BootUnit struct {
    IBootUnit
    stage BootStage
}

func NewBootUnit() *BootUnit {
    bootUnit := BootUnit {
        stage: BootStagePrepare,
    }
    return &bootUnit
}

func (b *BootUnit) OnBootStage() {
    switch b.stage {
    case BootStagePrepare:
        b.BootPrepare()
    default:
        logs.Error("boot failed")
        logs.Flush()
        os.Exit(-1)
    }
}

func (b *BootUnit) onAppEnterStageEnd() {
    b.stage = newStage(b.stage)

    if MaxBootStage == b.stage {
        b.bootFinished()
    }

    logs.Info("current boot unit enter next stage[%d]", b.stage)
}

func (b *BootUnit) bootFinished() {
    logs.Info("current boot unit boot finished")
}