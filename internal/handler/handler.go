package handler

import (
	"midlewerego/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
	// "gorm.io/driver/postgres"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate = validator.New()

func NewUser (db *gorm.DB) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		var user *model.Models
		ctx.ShouldBindJSON(&user)
		err := validate.Struct(user)
		if err != nil {
			ctx.JSON(400, gin.H{
				"status" : 400,
				"data" : err.Error(),
			})
			return
		}
		db.Create(user)
		ctx.JSON(http.StatusOK, gin.H{
			"data" : user,
		})
	}
}

func GetUser (db *gorm.DB) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		var user []model.Models
		db.Find(&user)
		if len(user) == 0 {
			ctx.JSON(400, gin.H{
				"status" : 400,
				"data" : "Empty user data",
				})			
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status" : 200,
			"data" : user,
		})
	}
}