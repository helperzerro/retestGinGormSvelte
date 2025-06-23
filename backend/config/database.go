package database

import (
	"fmt"
	"log"
	"os"

	"github.com/helperzerro/ginGormSVELTE/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectDB mengembalikan koneksi *gorm.DB
func ConnectDB() *gorm.DB {
	// Muat variabel lingkungan dari file .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Ambil konfigurasi dari environment variables
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),  // Misalnya: root
		os.Getenv("DB_PASSWORD"), // Misalnya: kosong
		os.Getenv("DB_HOST"),  // Misalnya: localhost
		os.Getenv("DB_PORT"),  // Misalnya: 3306
		os.Getenv("DB_NAME"),  // Misalnya: crud-employee
	)

	// Membuka koneksi ke database
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal koneksi ke database: %v", err)
	}
	log.Println("✅ Berhasil terkoneksi ke database")

	if err := DB.AutoMigrate(&models.Employee{}, &models.User{}); err != nil {
	log.Fatal("Gagal melakukan AutoMigrate:", err)
}
log.Println("✅ AutoMigrate berhasil")


	return DB
}
