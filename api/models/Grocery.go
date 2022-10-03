package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Grocery struct {
	Id        int       `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"title"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (grocery *Grocery) Setup() {
	grocery.Id = 0
	grocery.Name = html.EscapeString(strings.TrimSpace(grocery.Name))
	grocery.CreatedAt = time.Now()
	grocery.UpdatedAt = time.Now()
}

func (grocery *Grocery) Validate() error {

	if grocery.Name == "" {
		return errors.New("Required Grocery Name")
	}

	return nil
}

func (grocery *Grocery) SaveGrocery(db *gorm.DB) (*Grocery, error) {
	var err error
	err = db.Debug().Model(&Grocery{}).Create(&grocery).Error
	if err != nil {
		return &Grocery{}, err
	}
	return grocery, nil
}
