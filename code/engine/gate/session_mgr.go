package gate

import (
    "os"

    "engine/app"
    "engine/logs"
)

type SessionMgr struct {
    listener interface{}
}

func (g *SessionMgr) BootPrepare() {
    createListener()
    //listener := createListener()
    //listener.Init()
    //
    //listener.Start()
}

func createListener() IListener {
   address := getListenAddress("ServicePort")
   logs.Info("address %s", address)
   return nil
   //return newTcpListener(address, true)
}

func getListenAddress(key string) string {
   port, err := app.ServiceConfig.String(key)
   if err != nil {
       logs.Error("%s is not configured", key)
       logs.Flush()
       os.Exit(-1)
   }

   address := "127.0.0.1" + ":" + port
   logs.Info("the listening address is %s", address)
   return address
}