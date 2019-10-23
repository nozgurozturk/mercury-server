package models

import (
	. "errors"
	. "github.com/jinzhu/gorm"
	. "time"
)

type Board struct {
	ID        uint32 `gorm:"primary_key;auto_increment" json:"id"`
	UserID    uint32 `gorm:"index;" json:"user_id"`
	Name      string `gorm:"type:varchar(40);not_null;" json:"name"`
	Items     []Item  `gorm:"foreignkey:BoardID;association_foreignkey:ID" json:"items"`
	CreatedAt Time   `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt Time   `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (b *Board) Validate() error {

	if b.Name == "" {
		return New("required name")
	}
	if b.UserID < 1 {
		return New("required owner")
	}
	return nil
}

func (b *Board) SaveBoard(db *DB) (*Board, error) {
	var err error
	err = db.Debug().Model(&Board{}).Create(&b).Error
	if err != nil {
		return &Board{}, err
	}
	return b, nil
}

func (b *Board) FindAllBoard(db *DB, uid uint32) (*[]Board, error) {
	var err error
	var boards []Board
	err = db.Debug().Preload("Items").Model(&User{ID:uid}).Related(&boards).Error
	if err != nil {
		return &[]Board{}, err
	}
	return &boards, err
}

func (b *Board) FindBoard(db *DB, bid uint32) (*Board, error) {
	var err error
	err = db.Preload("Items").First(&b, bid).Error
	if err != nil {
		return &Board{}, err
	}
	if IsRecordNotFoundError(err) {
		return &Board{}, New("board Not Found")
	}
	return b, nil
}

func (b *Board) UpdateBoard(db *DB, bid uint32) (*Board, error) {
	var err error

	db = db.Debug().Model(Board{}).Where("id = ?", bid).First(&Board{}).UpdateColumns(
		map[string]interface{}{
			"name":      b.Name,
			"update_at": Now(),
		},
	)
	err = db.Debug().Model(&Board{}).Where("id = ?", bid).First(&b).Error
	if err != nil {
		return &Board{}, db.Error
	}

	return b, nil
}

func (b *Board) DeleteBoard(db *DB, bid uint32) (int64, error) {

	db = db.Debug().Model(&Board{}).Where("id = ? ", bid).First(&b).Delete(&Board{})
	if db.Error != nil {
		if IsRecordNotFoundError(db.Error) {
			return 0, New("board not Found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
