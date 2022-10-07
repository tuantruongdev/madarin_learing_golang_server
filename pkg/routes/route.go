package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(router *gin.Engine, db *gorm.DB) {
	RegisterWord(router, db)
	RegisterAuth(router, db)
}
