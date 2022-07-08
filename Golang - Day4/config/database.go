package config

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
)

type Database struct {
	Orm *gorm.DB
	Mu  *sync.Mutex
}

var DB *Database

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
		Port:     3306,
		User:     "root",
		Password: "password@123",
		DBName:   "ecommerce",
	}
	return &dbConfig
}

func DBUrl(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName)

}
