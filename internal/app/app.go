package app

import (
	"github.com/neo7337/go-microservice-template/internal/api/rest"
	"oss.nandlabs.io/golly/lifecycle"
)

var manager = lifecycle.NewSimpleComponentManager()

func Start() {
	restServer := rest.GetRestServer(manager)
	manager.Register(restServer)
	manager.StartAndWait()
}

func Shutdown() {
	manager.StopAll()
}
