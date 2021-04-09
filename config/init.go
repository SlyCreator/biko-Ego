package config

import (
	"fmt"
	"github.com/SlyCreator/biko-Ego/entity"
	"github.com/joho/godotenv"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//OpenDatabaseConnection is connecting to database
func OpenDatabaseConnection() *gorm.DB  {
		errEnv := godotenv.Load(".env")
		if errEnv != nil {
			panic("Failed to load env file" )
		}

		dbUser := os.Getenv("DB_USER")
		dbPass := os.Getenv("DB_PASSWORD")
		dbHost := os.Getenv("DB_HOST")
		dbName := os.Getenv("DB_NAME")


		dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",dbUser,dbPass,dbHost,dbName)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("Failed to create a database connection")
		}
		db.AutoMigrate(&entity.User{})
		return db
}

func CloseDatabaseConnection(db *gorm.DB)   {
	dbSQL, err := db.DB()
	if err == nil {
		panic("Failed to close connection")
	}
	dbSQL.Close()
}