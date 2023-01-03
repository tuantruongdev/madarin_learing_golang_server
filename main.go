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
	//res := models.NetResponse{}.Build()
	//res.SetStatus(200, "ok", "mess")
	//res.H["data"] = []int{10, 5, 1, 1, 4}
	//fmt.Println(res)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}
	log.Println("Connected to MySQL:", db)
	router := gin.Default()
	routes.Register(router, db)

	router.Run()
}
