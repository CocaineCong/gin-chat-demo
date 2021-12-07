package main

import (
	"chat/conf"
	"chat/router"
	"chat/service"
)

func main() {
	conf.Init()
	go service.Manager.Start()
	r:=router.NewRouter()
	_ = r.Run(conf.HttpPort)
}
