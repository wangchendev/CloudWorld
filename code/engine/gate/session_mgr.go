package gate

import (
    "fmt"
    "os"

    "engine/app"
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
   fmt.Printf("Info: address %s\n", address)
   return nil
   //return newTcpListener(address, true)
}

func getListenAddress(key string) string {
   port, err := app.ServiceConfig.String(key)
   if err != nil {
       fmt.Printf("Error: %s is not configured\n", key)
       //logs.Flush()
       os.Exit(-1)
   }

   address := "127.0.0.1" + ":" + port
   fmt.Printf("Info: the listening address is %s\n", address)
   return address
}