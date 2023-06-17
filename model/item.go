package model

import (
	"fmt"

	"github.com/gofrs/uuid"
)

type Item struct {
	UUID        uuid.UUID `json:"uuid" gorm:"primary_key;type:char(36)"`
	ID          string    `json:"id" gorm:"size:32;unique"`
	Description string    `json:"description" gorm:"size:86"`
	Point       int       `json:"point" validate:"omitempty"`
	Report      int       `json:"report" validate:"omitempty"`
}
type PublicItem struct {
	UUID        uuid.UUID `json:"uuid"`
	ID          string    `json:"id"`
	Description string    `json:"description" `
	Point       int       `json:"point"`
	Report      int       `json:"report"`
}

func GetItems() ([]*Item, error) {
	var items []*Item

	if err := db.Find(&items).Error; err != nil {

		return nil, err
	}

	return items, nil
}
func GetActiveItems() ([]*Item, error) {
	var items []*Item
	if err := db.Where("report <= ?", 2).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
func EnsureExisistenceID(id string) ([]*Item, error) {
	item := []*Item{}
	if err := db.Where("id = ?", id).Find(&item).Error; err != nil {
		return nil, err
	}
	fmt.Println(item)
	return item, nil
}

func AddItems(rawitem PublicItem) (*Item, error) {
	newuuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	item := Item{
		UUID:        newuuid,
		ID:          rawitem.ID,
		Description: rawitem.Description,
		Point:       rawitem.Point,
		Report:      rawitem.Report,
	}

	err = db.Create(&item).Error
	if err != nil {

		return nil, err
	}
	return &item, nil
}
