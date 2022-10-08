package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Group struct {
	GroupID   int       `gorm:"primary_key;auto_increment" json:"id"`
	GroupName string    `gorm:"size:255;not null;" json:"firstname"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (group *Group) FindGroupByID(db *gorm.DB, gid uint32) (*Group, error) {
	var err error
	err = db.Debug().Model(Group{}).Where("id = ?", gid).Take(&group).Error
	if err != nil {
		return &Group{}, err
	}
	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return &Group{}, errors.New("Group Not Found")
	}
	return group, err
}
