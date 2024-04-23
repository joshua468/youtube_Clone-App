package main

import (
	"net/http"
	"os"

	"gorm.io/driver/mysql"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/youtube_Clone/handlers"
	"github.com/joshua468/youtube_Clone/repository"
	"github.com/joshua468/youtube_Clone/services"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Panic().Err(err).Msg("Error connecting to MySQL database")
	}
	defer func() {
		dbSQL, err := db.DB()
		if err != nil {
			logger.Error().Err(err).Msg("Error getting database connection")
		}
		dbSQL.Close()
	}()
	userRepo := repository.NewUserRepository(db)
	videoRepo := repository.NewVideoRepository(db)

	userService := services.NewUserService(userRepo)
	videoService := services.NewVideoService(videoRepo)

	userHandler := handlers.NewUserHandlers(userService)
	videoHandler := handlers.NewVideoHandlers(videoService)

	router := gin.Default()

	userHandler.RegisterRoutes(router)
	videoHandler.RegisterRoutes(router)

	port := "8080"
	if err := http.ListenAndServe(":"+port, router); err != nil {
		logger.Panic().Err(err).Msg("Error starting HTTP server")
	}
}
