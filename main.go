package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/xvbnm48/go-network-media/auth"
	"github.com/xvbnm48/go-network-media/config"
	"github.com/xvbnm48/go-network-media/handler"
	"github.com/xvbnm48/go-network-media/helper"
	"github.com/xvbnm48/go-network-media/post"
	"github.com/xvbnm48/go-network-media/user"
	"log"
	"net/http"
	"strings"
)

func main() {
	db, err := config.SetUpDatabase()
	if err != nil {
		log.Fatal(err)
	}
	err = config.RunMigrations(db)
	// user repo
	userRepository := user.NewRepository(db)
	postRepository := post.NewPostRepository(db)

	//service user
	userService := user.NewService(userRepository)
	postService := post.NewServicePost(postRepository)

	//auth service
	authService := auth.NewService()
	// handler user
	userHandler := handler.NewUserHandler(userService, authService)
	postHandler := handler.NewPostHandler(postService)

	router := gin.Default()

	//user
	api := router.Group("/api/v1")
	api.POST("/user/register", userHandler.RegisterUser)
	api.POST("/user/login", userHandler.LoginUser)
	api.POST("/users/:id/follow", authMiddleWare(authService, userService), userHandler.FollowFriend)
	api.POST("/users/:id/unfollow", authMiddleWare(authService, userService), userHandler.UnfollowFriend)
	api.GET("/users/:id", authMiddleWare(authService, userService), userHandler.GetUserById)
	api.GET("/users", authMiddleWare(authService, userService), userHandler.FetchUser)

	// post
	api.POST("/post", authMiddleWare(authService, userService), postHandler.CreatePost)
	api.POST("/post/:id", authMiddleWare(authService, userService), postHandler.UpdatePost)
	api.GET("/posts", postHandler.GetAllPost)
	api.POST("/post/:id/delete", authMiddleWare(authService, userService), postHandler.DeletePost)
	api.GET("/post/:id", postHandler.GetPostById)

	router.Run(":8081")
}

func authMiddleWare(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			error := helper.FormatValidationError(err)
			errorMessage := gin.H{"errors": error}
			response := helper.ApiResponse("unauthorized", http.StatusUnauthorized, "error", errorMessage)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userId := int(claim["user_id"].(float64))
		user, err := userService.GetUserById(userId)
		if err != nil {
			response := helper.ApiResponse("unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)
	}
}
