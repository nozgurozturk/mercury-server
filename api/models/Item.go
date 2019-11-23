package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Item struct {
	ID          uint32    `gorm:"primary_key;" json:"id"`
	BoardID     uint32    `gorm:"index;not_null;" json:"board_id"`
	Name        string    `gorm:"type:varchar(40);not_null;" json:"name"`
	OrderNumber uint32    `gorm:"type:smallint;" json:"order_number"`
	Links       []Link    `gorm:"foreignkey:ItemID:association_foreignkey:ID"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (i *Item) Validate() error {

	if i.Name == "" {
		return errors.New("required name")
	}
	if i.BoardID < 1 {
		return errors.New("required board")
	}
	return nil
}

func (i *Item)BeforeSaveItem(db *gorm.DB) uint32{
	var count uint32
	db.Model(&Item{}).Where("board_id = ?", &i.BoardID).Count(&count)
	fmt.Println(count)
	return count+1
}


func (i *Item) SaveItem(db *gorm.DB) (*Item, error) {
	var err error
	err = db.Debug().Model(&Item{}).Create(&i).Error
	if err != nil {
		return &Item{}, err
	}
	return i, nil
}

func (i *Item) FindAllItem(db *gorm.DB, bid uint32) (*[]Item, error) {
	var err error
	var items []Item
	err = db.Debug().Model(&Board{ID: bid}).Related(&items).Error
	if err != nil {
		return &items, err
	}
	return &items, err
}
func (i *Item) FindItem(db *gorm.DB, iid uint32) (*Item, error) {
	var err error
	err = db.Preload("Links").First(&i, iid).Error
	//err = db.Debug().Model(Item{}).Where("id = ?", iid).First(&i).Error
	if err != nil {
		return &Item{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Item{}, errors.New("item not found")
	}
	return i, nil
}

func (i *Item) UpdateItem(db *gorm.DB, iid uint32) (*Item, error) {
	var err error

	db = db.Debug().Model(Item{}).Where("id = ?", iid).First(&Item{}).UpdateColumns(
		map[string]interface{}{
			"name":      i.Name,
			"order_number": i.OrderNumber,
			"update_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Item{}).Where("id = ?", iid).First(&i).Error
	if err != nil {
		return &Item{}, db.Error
	}

	return i, nil
}

func (i *Item) DeleteItem(db *gorm.DB, iid uint32) (int64, error) {

	db = db.Debug().Model(&Item{}).Where("id = ? ", iid).First(&i).Delete(&Item{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("item not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
