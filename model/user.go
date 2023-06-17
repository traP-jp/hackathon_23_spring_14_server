package model

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	UUID  uuid.UUID `json:"uuid" gorm:"primary_key;type:char(36)"`
	ID    string    `json:"id" gorm:"size:32;unique"`
	Point int       `json:"point" validate:"omitempty"`
	Date  time.Time `json:"date" gorm:"primary_key"`
}

type TimeCards struct {
	ID     string    `json:"id" gorm:"size:32;primary_key"`
	Date   time.Time `json:"date" gorm:"primary_key"`
	ItemID string    `json:"itemID" gorm:"size:86"`
}

type DataSet struct {
	Point int       `json:"point"`
	Date  time.Time `json:"date"`
}

type PublicUser struct {
	UUID    uuid.UUID  `json:"uuid"`
	ID      string     `json:"id"`
	DataSet []*DataSet `json:"dataset"`
}
type ID struct {
	UUID uuid.UUID `json:"uuid"`
	ID   string    `json:"id"`
}

func GetUsers() ([]*PublicUser, error) {
	ids := []*ID{}
	users := []*PublicUser{}

	tx := db.Begin()
	if err := db.Model(&User{}).Select("uuid", "id").Group("uuid").Find(&ids).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, id := range ids {
		data := []*DataSet{}
		if err := db.Model(&User{}).Select("point", "date").Where("uuid = ?", id.UUID).Find(&data).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
		users = append(users, &PublicUser{
			UUID:    id.UUID,
			ID:      id.ID,
			DataSet: data,
		})
	}

	return users, tx.Commit().Error
}
