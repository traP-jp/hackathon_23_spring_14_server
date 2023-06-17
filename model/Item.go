package model

type Items struct {
	ID          string `json:"id" gorm:"size:32;primary_key"`
	Description string `json:"description"`
	Point       int    `json:"point" validate:"omitempty"`
	Report      int    `json:"report" gorm:"type:tinytext" validate:"omitempty"`
}
