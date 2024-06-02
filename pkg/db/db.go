package db

import (
	"fmt"
	"log"

	"github.com/celalettin-ceylan/go-bookstore/pkg/utils"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Setup() *gorm.DB {

	utils.ReadEnvFile()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		viper.GetString("DB_HOST"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	log.Println("Database connected successfuly")
	return DB
}
