package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mandarinLearningBE/pkg/controllers"
)

func RegisterAuth(router *gin.Engine, db *gorm.DB) {
	account := router.Group("/api/v1/account")
	{
		//account.POST("/logout", controllers.Authorize(db), controllers.Logout(db))
		account.POST("/login", controllers.Login(db)).
			GET("/auth", controllers.Authorize(db)).
			POST("/signup", controllers.SignUp(db)).
			POST("/logout", controllers.Authorize(db), controllers.Logout(db)).
			PATCH("/changePass", controllers.Authorize(db), controllers.ChangePass(db)).
			POST("/forgot")
	}
}
