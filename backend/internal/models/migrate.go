package models

import (
	"fmt"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	mods := []interface{}{
		&Merchandise{},
		&Category{},
		&Reservation{},
		&User{},
		&Customer{},
	}
	for _, model := range mods {
		if err := db.AutoMigrate(model); err != nil {
			fmt.Println("Error auto-migrating model:", err)
		}
	}
}
