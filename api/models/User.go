package models

import (
	"errors"
	"html"
	"net/mail"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	UserID        int        `gorm:"primary_key;auto_increment" json:"id"`
	UserFirstName string     `gorm:"size:255;not null;" json:"firstname"`
	UserLastName  string     `gorm:"size:255;not null;" json:"lastname"`
	UserEmail     string     `gorm:"size:100;not null;unique" json:"email"`
	Password      string     `gorm:"size:100;not null;" json:"password"`
	CreatedAt     time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt     *time.Time `gorm:"default:null" json:"deleted_at"`
	// Group         Group     `gorm:"foreignkey:GroupID"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (u *User) Setup() {
	u.UserID = 0
	u.UserFirstName = html.EscapeString(strings.TrimSpace(u.UserFirstName))
	u.UserLastName = html.EscapeString(strings.TrimSpace(u.UserLastName))
	u.UserEmail = html.EscapeString(strings.TrimSpace(u.UserEmail))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	u.DeletedAt = nil
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	// TODO: add validation stuff
	// case "update":

	// 	return nil

	default:
		if u.UserFirstName == "" {
			return errors.New("Required User First Name")
		}
		if u.UserLastName == "" {
			return errors.New("Required User Last Name")
		}
		if u.Password == "" {
			return errors.New("Required User Password")
		}
		if u.UserEmail == "" {
			return errors.New("Required User Email")
		}
		if _, err := mail.ParseAddress(u.UserEmail); err != nil {
			return errors.New("Invalid Email Address")
		}
		return nil
	}
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) SecurePassword() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
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

func (u *User) SaveUser(db *gorm.DB) (*User, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}
