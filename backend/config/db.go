package config

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase เชื่อมต่อ SQLite database
func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("Manly_Massage.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	DB = database
	fmt.Println("✅ SQLite database connected successfully")
}
