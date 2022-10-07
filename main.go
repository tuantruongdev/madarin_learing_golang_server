package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"mandarinLearningBE/pkg/routes"
)

const (
	dsn = "root:@tcp(127.0.0.1:3306)/developer?charset=utf8mb4&parseTime=True&loc=Local"
)

func main() {
	//fmt.Println(gin.H{"time": time.Now().Format("2022-10-07 14:53:02 +0700 +07")})
	//
	//return

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}
	log.Println("Connected to MySQL:", db)
	router := gin.Default()
	routes.Register(router, db)

	router.Run()
}
