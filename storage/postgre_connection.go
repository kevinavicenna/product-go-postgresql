package storage

import "fmt"

//import (
//	"fmt"
//	"gorm.io/driver/postgres"
//	"gorm.io/gorm"
//)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	db       string
	SSLMode  string
}

func NewConnection(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s password=%s user=%s db=%s sslmode=%s",
		config.Host,config.Port,config.Password,config.User,config.db,config.SSLMode
	)
	db,err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		return db,err
	}
	return db,nil
}
