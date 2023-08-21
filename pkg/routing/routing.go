package routing

import (
	"github.com/AbdulrahmanMasoud/goblog/internal/providers/routes"
	"github.com/gin-gonic/gin"
)

func Init() {
	router = gin.Default()
}

func GetRouter() *gin.Engine {
	return router
}

func RegisterRouter() {
	routes.RegisterRoutes(GetRouter())
}
