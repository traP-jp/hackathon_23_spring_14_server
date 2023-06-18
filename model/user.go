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

type DataSet struct {
	Point int       `json:"point"`
	Date  time.Time `json:"date"`
}

type DataSetDetail struct {
	Score    int       `json:"score"`
	Date     time.Time `json:"date"`
	ItemList []*Item   `json:"itemList"`
}

type PublicUser struct {
	UUID    uuid.UUID  `json:"uuid"`
	ID      string     `json:"id"`
	DataSet []*DataSet `json:"dataset"`
}

type UserDetail struct {
	UUID    uuid.UUID        `json:"uuid"`
	ID      string           `json:"id"`
	DataSet []*DataSetDetail `json:"dataset"`
}
type ID struct {
	UUID uuid.UUID `json:"uuid"`
	ID   string    `json:"id"`
}
type Ranking struct {
	UUID  uuid.UUID `json:"uuid" gorm:"primary_key;type:char(36)"`
	ID    string    `json:"id" gorm:"size:32;unique"`
	Score int       `json:"score" validate:"omitempty"`
}

func GetUsers() ([]*PublicUser, error) {
	ids := []*ID{}
	users := []*PublicUser{}

	tx := db.Begin()
	if err := db.Raw("SELECT `uuid`,`id` FROM `users` GROUP BY `uuid`").Scan(&ids).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for _, id := range ids {
		data := []*DataSet{}
		if err := db.Raw("SELECT point,date FROM `users` WHERE uuid = ? GROUP BY `uuid`", id.UUID).Scan(&data).Error; err != nil {
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

func GetRanking() ([]*Ranking, error) {
	ranks := []*Ranking{}

	tx := db.Begin()
	if err := db.Raw("SELECT `uuid`,`id`,sum(point) FROM `users` GROUP BY `uuid` ORDER BY `point` DESC").Scan(&ranks).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return ranks, tx.Commit().Error
}

func GetMe() ([]*UserDetail, error) {
	ids := []*ID{}
	users := []*UserDetail{}

	tx := db.Begin()
	if err := db.Raw("SELECT `uuid`,`id` FROM `users` GROUP BY `uuid`").Scan(&ids).Error; err != nil {

		return nil, err
	}

	for _, id := range ids {
		data := []*DataSet{}
		dataDetail := []*DataSetDetail{}
		if err := db.Raw("SELECT point,date FROM `users` WHERE uuid = ? GROUP BY `uuid`", id.UUID).Scan(&data).Error; err != nil {

			return nil, err
		}
		for _, d := range data {
			timeCards := []*TimeCards{}
			itemList := []*Item{}
			if err := db.Raw("SELECT * FROM `time_cards` WHERE id = ? AND date = ?", id.UUID, d.Date).Scan(&timeCards).Error; err != nil {

				return nil, err
			}
			for _, t := range timeCards {
				item := &Item{}
				if err := db.Raw("SELECT * FROM `items` WHERE uuid = ?", t.ItemID).Scan(&item).Error; err != nil {

					return nil, err
				}
				itemList = append(itemList, item)
			}
			dataDetail = append(dataDetail, &DataSetDetail{
				Score:    d.Point,
				Date:     d.Date,
				ItemList: itemList,
			})
		}
		users = append(users, &UserDetail{
			UUID:    id.UUID,
			ID:      id.ID,
			DataSet: dataDetail,
		})
	}

	return users, tx.Commit().Error
}

func GetUserSpecific(uid string) ([]*UserDetail, error) {
	ids := []*ID{}
	users := []*UserDetail{}

	tx := db.Begin()
	if err := db.Raw("SELECT `uuid`,`id` FROM `users` GROUP BY `uuid` HAVING id = ?", uid).Scan(&ids).Error; err != nil {

		return nil, err
	}

	for _, id := range ids {
		data := []*DataSet{}
		dataDetail := []*DataSetDetail{}
		if err := db.Raw("SELECT point,date FROM `users` WHERE uuid = ? GROUP BY `uuid`", id.UUID).Scan(&data).Error; err != nil {

			return nil, err
		}
		for _, d := range data {
			timeCards := []*TimeCards{}
			itemList := []*Item{}
			if err := db.Raw("SELECT * FROM `time_cards` WHERE id = ? AND date = ?", id.UUID, d.Date).Scan(&timeCards).Error; err != nil {

				return nil, err
			}
			for _, t := range timeCards {
				item := &Item{}
				if err := db.Raw("SELECT * FROM `items` WHERE uuid = ?", t.ItemID).Scan(&item).Error; err != nil {

					return nil, err
				}
				itemList = append(itemList, item)
			}
			dataDetail = append(dataDetail, &DataSetDetail{
				Score:    d.Point,
				Date:     d.Date,
				ItemList: itemList,
			})
		}
		users = append(users, &UserDetail{
			UUID:    id.UUID,
			ID:      id.ID,
			DataSet: dataDetail,
		})
	}

	return users, tx.Commit().Error
}
