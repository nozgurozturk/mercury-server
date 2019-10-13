package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
	"time"
)

type User struct {

	ID        uint32    `gorm:"primary_key;auto_increment;" json:"id"`
	Name      string    `gorm:"type:varchar(40);not_null;" json:"name"`
	Email     string    `gorm:"type:varchar(120);not_null;unique;" json:"email"`
	Password  string    `gorm:"type:varchar(60);not_null;" json:"password"`
	CreatedAt time.Time `gorm:default:"CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:default:"CURRENT_TIMESTAMP" json:"updated_at"`
}

func HashPass(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPass(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := HashPass(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Name == "" {
			return errors.New("required name")
		}
		if u.Password == "" {
			return errors.New("required password")
		}
		if u.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("required password")
		}
		if u.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil

	default:
		if u.Name == "" {
			return errors.New("required name")
		}
		if u.Password == "" {
			return errors.New("required password")
		}
		if u.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil
	}
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) FindUser(db *gorm.DB, uid uint32)(*User, error){
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).First(&u).Error
	if err != nil{
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err){
		return &User{}, errors.New("User Not Found")
	}
	return u, nil
}

func (u *User) UpdateUser(db *gorm.DB, uid uint32)(*User, error){
	var err error
	err = u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}

	db = db.Debug().Model(User{}).Where("id = ?", uid).First(&u).UpdateColumns(
		map[string]interface{}{
			"password" : u.Password,
			"name" : u.Name,
			"email" : u.Email,
			"update_at" : time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	err = db.Debug().Model(User{}).Where("id = ?", uid).First(&u).Error
	if err != nil{
		return &User{}, err
	}
	return u, nil
}

func (u *User) DeleteUser(db *gorm.DB, uid uint32)(int64, error){

	db = db.Debug().Model(User{}).Where("id = ?", uid).First(&u).Delete(&User{})
	if db.Error != nil{
		return 0, db.Error
	}
	return db.RowsAffected, nil
}