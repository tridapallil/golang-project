package utils

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func ConnectDB() *gorm.DB {
	
	err := godotenv.Load()

	err != nil{
		log.Fatalf( format: "Error loading .env file" )
	}

	dsn := os.Getenv( key: "dsn" )
	
	db, err := gorm.Open( dialect: os.Getenv( key: "dbtype" ), dsn )

	if (err != nil) {
		log.Fatalf(format: "Error connecting to database: %% v", err)
		panic(err)
	}

	defer db.Close()

	db.AutoMigrate(&domain.User{})

	return db
}

