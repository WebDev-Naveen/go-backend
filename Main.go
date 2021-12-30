package main

import (
	"log"

	"github.com/gin-gonic/gin"

	config "github.com/WebDev-Naveen/go-backend/config"
	routes "github.com/WebDev-Naveen/go-backend/routes"
)

func main() {
	// Connect DB
	config.Connect()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)
    
	log.Fatal(router.Run(":4747"))
}