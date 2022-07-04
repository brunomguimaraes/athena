package models

import (
	"gorm.io/gorm"
)

type User struct {
	UserID        int    `gorm:"primary_key;auto_increment" json:"id"`
	UserFirstName string `gorm:"size:255;not null;" json:"firstname"`
	UserLastName  string `gorm:"size:255;not null;" json:"lastname"`
	UserEmail     string `gorm:"size:100;not null;unique" json:"email"`
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error

	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}
