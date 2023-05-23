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
	}
	newUser, err := h.service.RegisterUser(userInput)
	if err != nil {
		c.JSON(422, gin.H{"errors": err.Error()})
	}
	c.JSON(200, newUser)
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
	}
	token, err := h.auth.GenerateToken(logginUser.Id)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Login failed", 422, "error", errorMessage)
		c.JSON(422, response)
	}
	formatUser := user.FormatUser(logginUser, token)
	response := helper.ApiResponse("Login success", 200, "success", formatUser)
	c.JSON(200, response)
	//c.JSON(200, user)
}
