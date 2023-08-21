package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func Routes(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "pong",
			"app_name": viper.Get("App.Name"), // We can get the keys through config struct
			"email":    viper.Get("Auth.Email"),
			"password": viper.Get("Auth.Password"),
		})
	})
}
