package model

import (
	"time"

	"github.com/gofrs/uuid"
)

type TimeCards struct {
	ID     uuid.UUID `json:"uuid" gorm:"primary_key;type:char(36)"`
	Date   time.Time `json:"date" gorm:"primary_key"`
	ItemID uuid.UUID `json:"uuid" gorm:"primary_key;type:char(36)"`
}

func AddTimeCards(rawID string, rawtuid string, userid string) (*TimeCards, error) {
	if err := db.Exec("INSERT INTO `time_cards` (`id`,`date`,`item_id`) VALUES (?,?,?)", rawID, time.Now().Format(time.DateOnly), rawtuid).Error; err != nil {
		return nil, err
	}
	item := Item{}
	if err := db.Raw("SELECT * FROM `items` WHERE uuid = ?", rawtuid).Scan(&item).Error; err != nil {
		return nil, err
	}
	user := User{}
	if err := db.Raw("SELECT * FROM `users` WHERE uuid = ? AND date = ?", rawID, time.Now().Format(time.DateOnly)).Scan(&user).Error; err != nil {
		return nil, err
	}
	if user.UUID.IsNil() {
		if err := db.Exec("INSERT INTO `users` (`uuid`,`id`,`point`,`date`) VALUES (?,?,?,?)", rawID, userid, item.Point, time.Now()).Error; err != nil {
			return nil, err
		}

	} else {

		user.Point += item.Point
		if err := db.Exec("UPDATE `users` SET `point` = ? WHERE `uuid` = ?", user.Point, user.UUID).Error; err != nil {
			return nil, err
		}
	}
	card := &TimeCards{}

	return card, nil
}
