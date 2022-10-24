package models

import (
	"errors"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type GroceryList struct {
	Id         int            `gorm:"primary_key;auto_increment" json:"id"`
	GroceryIds pq.StringArray `gorm:"type:text[]"`
	CreatedAt  time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt  *time.Time     `gorm:"default:null" json:"deleted_at"`
	GroupId    *int           `gorm:"default:null" json:"group_id"`
	Group      *Group         `gorm:"foreignkey:GroupId"`
}

func (g *GroceryList) FindGroceryListByID(db *gorm.DB, uid uint32) (*GroceryList, error) {
	var err error
	err = db.Debug().Model(GroceryList{}).Where("id = ?", uid).Take(&g).Error
	if err != nil {
		return &GroceryList{}, err
	}
	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return &GroceryList{}, errors.New("User Not Found")
	}
	return g, err
}

// func (grocery *Grocery) Setup() {
// 	grocery.Id = 0
// 	grocery.Name = html.EscapeString(strings.TrimSpace(grocery.Name))
// 	grocery.CreatedAt = time.Now()
// 	grocery.UpdatedAt = time.Now()
// }

// func (grocery *Grocery) Validate() error {

// 	if grocery.Name == "" {
// 		return errors.New("Required Grocery Name")
// 	}

// 	return nil
// }

// func (grocery *Grocery) SaveGrocery(db *gorm.DB) (*Grocery, error) {
// 	var err error
// 	err = db.Debug().Model(&Grocery{}).Create(&grocery).Error
// 	if err != nil {
// 		return &Grocery{}, err
// 	}
// 	return grocery, nil
// }

// func (g *Group) FindAllGroceries(db *gorm.DB) (*[]Grocery, error) {
// 	var err error

// 	groceries := []Grocery{}
// 	err = db.Debug().Model(&Grocery{}).Limit(100).Find(&groceries).Error
// 	if err != nil {
// 		return &[]Grocery{}, err
// 	}
// 	return &groceries, err
// }
