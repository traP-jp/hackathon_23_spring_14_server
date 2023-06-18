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

	if err := db.Raw("SELECT * FROM `items`").Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}
func GetActiveItems() ([]*Item, error) {
	var items []*Item
	if err := db.Raw("SELECT * FROM `items` WHERE report <= 2").Scan(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
func EnsureExisistenceID(id string) ([]*Item, error) {
	item := []*Item{}
	if err := db.Raw("SELECT * FROM `items` WHERE id = ?", id).Scan(&item).Error; err != nil {
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

	err = db.Table("items").Create(&item).Error
	if err != nil {

		return nil, err
	}
	return &item, nil
}

func ReportItem(id string) (*Item, error) {
	item := Item{}
	if err := db.Raw("SELECT * FROM `items` WHERE id = ?", id).Scan(&item).Error; err != nil {
		return nil, err
	}
	item.Report += 1
	if err := db.Exec("UPDATE `items` SET report = ? WHERE uuid = ?", item.Report, item.UUID).Error; err != nil {
		return nil, err
	}
	return &item, nil
}
