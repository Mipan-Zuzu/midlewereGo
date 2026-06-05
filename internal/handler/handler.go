package handler

import (
	"midlewerego/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	// "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewUser (db *gorm.DB) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		var user *model.Model
		ctx.ShouldBindJSON(&user)
		ctx.JSON(http.StatusOK, gin.H{
			"data" : user,
		})
	}
}

func GetUser () {

}