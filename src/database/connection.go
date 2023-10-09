package database

import (
	"github.com/geraldojrcg/go-sample-api/src/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDatabaseConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(".db/database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&model.User{}, &model.Todo{})
	if err != nil {
		panic("failed to create migrations")
	}

	return db
}
