package models

import (
	"time"
)

type GroceryCategory struct {
	Id         int              `gorm:"primary_key;auto_increment" json:"id"`
	Name       string           `gorm:"size:255;not null;unique" json:"title"`
	CreatedAt  time.Time        `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time        `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt  *time.Time       `gorm:"default:null" json:"deleted_at"`
	CategoryId *int             `gorm:"default:null" json:"category_id"`
	Category   *GroceryCategory `gorm:"foreignkey:CategoryId"`
}
