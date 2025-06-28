package rest

import (
	"github.com/neo7337/go-microservice-template/internal/api/rest/router"
	"github.com/neo7337/go-microservice-template/internal/config"
	"oss.nandlabs.io/golly/lifecycle"
	"oss.nandlabs.io/golly/rest"
)

var componentManager lifecycle.ComponentManager

func GetRestServer(manager lifecycle.ComponentManager) lifecycle.Component {
	// Load application configuration
	conf := config.GetConfig()
	serverOptions := rest.DefaultSrvOptions()
	serverOptions.PathPrefix = "/api"
	serverOptions.ListenPort = int16(conf.System.Port)
	serverOptions.ReadTimeout = int64(conf.System.ReadTimeout)   // seconds
	serverOptions.WriteTimeout = int64(conf.System.WriteTimeout) // seconds
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
