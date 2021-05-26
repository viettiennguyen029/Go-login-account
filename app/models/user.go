package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Address     string    `gorm:"primaryKey" json:"address"`
	ImageUrl    string    `json:"image_url"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
}

func (u *User) SaveNewUser(db *gorm.DB) (*User, error) {
	err := db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindUserByAddress(db *gorm.DB, addr string) (*User, error) {
	err := db.Debug().Model(&User{}).Where("address = ?", addr).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) UpdateUser(db *gorm.DB, addr string) (*User, error) {
	db = db.Debug().Model(&User{}).Where("address = ?", addr).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"image_url":    u.ImageUrl,
			"updated_time": time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	err := db.Debug().Model(&User{}).Where("address = ?", addr).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
