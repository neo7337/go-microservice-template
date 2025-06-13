package handler

import (
	"github.com/neo7337/go-microservice-template/internal/service"
	"github.com/neo7337/go-microservice-template/pkg"
	"oss.nandlabs.io/golly/l3"
	"oss.nandlabs.io/golly/rest"
)

var logger = l3.Get()

type UsersHandler struct {
	svc *service.UsersService
}

func NewUsersHandler(svc *service.UsersService) *UsersHandler {
	return &UsersHandler{
		svc: svc,
	}
}

func (h *UsersHandler) UsersHandler(ctx rest.ServerContext) {
	users, err := h.svc.GetUsers()
	if err != nil {
		logger.Error("Failed to get users", "error", err)
		pkg.ResponseJSON(ctx, 500, map[string]string{"error": "Internal Server Error"})
		return
	}

	pkg.ResponseJSON(ctx, 200, users)
}
