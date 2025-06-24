package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"todo-list/internal/application"
	"todo-list/internal/infrastructure"
	"todo-list/internal/interfaces"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: .env files not found, using default values")
	} else {
		log.Println("Loaded environment from .env")
	}

	mode := os.Getenv("MODE")
	if mode == "" {
		mode = "debug"
	}
	gin.SetMode(mode)

	db, err := infrastructure.NewDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	todoRepo := infrastructure.NewTodoRepository(db)
	userRepo := infrastructure.NewUserRepository(db)

	todoService := application.NewTodoService(todoRepo)
	userService := application.NewUserService(userRepo)

	todoHandler := interfaces.NewTodoHandler(todoService)
	userHandler := interfaces.NewUserHandler(userService)

	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	interfaces.SetupRoutes(router, todoHandler, userHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s in %s mode", port, mode)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
