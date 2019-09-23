package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// User Model
type User struct {
	gorm.Model
	ID       string
	Name     string
	Password string
	Email    string
	Page     []Page
	Project  []Project
	Guide    []Guide
}
