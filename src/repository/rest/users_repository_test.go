package rest

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	fmt.Println("about to start test cases")
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserNewRepository(t *testing.T) {
	repo := NewUserRestRepository()
	assert.NotNil(t, repo)
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8081/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"email@email.com","password":"password"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})

	repository := usersRestRepository{}

	user, err := repository.LoginUser("email@email.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Code, http.StatusInternalServerError)
	assert.EqualValues(t, "invalid rest client response when trying to login user", err.Message)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8081/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"email@email.com","password":"password"}`,
		RespHTTPCode: http.StatusNotExtended,
		RespBody:     `{"message": "invalid login credentials", "status": "404", "error": "not_found"}`,
	})

	repository := usersRestRepository{}

	user, err := repository.LoginUser("email@email.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Code, http.StatusInternalServerError)
	assert.EqualValues(t, "invalid error when trying to login user", err.Message)
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8081/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"email@email.com","password":"password"}`,
		RespHTTPCode: http.StatusNotExtended,
		RespBody:     `{"message": "invalid user credentials", "status": 404, "error": "not_found"}`,
	})

	repository := usersRestRepository{}

	user, err := repository.LoginUser("email@email.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Code, http.StatusNotFound)
	assert.EqualValues(t, "invalid user credentials", err.Message)
}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8081/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"email@email.com","password":"password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": "1", "first_name": otavio, "last_name": "celestino dos santos", "email": "otavio@otavio.com"}`,
	})

	repository := usersRestRepository{}

	user, err := repository.LoginUser("email@email.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Code, http.StatusInternalServerError)
	assert.EqualValues(t, "invalid error when trying to login user", err.Message)
}

func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "http://localhost:8081/users/login",
		HTTPMethod:   http.MethodPost,
		ReqBody:      `{"email":"email@email.com","password":"password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": 1, "first_name": "otavio", "last_name": "celestino dos santos", "email": "otavio@otavio.com"}`,
	})

	repository := usersRestRepository{}

	user, err := repository.LoginUser("email@email.com", "password")

	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.EqualValues(t, user.Email, "otavio@otavio.com")
	assert.EqualValues(t, user.Id, 1)
	assert.EqualValues(t, user.FirstName, "otavio")
	assert.EqualValues(t, user.LastName, "celestino dos santos")
}
