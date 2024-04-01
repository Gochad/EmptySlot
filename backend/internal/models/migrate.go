package models

import (
	"fmt"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	models := []any{
		&Merchandise{},
		&Category{},
		&Reservation{},
		&User{},
		&History{},
	}

	if err := db.AutoMigrate(models...); err != nil {
		fmt.Println("Error auto-migrating model:", err)
	}
}
