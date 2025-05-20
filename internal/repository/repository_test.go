package repository

import (
	"testing"

	"github.com/neo7337/go-microservice-template/pkg"
	"oss.nandlabs.io/golly/assertion"
)

func TestGetUsers(t *testing.T) {
	users, err := GetUsers()
	assertion.Empty(err)
	assertion.NotEmpty(users)
	assertion.Equal(2, len(users))

	expected := []pkg.User{
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

	for i, user := range users {
		assertion.Equal(expected[i].ID, user.ID)
		assertion.Equal(expected[i].Name, user.Name)
		assertion.Equal(expected[i].Age, user.Age)
		assertion.Equal(expected[i].Email, user.Email)
	}
}
