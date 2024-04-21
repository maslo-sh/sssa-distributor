package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"privileges-management/model"
)

func ConnectToDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitializeSchema() {
	db, err := ConnectToDatabase()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.ApprovingPermission{}, &model.AccessRequest{}, &model.Resource{})

	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDb.SetMaxIdleConns(20)
	sqlDb.SetMaxOpenConns(5)
}
