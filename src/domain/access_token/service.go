package access_token

import (
	"strings"

	"github.com/HunnTeRUS/bookstore_oauth-api/src/utils/errors"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
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
	access_token_id = strings.TrimSpace(access_token_id)

	if len(access_token_id) == 0 {
		return nil, errors.NewBadRequestError("invalid access token id")
	}

	accessToken, err := s.repository.GetById(access_token_id)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (s *service) Create(access_token AccessToken) *errors.RestErr {
	if err := access_token.Validate(); err != nil {
		return err
	}

	return s.repository.Create(access_token)
}

func (s *service) UpdateExpirationTime(access_token AccessToken) *errors.RestErr {
	if err := access_token.Validate(); err != nil {
		return err
	}

	return s.repository.UpdateExpirationTime(access_token)
}
