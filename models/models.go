package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB
var err error

var InitMigration = func() {
	// Load .env file
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	// db Connection
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER_NAME")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

	connection, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	db = connection
	db.Debug().AutoMigrate(&User{}, &Page{}, &Project{}, &Guide{})
}

var GetDB = func() *gorm.DB {
	return db
}
