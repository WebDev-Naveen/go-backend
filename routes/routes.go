package routes

import (
	"net/http"

	controllers "github.com/WebDev-Naveen/go-backend/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	userHandler := controllers.NewUserHandler()

	coinHandler := controllers.NewCoinHandler()
	router.GET("/", welcome)
	router.GET("/portfolio", userHandler.GetAllUser)
	router.GET("/portfolio/:user/coins", userHandler.GetUserCoins)

	router.POST("/portfolio/:user/entries", coinHandler.UserCoin)
	router.NoRoute(notFound)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})

}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})

}
