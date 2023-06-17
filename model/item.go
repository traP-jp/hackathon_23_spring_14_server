package model

import (
	"log"

	"github.com/gofrs/uuid"
)

type Items struct {
	UUID        uuid.UUID `json:"uuid" gorm:"primary_key;type:char(36)"`
	ID          string    `json:"id" gorm:"size:32;unique"`
	Description string    `json:"description" gorm:"size:86;unique"`
	Point       int       `json:"point" validate:"omitempty"`
	Report      int       `json:"report" validate:"omitempty"`
}

func GetItems(include bool) []Items {
	var items []Items
	var err error
	if include {
		err = db.Find(&items).Error
	} else {
		err = db.Where("report <= ?", 2).Find(&items).Error
	}
	if err != nil {
		log.Fatal(err)
	}
	return items
}

func AddItems(item Items) {
	err := db.Create(item).Error //
	if err != nil {
		log.Fatal(err)
	}
}
