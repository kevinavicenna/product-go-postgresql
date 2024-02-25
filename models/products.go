package models

type Product struct {
	ID          int8    `gorm:"primary key;autoIncrement" json:"id"`
	Name        *string `json:"name"`
	Description *string `json:description`
	Category    *string `json:category`
}

func MigrateProduct(db *gorm.DB) error {
	err := db.AutoMigrate{&Product{}}
	return err
}
