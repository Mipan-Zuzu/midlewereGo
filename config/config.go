package config

import (
	"fmt"
	"midlewerego/internal/model"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



func DbConection () *gorm.DB{
	godotenv.Load()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable search_path=public", 
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, erro := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if erro != nil {
		panic("Cannot conect to database")
	}
	if err := db.AutoMigrate(&model.Models{}); err != nil {
		panic(err.Error())
	}
	return db
}