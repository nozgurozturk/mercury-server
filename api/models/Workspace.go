package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type Workspace struct {
	ID        uint32    `gorm:"primary_key;auto_increment;" json:"id"`
	Name      string    `gorm:"type:varchar(40);not_null;" json:"name"`
	Users     []User    `gorm:"many2many:user_workspaces;" json:"users"`
	Boards    []Board   `gorm:"foreignkey:WorkspaceID;association_foreignkey:ID" json:"boards"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (ws *Workspace) Validate() error {

	if ws.Name == "" {
		return errors.New("required name")
	}
	return nil
}

func (ws *Workspace) SaveWorkspace(db *gorm.DB, uid uint32) (*Workspace, error) {

	var err error
	err = db.Debug().Model(&Workspace{}).Create(&ws).Association("Users").Append(User{ID:uid}).Error
	if err != nil {
		return &Workspace{}, err
	}
	return ws, nil
}
func (ws *Workspace) FindAllWorkspace(db *gorm.DB, uid uint32) (*[]Workspace, error) {
	var err error
	var workspaces []Workspace
	err = db.Preload("Users").Preload("Boards").Preload("Items").Model(&User{ID:uid}).Association("Workspaces").Find(&workspaces).Error
	//err = db.Debug().Preload("Users").Model(&Workspace{}).Find(&workspaces).Error
	if err != nil {
		return &[]Workspace{}, err
	}
	return &workspaces, err
}

func (ws *Workspace) FindWorkspace(db *gorm.DB, wid uint32) (*Workspace, error) {
	var err error
	err = db.Debug().Preload("Boards").Model(Workspace{}).First(&ws, wid).Error
	if err != nil {
		return &Workspace{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Workspace{}, errors.New("user Not Found")
	}
	return ws, nil
}

func (ws *Workspace) UpdateWorkspace(db *gorm.DB, wid uint32) (*Workspace, error) {
	var err error

	db = db.Debug().Model(Workspace{}).Where("id = ?", wid).First(&ws).UpdateColumns(
		map[string]interface{}{
			"name":  ws.Name,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Workspace{}, db.Error
	}
	err = db.Debug().Model(Workspace{}).Where("id = ?", wid).First(&ws).Error
	if err != nil {
		return &Workspace{}, err
	}
	return ws, nil
}

func (ws *Workspace) DeleteWorkspace(db *gorm.DB, wid uint32) (int64, error) {

	db = db.Debug().Model(Workspace{}).Where("id = ?", wid).First(&ws).Delete(&Workspace{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
