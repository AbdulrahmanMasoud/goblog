package routes

import (
	homeRoute "github.com/AbdulrahmanMasoud/goblog/internal/modules/home/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	homeRoute.Routes(router)
}
