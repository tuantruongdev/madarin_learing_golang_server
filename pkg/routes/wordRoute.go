package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mandarinLearningBE/pkg/controllers"
)

func RegisterWord(router *gin.Engine, db *gorm.DB) {
	word := router.Group("/api/v1/word")
	{
		word.GET("/audio/:character").
			GET("/lookup/:character", controllers.LookupCharacter(db)).
			GET("/sentences/:character", controllers.LookupExample(db))
	}
}
