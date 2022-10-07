package models

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"regexp"
	"time"
)

type User struct {
	Id                int        `json:"id" gorm:"id"`
	Email             string     `json:"email" gorm:"email"`
	Password          string     `json:"password" gorm:"password"`
	NewPassword       string     `json:"newPassword" gorm:"-"`
	CreatedAt         *time.Time `json:"create_at" gorm:"created_at"`
	PasswordUpdatedAt *time.Time `json:"passwordUpdatedAt" gorm:"password_updated_at"`
	JwtToken          string     `json:"jwtToken" gorm:"-"`
}

func (*User) TableName() string {
	return "user"
}

func HashPassword(u *User) {
	hash := sha256.New()
	hash.Write([]byte(u.Password))
	bs := hash.Sum(nil)
	u.Password = hex.EncodeToString(bs)
	//return hex.EncodeToString(bs)
	//fmt.Println(hex.EncodeToString(bs))
}

func ValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

func SignJwt(user *User) error {
	var mySigningKey = []byte{69, 99, 77, 88, 55, 45, 45, 14, 14, 54, 77, 98, 98, 78, 11, 22, 33, 55}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = user.Email
	claims["updated_at"] = user.PasswordUpdatedAt.String()
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return err
	}
	//fmt.Println(tokenString)
	user.JwtToken = tokenString
	return nil
}

func ValidateJwt(db *gorm.DB, user *User) bool {
	var mySigningKey = []byte{69, 99, 77, 88, 55, 45, 45, 14, 14, 54, 77, 98, 98, 78, 11, 22, 33, 55}
	token, err := jwt.Parse(user.JwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return mySigningKey, nil
	})

	if err != nil {
		fmt.Println("something went wrong", err.Error())
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["email"].(string)
		dbUser, err := GetUserInfo(db, email)
		if err != nil {
			fmt.Println(err.Error())
			return false
		}
		//	jwtUpdatedAt := claims["updated_at"].(string)

		//fmt.Println(dbUser.PasswordUpdatedAt.String(), claims["updated_at"])
		//fmt.Println(dbUser.PasswordUpdatedAt.String())
		if dbUser.PasswordUpdatedAt.String() == claims["updated_at"] {
			*user = dbUser
			return true
		}
		//if claims["role"] == "admin" {
		//
		//	r.Header.Set("Role", "admin")
		//	handler.ServeHTTP(w, r)
		//	return
		//
		//}
	}

	fmt.Errorf("not auth")
	return false
}

func IsUserExist(db *gorm.DB, user *User) bool {
	var userCount int64
	if error := db.Model(&User{}).Where("email = ?", user.Email).Count(&userCount).Error; error != nil || userCount > 0 {
		//fmt.Println(error.Error())
		//fmt.Println(userCount, "hgghgh")
		return true
	}
	return false
}

func InsertUser(db *gorm.DB, user *User) bool {
	if err := db.Create(&user).Error; err != nil {
		fmt.Println(err.Error())
		return true
	}
	return false
}

func GetUserInfo(db *gorm.DB, email string) (User, error) {
	var user User
	if err := db.Where("email = ? ", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func UpdatePasswordChangeTime(db *gorm.DB, user *User) error {
	*user.PasswordUpdatedAt = time.Now()
	if err := db.Model(&user).Where("email =?", user.Email).Update("password_updated_at", user.PasswordUpdatedAt.String()).Error; err != nil {
		return err
	}
	dbUser, _ := GetUserInfo(db, user.Email)
	*user = dbUser
	return nil
}

func UpdatePassword(db *gorm.DB, user *User) error {

	if err := db.Model(&user).Where("email =?", user.Email).Update("password", user.Password).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
