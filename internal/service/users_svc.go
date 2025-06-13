package service

import (
	"github.com/neo7337/go-microservice-template/internal/config"
	"github.com/neo7337/go-microservice-template/internal/repository"
	"github.com/neo7337/go-microservice-template/pkg"
)

type UsersService struct {
	repo   repository.UsersRepository
	config *config.Config
}

func NewUsersService(config *config.Config) *UsersService {
	return &UsersService{
		repo:   repository.Get[repository.UsersRepository](repository.UsersRepo),
		config: config,
	}
}

func (s *UsersService) GetUsers() ([]pkg.User, error) {
	users, err := s.repo.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
