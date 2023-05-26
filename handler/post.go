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
	input.User.Id = currentUser.Id

	NewPost, err := h.service.CreatePost(input)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": error}
		response := helper.ApiResponse("Create post failed", 422, "error", errorMessage)
		c.JSON(422, response)
		return
	}
	formatPost := post.FormatPost(NewPost)
	response := helper.ApiResponse("Create post success", 200, "success", formatPost)
	c.JSON(200, response)
}

//func (h *postHandler) UpdatePost(c *gin.Context) {
//	var inputID post.GetPostDetailInput
//	err := c.ShouldBindUri(&inputID)
//	if err != nil {
//		errors := helper.FormatValidationError(err)
//		messageErrors := gin.H{"errors": errors}
//		response := helper.ApiResponse("id post not found", 422, "error", messageErrors)
//		c.JSON(422, response)
//		return
//	}
//	Oldpost := post.PostInput{}
//	err = c.ShouldBindJSON(&Oldpost)
//	if err != nil {
//		errors := helper.FormatValidationError(err)
//		messageErrors := gin.H{"errors": errors}
//		response := helper.ApiResponse("update post failed", 422, "error", messageErrors)
//		c.JSON(422, response)
//		return
//	}
//
//	currentUser := c.MustGet("currentUser").(model.User)
//	userId := currentUser.Id
//	//Oldpost.Author = currentUser.Name
//	Oldpost.Author = currentUser.Name
//
//	UpdatedPost, err := h.service.UpdatePost(Oldpost, inputID.Id, userId)
//	if err != nil {
//		errors := helper.FormatValidationError(err)
//		messageErrors := gin.H{"errors": errors}
//		response := helper.ApiResponse("update post failed", 422, "error", messageErrors)
//		c.JSON(422, response)
//		return
//	}
//
//	formatPost := post.FormatPost(UpdatedPost)
//	response := helper.ApiResponse("update post success", 200, "success", formatPost)
//	c.JSON(200, response)
//}

func (h *postHandler) UpdatePost(c *gin.Context) {
	var inputId post.GetPostDetailInput
	err := c.ShouldBindUri(&inputId)
	fmt.Println(inputId.Id)
	if err != nil {
		response := helper.ApiResponse("id post not found", 422, "error", nil)
		c.JSON(422, response)
		return
	}

	var inputPost post.UpdatePost
	err = c.ShouldBindJSON(&inputPost)
	fmt.Println("isi post: ", inputPost)
	currentUser := c.MustGet("currentUser").(model.User)
	inputPost.User = currentUser
	inputPost.Author = currentUser.Name
	inputPost.User.Id = currentUser.Id
	fmt.Println("id dari jwt:", currentUser.Id)
	//fmt.Println("isi input post setelah di isi currentUser", inputPost)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": error}
		response := helper.ApiResponse("update post failed", 422, "error", errorMessage)
		c.JSON(422, response)
		return
	}

	updatePost, err := h.service.UpdatePost(inputId, inputPost)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": error}
		response := helper.ApiResponse("update post failed", 422, "error", errorMessage)
		c.JSON(422, response)
		return
	}

	formatPost := post.FormatPost(updatePost)
	response := helper.ApiResponse("update post success", 200, "success", formatPost)
	c.JSON(200, response)
}

func (h *postHandler) GetAllPost(c *gin.Context) {
	posts, err := h.service.GetAllPost()
	if err != nil {
		message := helper.FormatValidationError(err)
		ErrorMessage := gin.H{"errors": message}
		response := helper.ApiResponse("Get all post failed", 422, "error", ErrorMessage)
		c.JSON(422, response)
		return
	}

	formatPost := post.FormatPosts(posts)
	response := helper.ApiResponse("Get all post success", 200, "success", formatPost)
	c.JSON(200, response)
}

func (h *postHandler) DeletePost(c *gin.Context) {
	postId := post.GetPostDetailInput{}
	err := c.ShouldBindUri(&postId)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": error}
		response := helper.ApiResponse("id post not found", 422, "error", errorMessage)
		c.JSON(422, response)
		return
	}
	currentUser := c.MustGet("currentUser").(model.User)
	err = h.service.DeletePost(postId.Id, currentUser.Id)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": error}
		response := helper.ApiResponse("id post not found", 422, "error", errorMessage)
		c.JSON(422, response)
		return
	}

	SuccessMessage := gin.H{
		"is_deleted": true,
		"message":    "success deleted post! with id: " + string(postId.Id),
	}

	response := helper.ApiResponse("delete post success", 200, "success", SuccessMessage)
	c.JSON(200, response)
}

func (h *postHandler) GetPostById(c *gin.Context) {
	postId := post.GetPostDetailInput{}
	err := c.ShouldBindUri(&postId)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": error}
		response := helper.ApiResponse("id post not found", 422, "error", errorMessage)
		c.JSON(422, response)
		return
	}
	Findpost, err := h.service.GetPostById(postId.Id)
	if err != nil {
		error := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": error}
		response := helper.ApiResponse("id post not found", 422, "error", errorMessage)
		c.JSON(422, response)
		return
	}
	formatPost := post.FormatPost(Findpost)
	response := helper.ApiResponse("Get post by id success", 200, "success", formatPost)
	c.JSON(200, response)
}
