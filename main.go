package main

import (
	"assignment2/config"
	"assignment2/routers"
	"fmt"
	"log"
	"os"

	_ "assignment2/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/joho/godotenv"
)

// @title Hacktiv8 Golang Assignment 2 - API Documentation
// @description This is a REST API documentation for Hacktiv8 Golang Assignment 2.
// @version 1.0
//
// @contact.name API Support
// @contact.email thedaamandaa@gmail.com
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config.StartDB()

	serverAddress := setupServer()
	router := routers.SetupRouter()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run(serverAddress)
	if err != nil {
		log.Fatal("Error running server: " + err.Error())
	}
}

func setupServer() string {
	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")
	return fmt.Sprintf("%s:%s", host, port)
}
