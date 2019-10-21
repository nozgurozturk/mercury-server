package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type Link struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Name	  string    `gorm:"type:varchar(10);not_null;" json:"name"`
	URL       string    `gorm:"type:varchar(120);not_null:" json:"url"`
	ItemID    uint32    `gorm:"not_null" json:"item_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (l *Link) Validate() error {

	if l.Name == "" {
		return errors.New("required name")
	}
	if l.ItemID < 1 {
		return errors.New("required item")
	}
	return nil
}

func (l *Link) SaveLink(db *gorm.DB) (*Link, error) {
	var err error
	err = db.Debug().Model(&Link{}).Create(&l).Error
	if err != nil {
		return &Link{}, err
	}
	return l, nil
}

func (l *Link) FindAllLink(db *gorm.DB) (*[]Link, error) {
	var err error
	var links []Link
	err = db.Debug().Model(&Link{}).Find(&links).Error
	if err != nil {
		return &links, err
	}
	return &links, err
}
func (l *Link) FindLink(db *gorm.DB, lid uint32) (*Link, error) {
	var err error
	err = db.Debug().Model(Link{}).Where("id = ?", lid).First(&l).Error
	if err != nil {
		return &Link{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Link{}, errors.New("link not found")
	}
	return l, nil
}

func (l *Link) UpdateLink(db *gorm.DB, lid uint32) (*Link, error) {
	var err error

	db = db.Debug().Model(Link{}).Where("id = ?", lid).First(&Link{}).UpdateColumns(
		map[string]interface{}{
			"name":     l.Name,
			"update_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Link{}).Where("id = ?", lid).First(&l).Error
	if err != nil {
		return &Link{}, db.Error
	}

	return l, nil
}

func (l *Link) DeleteLink(db *gorm.DB, lid uint32) (int64, error) {

	db = db.Debug().Model(&Link{}).Where("id = ? ", lid).First(&l).Delete(&Link{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("link not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}