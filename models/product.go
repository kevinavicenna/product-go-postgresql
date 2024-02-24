package models

type Products struct {
	ID          int8    `gorm:"primary key;autoIncrement" json:"id"`
	Name        *string `json:"name"`
	Description *string `json:description`
	Category    *string `json:category`
}

func MigrateProduct(db *gorm.DB) error {
	err := db.AutoMigrate{&Products{}}
	return err
}
