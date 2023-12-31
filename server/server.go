package server

import (
	"goweb/router"
	"log"
	"net/http"
)

var App = &MsServer{}

type MsServer struct {
}

func (*MsServer) Start(ip, port string) {
	server := http.Server{
		Addr: ip + ":" + port,
	}
	//路由
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
