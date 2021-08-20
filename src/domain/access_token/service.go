package access_token

import "github.com/HunnTeRUS/bookstore_oauth-api/src/utils/errors"

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(access_token_id string) (*AccessToken, *errors.RestErr) {
	return s.repository.GetById(access_token_id)
}
