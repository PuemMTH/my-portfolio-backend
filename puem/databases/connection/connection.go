package connection

import (
	"fmt"
	"log"
	"os"
	"puem/puem/types"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initializeDatabase() (*gorm.DB, error) {
	// get env from Environtment
	err := godotenv.Load()
	if err != nil {
		fmt.Println("DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME are required")
		log.Fatal("Error loading .env file")
	}

	Database := types.DB{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Database.User,
		Database.Password,
		Database.Host,
		Database.Port,
		Database.Name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db, nil
}

func GetDatabase() (*gorm.DB, error) { // GetDatabase is a function to get database connection
	return initializeDatabase()
}
