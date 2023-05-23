package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xvbnm48/go-network-media/auth"
	"github.com/xvbnm48/go-network-media/helper"
	"github.com/xvbnm48/go-network-media/user"
)

type userHandler struct {
	service user.Service
	auth    auth.Service
}

func NewUserHandler(service user.Service, auth auth.Service) *userHandler {
	return &userHandler{service, auth}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var userInput user.RegisterUserInput
	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := helper.ApiResponse("Register account failed", 422, "error", errors)
		response := helper.ApiResponse("Register account failed", 422, "error", errorMessage)
		c.JSON(422, response)
		return
	}
	newUser, err := h.service.RegisterUser(userInput)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := helper.ApiResponse("Register account failed", 422, "error", errors)
		response := helper.ApiResponse("Register account failed", 422, "error", errorMessage)
		c.JSON(422, response)
		return
	}
	token, err := h.auth.GenerateToken(newUser.Id)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := helper.ApiResponse("Register account failed", 422, "error", errors)
		response := helper.ApiResponse("Register account failed", 422, "error", errorMessage)
		c.JSON(422, response)
		return
	}
	format := user.FormatUser(newUser, token)
	response := helper.ApiResponse("Account has been registered", 200, "success", format)
	c.JSON(200, response)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var userInput user.LoginUserInput
	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		c.JSON(422, gin.H{"errors": err.Error()})
	}
	logginUser, err := h.service.LoginUser(userInput)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Login failed", 422, "error", errorMessage)
		c.JSON(422, response)
		return
	}
	token, err := h.auth.GenerateToken(logginUser.Id)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Login failed", 422, "error", errorMessage)
		c.JSON(422, response)
		return
	}
	formatUser := user.FormatUser(logginUser, token)
	response := helper.ApiResponse("Login success", 200, "success", formatUser)
	c.JSON(200, response)
	//c.JSON(200, user)
}
