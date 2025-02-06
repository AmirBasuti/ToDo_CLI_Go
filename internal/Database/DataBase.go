package Database

import (
	"ToDo/internal/Model"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Connection() error {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("Error loading .env file")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("connection to the data base faild \n %s\n\n", err)
	}
	return nil
}

func Migration() error {
	if err := DB.AutoMigrate(&Model.Todo{}); err != nil {
		return fmt.Errorf("Migration faild \n %s \n\n", err)
	}
	return nil
}
