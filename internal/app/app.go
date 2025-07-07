package app

import (
	"github.com/neo7337/go-microservice-template/internal/api/rest"
	"github.com/neo7337/go-microservice-template/internal/cache"
	"oss.nandlabs.io/golly/lifecycle"
)

var manager = lifecycle.NewSimpleComponentManager()

func Start() {
	cacheConnection := cache.GetConnection(manager)
	manager.Register(cacheConnection)
	restServer := rest.GetRestServer(manager)
	manager.Register(restServer)
	manager.AddDependency(restServer.Id(), cacheConnection.Id())
	manager.StartAndWait()
}

func Shutdown() {
	manager.StopAll()
}
