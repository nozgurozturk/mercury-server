package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Project Model
type Project struct {
	gorm.Model
	ID       string `json:"id"`
	Name     string `json:"description"`
	Link     string `json:"link"`
	TestLink string `json:"testlink"`
	UserID   string `json:"user_id"`
}
