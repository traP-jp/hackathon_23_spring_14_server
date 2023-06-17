package model

type User struct {
	ID string `json:"id" gorm:"size:32;primary_key"`

	Password string `json:"-" gorm:"size:86"`

	Email         string `json:"email" gorm:"size:191;unique" validate:"omitempty,email"`
	StudentNumber string `json:"studentNumber" gorm:"size:8;unique" validate:"omitempty,studentNumber"`
	PhoneNumber   string `json:"phoneNumber" gorm:"type:tinytext" validate:"omitempty,phoneNumber"`
}
