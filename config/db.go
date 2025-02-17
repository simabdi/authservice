package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
)

var DB *gorm.DB

func ConnectDatabase() {
	rootEnvPath := filepath.Join("..", ".env") // Adjust if needed
	localEnvPath := ".env"                     // Default location

	if err := godotenv.Load(localEnvPath); err != nil {
		log.Println("⚠️ Could not load local .env, trying parent directory...")
		if err := godotenv.Load(rootEnvPath); err != nil {
			log.Fatal("❌ Failed to load .env file from any location:", err)
		}
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Database connection failed:", err)
	}

	DB = db
	fmt.Println("✅ Database connected successfully")
}
