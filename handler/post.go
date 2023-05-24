package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xvbnm48/go-network-media/helper"
	"github.com/xvbnm48/go-network-media/model"
	"github.com/xvbnm48/go-network-media/post"
)

type postHandler struct {
	service post.Service
}

func NewPostHandler(service post.Service) *postHandler {
	return &postHandler{service: service}
}

func (h *postHandler) CreatePost(c *gin.Context) {
	var input post.PostInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.ApiResponse("Create post failed", 422, "error", nil)
		c.JSON(422, response)
		return
	}

	currentUser := c.MustGet("currentUser").(model.User)
	fmt.Println(currentUser.Id)
	fmt.Println(currentUser.Name)
	input.Author = currentUser.Name

	NewPost, err := h.service.CreatePost(input)
	if err != nil {
		response := helper.ApiResponse("Create post failed", 422, "error", nil)
		c.JSON(422, response)
		return
	}

	response := helper.ApiResponse("Create post success", 200, "success", NewPost)
	c.JSON(200, response)
}
