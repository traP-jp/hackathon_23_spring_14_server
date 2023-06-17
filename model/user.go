package model

import "time"

type User struct {
	ID    string    `json:"id" gorm:"size:32;primary_key"`
	Point int       `json:"point" validate:"omitempty"`
	Date  time.Time `json:"date" validate:"omitempty"`
}

type TimeCards struct {
	ID     string    `json:"id" gorm:"size:32;primary_key"`
	Date   time.Time `json:"date" gorm:"primary_key"`
	ItemID string    `json:"itemID" gorm:"size:86"`
}
