package models

import "gorm.io/gorm"

type Products struct {
	ID          uint    `gorm:"primary key;autoIncrement" json:"id"`
	Name        *string `json:"name"`
	Description *string `json:description`
	Category    *string `json:category`
}

func MigrateProduct(db *gorm.DB) error {
	err := db.AutoMigrate(&Products{})
	return err
}
