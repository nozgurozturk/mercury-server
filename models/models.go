package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

var initMigration = func() {
	// Load .env file
	// e := godotenv.Load()
	// if e != nil {
	// 	fmt.Print(e)
	// }
	// // db Connection
	// host := os.Getenv("db_host")
	// user := os.Getenv("db_user")
	// pass := os.Getenv("db_pass")
	// dbname := os.Getenv("db_name")

	// psqlInfo := fmt.Sprint("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, pass, dbname)
	// fmt.Print(psqlInfo)
	db, err = gorm.Open("postgres", "host=localhost user=postgres password=Postgres.69487 dbname=startpage sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.Debug().AutoMigrate(&User{}, &Page{}, &Project{}, &Guide{})

}

var GetDB = func() *gorm.DB {
	return db
}
