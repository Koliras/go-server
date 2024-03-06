package config

import (
	"fmt"
	"os"

	"github.com/Koliras/go-server/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
    if err := godotenv.Load(); err != nil {
        panic(err)
    }
    sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("HOST"),
        os.Getenv("PORT"),
        os.Getenv("USER"),
        os.Getenv("PASSWORD"),
        os.Getenv("DBNAME"),
    )

    db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
    db.AutoMigrate(&model.User{})
    if err != nil {
        panic(err)
    }

    return db
}
