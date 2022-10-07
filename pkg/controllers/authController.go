package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
	"mandarinLearningBE/pkg/models"
	"net/http"
	"strings"
)

func validateEmailAndPassword(context gin.Context) (models.User, bool) {
	var user models.User
	context.BindJSON(&user)
	user.JwtToken = ""
	if len(user.Email) == 0 || len(user.Password) < 0 {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "you must include email or password"})
		context.Abort()
		return user, false
	}
	if !models.ValidEmail(user.Email) {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "your email is not valid"})
		context.Abort()
		return user, false
	}
	if len(user.Password) < 6 {
		context.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "your password must be at least 6 characters"})
		context.Abort()
		return user, false
	}
	return user, true
}

func SignUp(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {

		user, ok := validateEmailAndPassword(*context)
		if !ok {
			return
		}
		if !models.IsUserExist(db, &user) {
			models.HashPassword(&user)
			models.InsertUser(db, &user)
			user, _ = models.GetUserInfo(db, user.Email)
			//us
			_ = models.SignJwt(&user)
			context.JSON(http.StatusOK, gin.H{"status": "ok", "message": "account created successfully", "token": user.JwtToken})
			return
		}
		context.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Your email has been linked with another account"})
	}
}

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		user, ok := validateEmailAndPassword(*context)
		if !ok {
			return
		}
		dbUser, err := models.GetUserInfo(db, user.Email)
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "account or password are in correct"})
			return
		}
		models.HashPassword(&user)
		if dbUser.Password != user.Password {
			//	fmt.Println(dbUser.Password, user.Password)
			context.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "account or password are in correct 2"})
			return
		}
		//fmt.Println(user)
		err = models.SignJwt(&dbUser)

		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"status": "error", "message": err.Error()})
		}
		context.JSON(http.StatusNotFound, gin.H{"status": "ok", "message": "login successfully", "token": dbUser.JwtToken})
	}
}

func Authorize(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var user models.User

		err := context.ShouldBindBodyWith(&user, binding.JSON)
		if err != nil {
			fmt.Println(err.Error())
		}

		jwtToken := context.GetHeader("Authorization")
		if len(jwtToken) < 1 {
			context.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "you're not logged in1"})
			context.Abort()
			return
		}
		jwtToken = strings.Split(jwtToken, " ")[1]
		if len(jwtToken) < 1 {
			context.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "you're not logged in32"})
			context.Abort()
			return
		}
		user.JwtToken = jwtToken
		//if no auth then stop, else continue to next middleware
		if !models.ValidateJwt(db, &user) {
			context.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "you're not logged in3"})
			context.Abort()
			return
		}
		//fmt.Println("auth ok,next middleware", user)
		context.Set("user", user)
		context.Set("authenticated", true)
		context.Next()
	}
}

func Logout(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		//there is 2 type, local delete token or server logout(single logout,force logout) in this case i gonna use force method because it will be something to do rather than local delete token but less work than single logout
		user, ok := context.Keys["user"].(models.User)
		if !ok {
			//spamming paste xd
			context.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "you're not logged in"})
			return
		}
		err := models.UpdatePasswordChangeTime(db, &user)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"status": "ok", "message": "you have logged out successfully"})
	}
}

func ChangePass(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var givenUser models.User
		context.ShouldBindBodyWith(&givenUser, binding.JSON)
		givenUser.JwtToken = ""

		//	fmt.Println(givenUser)
		if len(givenUser.Password) < 1 || len(givenUser.NewPassword) < 1 {
			context.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "you need to include your password"})
			return
		}
		if len(givenUser.NewPassword) < 6 {
			context.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "password must be more or equal than 6 character"})
			return
		}

		//less case
		user, ok := context.Keys["user"].(models.User)
		if !ok {
			//spamming paste xd
			context.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "you're not logged in"})
			return
		}

		//less case
		dbUser, err := models.GetUserInfo(db, user.Email)
		if err != nil {
			context.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "account or password are in correct"})
			return
		}
		//fmt.Println(dbUser.Password, givenUser.Password)
		models.HashPassword(&givenUser)
		if dbUser.Password != givenUser.Password {

			context.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "your current password is incorrect"})
			return
		}
		//	fmt.Println("password match")
		//	fmt.Println(user.NewPassword)
		dbUser.Password = givenUser.NewPassword
		//todo handle error here
		models.HashPassword(&dbUser)
		models.UpdatePassword(db, &dbUser)
		models.UpdatePasswordChangeTime(db, &dbUser)
		models.SignJwt(&dbUser)
		context.JSON(http.StatusCreated, gin.H{"status": "ok", "message": "password changed successfully", "token": dbUser.JwtToken})

	}
}
