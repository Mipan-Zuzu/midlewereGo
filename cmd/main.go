package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main () {
	godotenv.Load()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", 
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, erro := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if erro != nil {
		panic("Cannot conec to database")
	}
	route := gin.Default()
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
}