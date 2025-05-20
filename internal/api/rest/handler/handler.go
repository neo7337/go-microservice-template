package handler

import (
	"github.com/neo7337/go-microservice-template/internal/service"
	"github.com/neo7337/go-microservice-template/pkg"
	"oss.nandlabs.io/golly/rest"
)

func UsersHandler(ctx rest.ServerContext) {
	users, err := service.GetUsers()
	if err != nil {
		pkg.ResponseJSON(ctx, 500, map[string]string{"error": err.Error()})
		return
	}
	pkg.ResponseJSON(ctx, 200, users)
}
