package repository

import "github.com/neo7337/go-microservice-template/pkg"

func GetUsers() ([]pkg.User, error) {
	users := []pkg.User{
		{
			ID:    1,
			Name:  "John Doe",
			Age:   30,
			Email: "john.doe@mail.com",
		},
		{
			ID:    2,
			Name:  "Jane Smith",
			Age:   25,
			Email: "jane.smith@mail.com",
		},
	}

	return users, nil
}
