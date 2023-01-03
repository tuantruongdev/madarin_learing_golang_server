package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mandarinLearningBE/pkg/controllers"
)

func RegisterSyncWord(router *gin.Engine, db *gorm.DB) {
	word := router.Group("/api/v1/sync")
	{
		//not complete
		word.POST("/word", controllers.Authorize(db), controllers.SaveFavoriteCharacter(db)).
			GET("/word", controllers.Authorize(db), controllers.QueryFavorite(db))
	}
}
