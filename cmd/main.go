package main

import (
	"midlewerego/config"
	"midlewerego/internal/handler"
	"midlewerego/midlewere"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	route := gin.Default()
	midlewere.CorsMidllewere(route)
	db := config.DbConection()
	HandleRoute(route, db)
	route.Run(":3000")
}

func HandleRoute(route *gin.Engine, db *gorm.DB) {
	route.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": 200,
			"data":   "PONG",
		})
	})

	route.POST("/api/user/new", handler.NewUser(db))
	route.POST("/api/user/login", handler.Login(db))
}
