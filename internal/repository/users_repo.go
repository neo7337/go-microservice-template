package repository

import "github.com/neo7337/go-microservice-template/pkg"

type UsersRepository interface {
	GetUsers() ([]pkg.User, error)
}
