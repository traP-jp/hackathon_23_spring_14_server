package model

import "time"

type User struct {
	UUID  string    `json:"uuid" gorm:"size:32;primary_key"`
	ID    string    `json:"id" gorm:"size:32;unique"`
	Point int       `json:"point" validate:"omitempty"`
	Date  time.Time `json:"date" gorm:"primary_key"`
}

type TimeCards struct {
	ID     string    `json:"id" gorm:"size:32;primary_key"`
	Date   time.Time `json:"date" gorm:"primary_key"`
	ItemID string    `json:"itemID" gorm:"size:86"`
}

func GetUsers() ([]*User, error) {
	var users []*User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
