package main

import (
	"fmt"
	"growfunding/auth"
	"growfunding/handler"
	"growfunding/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/crowfunding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewHandler(userService, authService)

	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMX0.MGmCvCxgghh0tfE4sC1uMc5jm91BC9lTnVaqS7koowY")

	if err != nil {
		fmt.Println("Error")
	}

	if token.Valid {
		fmt.Println("Token JWT valid")
	} else {
		fmt.Println("Token JWT invalid")
	}

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.POST("/email-checker", userHandler.CheckEmailAvability)
	api.POST("/avatars", userHandler.UploadAvatar)
	router.Run(":4000")
}
