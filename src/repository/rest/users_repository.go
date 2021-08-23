package rest

import (
	"encoding/json"
	"time"

	"github.com/HunnTeRUS/bookstore_oauth-api/src/domain/users"
	"github.com/HunnTeRUS/bookstore_oauth-api/src/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8081",
		Timeout: 100 * time.Millisecond,
	}
)

type UsersRestRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRestRepository struct {
}

func NewUserRestRepository() UsersRestRepository {
	return &usersRestRepository{}
}

func (userRest *usersRestRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request := users.LoginRequest{
		Email:    email,
		Password: password,
	}
	response := usersRestClient.Post("/users/login", request)

	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("invalid rest client response when trying to login user")
	}

	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error when trying to login user")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("invalid error when trying to login user")
	}

	return &user, nil
}
