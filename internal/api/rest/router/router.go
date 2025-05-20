package router

import (
	"github.com/neo7337/go-microservice-template/internal/api/rest/handler"
	"github.com/neo7337/go-microservice-template/pkg"
	"oss.nandlabs.io/golly/rest"
)

func RouterHandler(server rest.Server) rest.Server {
	server.Get("/healthz", func(ctx rest.ServerContext) {
		pkg.ResponseJSON(ctx, 200, map[string]string{"status": "ok"})
	})

	server.Get("/users", handler.UsersHandler)
	return server
}
