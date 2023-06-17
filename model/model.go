package model

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db *gorm.DB

	schemas = []interface{}{
		&User{},
		&Items{},
		&TimeCards{},
	}
)

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

	db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %v", err)
	}

	db = db.
		Set("gorm:save_associations", false).
		Set("gorm:association_save_reference", false).
		Set("gorm:table_options", "CHARSET=utf8mb4").
		Set("gorm:table_options", "ENGINE=InnoDB")
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(2)
	sqlDB.SetMaxOpenConns(16)

	if os.Getenv("GORM_DEBUG") != "" {
		db.Logger = db.Logger.LogMode(logger.Info)
	}

	tx := db.Begin()

	// if err := tx.AutoMigrate(schemas...); err != nil {
	// 	tx.Rollback()
	// 	return fmt.Errorf("failed to connect database: %v", err)
	// }

	// var items []TimeCards
	// db.First(&items)

	// fmt.Println(items)
	return tx.Commit().Error
}
