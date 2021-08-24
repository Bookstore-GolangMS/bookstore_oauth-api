package http

import (
	"net/http"

	"github.com/Bookstore-GolangMS/bookstore_utils-go/errors"
	"github.com/HunnTeRUS/bookstore_oauth-api/src/domain/access_token"
	access_token_service "github.com/HunnTeRUS/bookstore_oauth-api/src/services/access_token"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetById(c *gin.Context)
	Create(c *gin.Context)
	UpdateExpirationTime(c *gin.Context)
}

type accessTokenHandler struct {
	service access_token_service.Service
}

func NewHandler(service access_token_service.Service) AccessTokenHandler {
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
	var at access_token.AccessTokenRequest

	if err := c.ShouldBindJSON(&at); err != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequestError("invalid json body for create an access token"))
		return
	}

	if _, err := h.service.Create(at); err != nil {
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
