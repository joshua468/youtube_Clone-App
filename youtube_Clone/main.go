package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joshua468/youtube_Clone/handlers"
	"github.com/joshua468/youtube_Clone/repository"
	"github.com/joshua468/youtube_Clone/services"
)

func main() {
	// Initialize MySQL database connection
	db, err := sql.Open("mysql", "Temi2080:170821002@tcp(127.0.0.1:3306)/database_name")
	if err != nil {
		panic("Error connecting to MySQL database: " + err.Error())
	}
	defer db.Close()

	// Initialize repository
	userRepo := repository.NewUserRepository(db)
	videoRepo := repository.NewVideoRepository(db)

	// Initialize services
	userService := services.NewUserService(*userRepo)
	videoService := services.NewVideoService(*videoRepo)

	// Initialize handlers
	userHandler := handlers.NewUserHandlers(userService)
	videoHandler := handlers.NewVideoHandlers(videoService)

	router := gin.Default()

	// Register routes using Gin router
	userHandler.RegisterRoutes(router)
	videoHandler.RegisterRoutes(router)

	// Start HTTP server
	port := "8080" // Or any port you prefer
	if err := http.ListenAndServe(":"+port, router); err != nil {
		panic("Error starting HTTP server: " + err.Error())
	}
}
