package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"backend/internal"
)

func ConnectDb() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432",
		internal.EnvConfig.Session.DbUser,
		internal.EnvConfig.Session.DbPwd,
		internal.EnvConfig.Session.DbName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	db.Logger = logger.Default.LogMode(logger.Info)

	return db
}
