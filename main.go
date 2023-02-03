package main

import (
	"growfunding/auth"
	"growfunding/campaign"
	"growfunding/handler"
	"growfunding/helper"
	"growfunding/user"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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
	campaignRepostitory := campaign.NewRepository(db)

	userService := user.NewService(userRepository)
	authService := auth.NewService()
	campaignService := campaign.NewService(campaignRepostitory)

	campaignInput := campaign.CreateCampaignInput{
		Name:             "Rizky pangestu update",
		ShortDescription: "short desc update",
		Description:      "desc update",
		GoalAmount:       20000000,
		Perks:            "perks 1 update, perks 2",
	}

	campaignIDInput := campaign.GetCampaignDetailInput{
		ID: 8,
	}

	campaignService.UpdataeCampaign(campaignIDInput, campaignInput)

	userHandler := handler.NewHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	router := gin.Default()
	router.Static("images/user", "./images/user")
	router.Static("images/campaign", "./images/campaign")

	api := router.Group("/api/v1")

	// user
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.POST("/email-checker", userHandler.CheckEmailAvability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	// campaign
	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaigns", authMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)

	router.Run(":4000")
}

// middleware auth
func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", 401, "error", nil)
			c.AbortWithStatusJSON(401, response)
			return
		}

		// ambil hanya nilai token
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			response := helper.APIResponse("Unauthorized", 401, "error", nil)
			c.AbortWithStatusJSON(401, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", 401, "error", nil)
			c.AbortWithStatusJSON(401, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)

		if err != nil {
			response := helper.APIResponse("Unauthorized", 401, "error", nil)
			c.AbortWithStatusJSON(401, response)
			return
		}

		c.Set("currentUser", user)
	}
}
