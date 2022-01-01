package main

import (
	"fmt"
	"log"

	routes "github.com/WebDev-Naveen/go-backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadenv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	// Connect DB

	fmt.Println("Main Application Starts")
	//Loading Environmental Variable
	loadenv()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":8080"))
}
