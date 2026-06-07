package handler

import (
	"midlewerego/internal/model"
	"midlewerego/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	// "gorm.io/driver/postgres"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate = validator.New()

func NewUser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user *model.Models
		ctx.ShouldBindJSON(&user)
		err := validate.Struct(user)
		if err != nil {
			ctx.JSON(400, gin.H{"status": 400,"data":   err.Error(),})
			return
		}
		hased, err := utils.HashPassword(user.Password)
		user.Password = hased
		if err != nil {
			ctx.JSON(400, gin.H{"status": 400, "data": "failed to hash"})
		}
		if err := db.Create(user).Error; err != nil {
			if strings.Contains(err.Error(), "email") {
				ctx.JSON(400, gin.H{"status" : 400, "data" : "email alredy exist"})
				return
			}
			if strings.Contains(err.Error(), "username") {
				ctx.JSON(400, gin.H{"status": 400, "data" : "username alredy taken"})
				return 
			}
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status": 200,
			"data":   "succses sign user",
		})
	}
}

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user *model.Models
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(400, gin.H{"status": 400, "data": err.Error()})
			return
		}
		var findUser *model.Models
		if err := db.Where("email = ?", user.Email).First(&findUser).Error; err != nil {
			ctx.JSON(404, gin.H{"status": 404, "data": "user not found"})
			return
		}
		if err := utils.ComperePassword(findUser.Password, user.Password); err != nil {
			ctx.JSON(401, gin.H{"status" : 401, "data" : "wrong password"})
			return
		}

		if user.Username != findUser.Username {
			ctx.JSON(401, gin.H{"satatus": 401, "data" : "wrong username"})
			return
		}
		jwtToken,err := utils.GenerateToken(user.ID)
		if err != nil {
			ctx.JSON(400, gin.H{"status" : 400, "data" : "failed to create jwt token"})
			return
		}
		ctx.Header("Authorization", "Barier" + jwtToken)
		ctx.JSON(http.StatusOK, gin.H{
			"status": 200,
			"data":   "succsesfull login",
		})
	}
}

