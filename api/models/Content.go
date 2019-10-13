package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type Content struct {
	gorm.Model
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Owner     User      `json:"owner"`
	OwnerId   uint32    `gorm:"not null" json:"owner_id"`
	Title     string    `gorm:"type:varchar(40);not_null;" json:"name"`
	Link      string    `gorm:"type:varchar(60);not_null;unique;" json:"email"`
	TestLink  string    `gorm:"type:varchar(70);" json:"password"`
	Type      string    `gorm:"type:varchar(70);not_null;" json:"type"`
	CreatedAt time.Time `gorm:default:"CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:default:"CURRENT_TIMESTAMP" json:"updated_at"`
}

func (c *Content) Validate() error{

	if c.Link == "" {
		return  errors.New("link is required")
	}
	if c.Title == ""{
		return errors.New("title is required")
	}
	if c.Type == ""{
		return errors.New("type is required")
	}
	if c.OwnerId < 1 {
		return errors.New("invalid")
	}
	return  nil
}


func (c *Content) SaveContent(db *gorm.DB) (*Content, error) {
	var err error
	err = db.Debug().Model(&Content{}).Create(&c).Error
	if err != nil {
		return &Content{}, err
	}
	if c.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", c.OwnerId).Take(&c.Owner).Error
		if err != nil {
			return &Content{}, err
		}
	}
	return c, nil
}

func (c *Content) GetAllContent(db *gorm.DB) (*[]Content, error) {
	var err error
	contents := []Content{}
	err = db.Debug().Model(&Content{}).Limit(50).Find(&contents).Error
	if err != nil {
		return &[]Content{}, err
	}
	if len(contents) > 0 {
		for i, _ := range contents {
			err = db.Debug().Model(&User{}).Where("id = ?", contents[i].OwnerId).Find(&contents).Error
			if err != nil {
				return &[]Content{}, err
			}
		}
	}
	return &contents, err
}
func (c *Content) FindContent(db *gorm.DB, cid uint32) (*Content, error) {
	var err error
	err = db.Debug().Model(Content{}).Where("id = ?", cid).First(&c).Error
	if err != nil {
		return &Content{}, err
	}
	if c.ID != 0 {
		err = db.Debug().Model(User{}).Where("id = ?", c.OwnerId).First(&c.Owner).Error
		if err != nil {
			return &Content{}, err
		}
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Content{}, errors.New("Content Not Found")
	}
	return c, nil
}

func (c *Content) UpdateContent(db *gorm.DB, cid uint32) (*Content, error) {
	var err error

	db = db.Debug().Model(Content{}).Where("id = ?", cid).First(&Content{}).UpdateColumns(
		map[string]interface{}{
			"title":     c.Title,
			"link":      c.Link,
			"test_link": c.TestLink,
			"type":      c.Type,
			"update_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Content{}).Where("id = ?", cid).First(&c).Error
	if db.Error != nil {
		return &Content{}, db.Error
	}
	if c.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", c.OwnerId).Find(&c.Owner).Error
		if err != nil {

			return &Content{}, err
		}
	}

	return c, nil
}

func (c *Content) DeleteContent(db *gorm.DB, cid uint32, uid uint32) (int64, error) {

	db = db.Debug().Model(&Content{}).Where("id = ? and owner_id = ?", cid, uid).First(&c).Delete(&Content{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Content Not Found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
