package controllers

import (
	"github.com/Tapu106/happy-meal/internal/models"
	"github.com/gin-gonic/gin"
)

func GetMeals(context *gin.Context) {
	meals := []models.Meal{}
	context.JSON(200, gin.H{
		"meals": meals,
	})
}
