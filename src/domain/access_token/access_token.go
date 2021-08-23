package access_token

import (
	"fmt"
	"strings"
	"time"

	"github.com/HunnTeRUS/bookstore_oauth-api/src/utils/crypto_utils"
	"github.com/HunnTeRUS/bookstore_oauth-api/src/utils/errors"
)

const (
	expirationTime = 24
)

var (
	GrantTypePassword          = "password"
	GrantTypeClientCredentials = "client_credentials"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)

	if at.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token id")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}
	if at.ClientId <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}

	return nil
}

type AccessTokenRequest struct {
	GrantType    string `json:"grant_type"`
	Scope        string `json:"scope"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at *AccessTokenRequest) Validate() *errors.RestErr {
	switch at.GrantType {
	case GrantTypePassword:
		return at.validatePasswordGrantType()
	case GrantTypeClientCredentials:
		return at.validateClientGrantType()
	default:
		return errors.NewBadRequestError("invalid grant type")
	}
}

func (at *AccessTokenRequest) validatePasswordGrantType() *errors.RestErr {
	if strings.TrimSpace(at.Username) == "" {
		return errors.NewBadRequestError("username can not be empty")
	}

	if strings.TrimSpace(at.Password) == "" {
		return errors.NewBadRequestError("password can not be empty")
	}

	return nil
}

func (at *AccessTokenRequest) validateClientGrantType() *errors.RestErr {
	if strings.TrimSpace(at.ClientId) == "" {
		return errors.NewBadRequestError("client_id can not be empty")
	}

	if strings.TrimSpace(at.ClientSecret) == "" {
		return errors.NewBadRequestError("client_secret can not be empty")
	}

	return nil
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}
