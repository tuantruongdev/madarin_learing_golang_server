package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterAuth(router *gin.Engine) {
	account := router.Group("/api/v1/account")
	{
		account.POST("/login").
			POST("/signup").
			POST("/logout").
			POST("/forgot").
			GET("/login")
	}
}
