package routes

import (
	"github.com/Tapu106/happy-meal/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/meals", controllers.GetMeals)
}
