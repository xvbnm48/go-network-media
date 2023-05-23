package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xvbnm48/go-network-media/auth"
	"github.com/xvbnm48/go-network-media/config"
	"github.com/xvbnm48/go-network-media/handler"
	"github.com/xvbnm48/go-network-media/user"
	"log"
)

func main() {
	db, err := config.SetUpDatabase()
	if err != nil {
		log.Fatal(err)
	}
	// user repo
	userRepository := user.NewRepository(db)

	//service user
	userService := user.NewService(userRepository)

	//auth service
	authService := auth.NewService()
	// handler user
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)

	router.Run()
}
