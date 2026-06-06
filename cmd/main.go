package main

import (
	"midlewerego/internal/handler"
	"midlewerego/config"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func main () {
	route := gin.Default()
	db := config.DbConection()
	HandleRoute(route, db)
	route.Run(":3000")
}

func HandleRoute (route *gin.Engine, db *gorm.DB) {
	route.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status" : 200,
			"data" : "PONG", 
		})
	})

	route.POST("/api/user/new", handler.NewUser(db))
	route.GET("/api/user/data", handler.GetUser(db))
}