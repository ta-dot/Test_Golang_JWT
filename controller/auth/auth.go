package auth

import (
	"fmt"
	"myproject/orm"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type RegisterBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
}
type LoginBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var hmacSampleSecret []byte

func Register(ctx *gin.Context) {
	var register RegisterBody
	err := ctx.ShouldBindJSON(&register)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var existingUser orm.User
	result := orm.Db.Where("username = ?", register.Username).First(&existingUser)
	if result.Error == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}
	ency, _ := bcrypt.GenerateFromPassword([]byte(register.Password), 10)
	user := orm.User{Username: register.Username, Password: string(ency), Fullname: register.Fullname}

	orm.Db.Create(&user)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Register successfully",
	})

}
func Login(ctx *gin.Context) {
	var login LoginBody
	err := ctx.ShouldBindJSON(&login)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var loginmatch orm.User
	orm.Db.Where("username = ? ", login.Username).First(&loginmatch)
	if loginmatch.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username or password Invalid"})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(loginmatch.Password), []byte(login.Password))
	if err == nil {
		hmacSampleSecret = []byte(os.Getenv("JWT_SECRET_KEY"))
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userId": loginmatch.ID,
			"exp":    time.Now().Add(time.Minute * 1).Unix(),
		})
		tokenString, err := token.SignedString(hmacSampleSecret)
		fmt.Println(tokenString, err)
		ctx.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Login Success", "token": tokenString})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username or password Invalid"})
		return  
	}

}
