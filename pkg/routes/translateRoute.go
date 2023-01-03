package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mandarinLearningBE/pkg/controllers"
)

func RegisterTranslate(router *gin.Engine, db *gorm.DB) {
	word := router.Group("/api/v1/translate")
	{
		word.POST("", controllers.Translate(db))
	}
}
