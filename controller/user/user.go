package user

import (
	"myproject/orm"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewAllUser(ctx *gin.Context) {
	var user []orm.User
	orm.Db.Find(&user)
	ctx.JSON(http.StatusOK, gin.H{
		"user":   user,
		"status": "ok",
	})

}
func SearchIndex(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(float64)
	var user []orm.User
	orm.Db.First(&user, userId)
	ctx.JSON(http.StatusOK, gin.H{
		"user":   user,
		"status": "ok",
	})

}
