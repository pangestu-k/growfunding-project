package handler

import (
	"growfunding/helper"
	"growfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service user.Service
}

func NewHandler(service user.Service) *userHandler {
	return &userHandler{service}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	input := user.RegisterUserInput{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{
			"errors": errors,
		}
		response := helper.APIResponse("Register user failed", 422, "error", errorMessage)
		c.JSON(422, response)
		return
	}

	newUser, err := h.service.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register user failed", 400, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userFormat := user.FormatUser(newUser, "token-generate")
	response := helper.APIResponse("Register user success", 200, "success", userFormat)

	c.JSON(http.StatusOK, response)
}
