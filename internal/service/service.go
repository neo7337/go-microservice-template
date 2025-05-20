package service

import (
	"github.com/neo7337/go-microservice-template/internal/repository"
	"github.com/neo7337/go-microservice-template/pkg"
)

func GetUsers() ([]pkg.User, error) {
	return repository.GetUsers()
}
