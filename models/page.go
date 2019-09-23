package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// User Model
type Page struct {
	gorm.Model
	ID     string
	Name   string
	Link   string
	UserID string
}
