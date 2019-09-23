package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// User Model
type Page struct {
	gorm.Model
	id     uint `gorm:"primary_key"`
	Name   string
	Link   string
	UserID uint
}
