package rest

import (
	"github.com/neo7337/go-microservice-template/internal/api/rest/router"
	"oss.nandlabs.io/golly/lifecycle"
	"oss.nandlabs.io/golly/rest"
)

var componentManager lifecycle.ComponentManager

func GetRestServer(manager lifecycle.ComponentManager) lifecycle.Component {
	serverOptions := rest.DefaultSrvOptions()
	serverOptions.PathPrefix = "/api"
	serverOptions.ListenPort = 8282
	componentManager = manager

	server, err := rest.NewServer(serverOptions)
	if err != nil {
		panic(err)
	}
	server.OnChange(func(prevState, newState lifecycle.ComponentState) {
		if prevState == lifecycle.Unknown && newState == lifecycle.Starting {
			server = router.RouterHandler(server)
		}
	})
	return server
}
