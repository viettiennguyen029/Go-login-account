package models

import "github.com/jinzhu/gorm"

type Token struct {
	gorm.Model
	Token string `gorm:"primaryKey" json:"token"`
}

func (t *Token) SaveNewToken(db *gorm.DB) (*Token, error) {
	err := db.Debug().Create(&t).Error
	if err != nil {
		return &Token{}, err
	}
	return t, nil
}
