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
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.POST("/createPost", authMiddleWare(authService, userService), postHandler.CreatePost)

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
			response := helper.ApiResponse("unauthorized", http.StatusUnauthorized, "error", nil)
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
