package model

import (
	"gorm.io/gorm"
)

type Models struct {
	gorm.Model
	Email    string `validate:"required,email" gorm:"unique"`
	Password string `validate:"required,min=8"`
	Username string `validate:"required" gorm:"unique"`
}
