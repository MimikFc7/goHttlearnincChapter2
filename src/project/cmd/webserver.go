package main

import (
	"fmt"
	"project/database/pgconnectror"
	"project/httpsrv"
	"project/web/controllers"
)

const (
	LISTEN_PORT = 8080
)

func main() {

	dbInstance := pgconnect.GetInstance()
	if dbInstance == nil {
		fmt.Println("error db not open")
		return
	}
	dbInstance.OpenConnect()

	h2 := userinfo.GetEP()
	var hadlers = []httpsrv.EPHandler{h2}
	httpServer := httpsrv.HTTPServer{
		Name:     "MainWeb",
		Port:     LISTEN_PORT,
		Handlers: hadlers,
	}
	httpServer.StartServer()
}
