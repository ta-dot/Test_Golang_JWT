package main

import (
	"log"
	"myproject/controller/auth"
	"myproject/controller/user"
	"myproject/middleware"
	"myproject/orm"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	orm.InitDB()
	r.POST("/register", auth.Register)
	r.POST("/login", auth.Login)
	authorized := r.Group("/user", middleware.JWTAuthen())
	authorized.GET("/readall", user.ViewAllUser)
	authorized.GET("/find", user.SearchIndex)

	r.Run(":7070")
}
