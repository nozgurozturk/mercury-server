package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Guide Model
type Guide struct {
	gorm.Model
	ID     string `json:"id"`
	Name   string `json:"description"`
	Link   string `json:"link"`
	UserID string `json:"user_id"`
}
