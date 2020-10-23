package gate

import (
    "os"

    "engine/app"
    "engine/env"
    "engine/logs"
)

type SessionMgr struct {
    listener interface{}
}

func (g *SessionMgr) BootPrepare() {
    listener := createListener()
    listener.Init()

    listener.Start()
}

func createListener() IListener {
   address := getListenAddress("ServicePort")
   return newTcpListener(address)
}

func getListenAddress(key string) string {
   port, err := app.ServiceConfig.String(key)
   if err != nil {
       logs.Error("%s is not configured", key)
       logs.Flush()
       os.Exit(-1)
   }

   address := env.HostIP() + ":" + port
   logs.Info("the listening address is %s", address)
   return address
}