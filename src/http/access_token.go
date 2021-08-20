package http

import (
	"net/http"

	"github.com/HunnTeRUS/bookstore_oauth-api/src/domain/access_token"
	"github.com/HunnTeRUS/bookstore_oauth-api/src/utils/errors"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(c *gin.Context)
	Create(c *gin.Context)
	UpdateExpirationTime(c *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	returned, err := h.service.GetById(c.Param("access_token_id"))

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, returned)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	var at access_token.AccessToken

	if err := c.ShouldBindJSON(&at); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("invalid json body for create an access token"))
		return
	}

	if err := h.service.Create(at); err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, at)
}

func (h *accessTokenHandler) UpdateExpirationTime(c *gin.Context) {
	returned, err := h.service.GetById(c.Param("access_token_id"))

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, returned)
}
