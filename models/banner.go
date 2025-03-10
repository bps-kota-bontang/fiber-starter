package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string `gorm:"not null"`
	gorm.Model
}
