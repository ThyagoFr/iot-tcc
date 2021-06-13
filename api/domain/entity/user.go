package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	ID         string `gorm:"type:uuid;primarykey"`
	Email      string
	Name       string
	Password   string
}