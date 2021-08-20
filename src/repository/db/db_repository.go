package db

import (
	"github.com/HunnTeRUS/bookstore_oauth-api/src/clients/cassandra"
	"github.com/HunnTeRUS/bookstore_oauth-api/src/domain/access_token"
	"github.com/HunnTeRUS/bookstore_oauth-api/src/utils/errors"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (db *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	session, err := cassandra.GetSession()
	defer session.Close()

	if err != nil {
		panic(err)
	}
	return nil, errors.NewInternalServerError("Database connection not implemented yet")
}
