package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Grocery struct {
	Id          int        `gorm:"primary_key;auto_increment" json:"id"`
	Name        string     `gorm:"size:255;not null;unique" json:"title"`
	Description string     `gorm:"size:255;" json:"description"`
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"default:null" json:"deleted_at"`
}

func (grocery *Grocery) Setup() {
	grocery.Id = 0
	grocery.Name = html.EscapeString(strings.TrimSpace(grocery.Name))
	grocery.Description = html.EscapeString(strings.TrimSpace(grocery.Description))
	grocery.CreatedAt = time.Now()
	grocery.UpdatedAt = time.Now()
	grocery.DeletedAt = nil
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

func (grocery *Grocery) FindAllGroceries(db *gorm.DB) (*[]Grocery, error) {
	var err error

	groceries := []Grocery{}
	err = db.Debug().Model(&Grocery{}).Limit(100).Find(&groceries).Error
	if err != nil {
		return &[]Grocery{}, err
	}
	return &groceries, err
}
