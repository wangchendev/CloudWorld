package main

import (
    "engine/gate"
)

func main() {
    gateApp := gate.Gate{}
    gateApp.Init()
    gateApp.BootPrepare()
}
