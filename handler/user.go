package handler

import (
	"fmt"
	"growfunding/auth"
	"growfunding/helper"
	"growfunding/user"
	"math/rand"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewHandler(service user.Service, authService auth.Service) *userHandler {
	return &userHandler{service, authService}
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

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register user failed", 400, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.APIResponse("Register user failed", 400, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userFormat := user.FormatUser(newUser, token)
	response := helper.APIResponse("Register user success", 200, "success", userFormat)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	input := user.LoginUserInput{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{
			"errors": errors,
		}
		response := helper.APIResponse("Login user failed", 422, "error", errorMessage)
		c.JSON(422, response)
		return
	}

	loginUser, err := h.userService.LoginUser(input)

	if err != nil {
		errorMessage := gin.H{
			"errors": err.Error(),
		}
		response := helper.APIResponse("Login user failed", 422, "error", errorMessage)
		c.JSON(422, response)
		return
	}

	token, err := h.authService.GenerateToken(loginUser.ID)
	if err != nil {
		response := helper.APIResponse("Login user failed", 400, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userFormat := user.FormatUser(loginUser, token)
	response := helper.APIResponse("Login user success", 200, "success", userFormat)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvability(c *gin.Context) {
	input := user.CheckEmailInput{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{
			"errors": errors,
		}
		response := helper.APIResponse("CheckEmail Avability failed", 422, "error", errorMessage)
		c.JSON(422, response)
		return
	}

	emailAvailable, err := h.userService.EmailIsAvailable(input)

	if err != nil {
		errorMessage := gin.H{
			"errors": "Ups something wrong please try again",
		}
		response := helper.APIResponse("Login user failed", 422, "error", errorMessage)
		c.JSON(422, response)
		return
	}

	var message string
	// var statusCode int
	// var code string

	if emailAvailable {
		message = "Email can use to register"
		// statusCode = 400
		// code = "failed"
	} else {
		message = "Email is available"
		// statusCode = 200
		// code = "success"
	}

	data := gin.H{
		"avability": emailAvailable,
	}

	// response := helper.APIResponse(message, statusCode, code, data)
	// c.JSON(http.StatusOK, response)
	response := helper.APIResponse(message, 200, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{
			"upload-file": false,
		}
		response := helper.APIResponse("Upload avatar failed", 422, "error", data)
		c.JSON(400, response)
		return
	}

	extFile := filepath.Ext(file.Filename)
	fileName := rand.Int()
	idUser := 1

	path := fmt.Sprintf("images/%d-user-%d%s", idUser, fileName, extFile)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{
			"upload-file": false,
		}
		response := helper.APIResponse("Upload avatar failed", 422, "error", data)
		c.JSON(400, response)
		return
	}

	_, err = h.userService.UploadAvatar(1, path)
	if err != nil {
		data := gin.H{
			"upload-file": false,
		}
		response := helper.APIResponse("Upload avatar failed", 422, "error", data)
		c.JSON(400, response)
		return
	}

	data := gin.H{
		"upload-file": true,
	}
	response := helper.APIResponse("Upload avatar success", 200, "success", data)
	c.JSON(200, response)
}
