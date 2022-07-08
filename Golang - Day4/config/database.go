package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func DBSetup() *DBConfig {
	dbConfig := DBConfig{
		Host:     "localhost",
		Port:     8080,
		User:     "root",
		Password: "password@123",
		DBName:   "ecommerce",
	}
	return &dbConfig
}

func DBUrl(dbConfig DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName)

}
