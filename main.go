package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID string `json:"id" gorm:"size:32;primary_key"`

	Password string `json:"-" gorm:"size:86"`

	Email         string `json:"email" gorm:"size:191;unique" validate:"omitempty,email"`
	StudentNumber string `json:"studentNumber" gorm:"size:8;unique" validate:"omitempty,studentNumber"`
	PhoneNumber   string `json:"phoneNumber" gorm:"type:tinytext" validate:"omitempty,phoneNumber"`
}

// var (
// 	db *gorm.DB

// 	schemas = []interface{}{
// 		&User{},
// 	}
// )

func Setup() error {
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "root"
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "password"
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		port = 3306
	}

	dbname := os.Getenv("DB_DATABASE")
	if dbname == "" {
		dbname = "portal"
	}

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %v", err)
	}
	var users []User
	db.Find(&users)
	fmt.Println(users)
	return nil
}
func main() {
	fmt.Println("Hello World")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	Setup()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	e.Logger.Fatal(e.Start(":" + port))

}
