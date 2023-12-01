package database

import (
	"fmt"

	"github.com/smugg99/coding-night-2023/config"
	"github.com/smugg99/coding-night-2023/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Setup() (*gorm.DB, error) {
    dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Warsaw",
		config.Conf.DbConfig.Host,
		config.Conf.DbConfig.User,
		config.Conf.DbConfig.Password,
		"haseus",
		"5432",
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
        &models.User{},
    )

	return db, nil
}
