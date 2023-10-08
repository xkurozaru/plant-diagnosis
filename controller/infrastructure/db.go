package infrastructure

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

var (
	db  *gorm.DB
	err error
)

var dbConfig = DBConfig{
	Host:     "db",
	User:     os.Getenv("POSTGRES_USER"),
	Password: os.Getenv("POSTGRES_PASSWORD"),
	DBName:   os.Getenv("POSTGRES_DB"),
	Port:     os.Getenv("POSTGRES_PORT"),
}

func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.Port,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	dbInstance, err := db.DB()
	if err != nil {
		panic(err)
	}

	err = dbInstance.Close()
	if err != nil {
		panic(err)
	}
}
