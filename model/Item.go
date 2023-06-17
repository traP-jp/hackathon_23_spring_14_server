package model

type Items struct {
	ID          string `json:"id" gorm:"size:32;unique"`
	Description string `json:"description" gorm:"size:86;unique"`
	Point       int    `json:"point" validate:"omitempty"`
	Report      int    `json:"report" validate:"omitempty"`
}
