package main

import (
	"log"
	"os"

	"github.com/brkss/dextrace/internal/delivery"
	"github.com/brkss/dextrace/internal/domain"
	"github.com/brkss/dextrace/internal/infrastructure"
	"github.com/brkss/dextrace/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// todo : move configs to seprate package ! 
	user := domain.User{
		Email:    os.Getenv("USER_EMAIL"),
		Password: os.Getenv("USER_PASSWORD"),
	}
	userID := os.Getenv("USER_ID")

	nightscoutConfig := domain.NightscoutConfig{
		Token: os.Getenv("NIGHTSCOUT_API_KEY"),
		NightscoutURL: os.Getenv("NIGHTSCOUT_API_URL"),
	}

	// Initialize repositories
	sibionicRepo := infrastructure.NewSibionicRepository(os.Getenv("API_URL"))
	nighscoutRepo := infrastructure.NewNightscoutRepository(nightscoutConfig.NightscoutURL, nightscoutConfig.Token)

	// Initialize use cases
	sibionicUseCase := usecase.NewSibionicUseCase(sibionicRepo, sibionicRepo)
	nighscoutUseCase := usecase.NewNightscoutUseCase(nighscoutRepo)

	// Initialize handlers
	handler := delivery.NewGlucoseHandler(sibionicUseCase, nighscoutUseCase, userID, user)

	// Setup router
	r := gin.Default()
	r.GET("/data", handler.GetGlucoseData)
	r.POST("/push-to-nightscout", handler.PushToNightscout);

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}