package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	AppConfig *Config
)

// Config menyimpan konfigurasi aplikasi
type Config struct {
	AppPort          string
	DBHost           string
	DBPort           string
	DBUser           string
	DBPassword       string
	DBName           string
	JWTSecret        string
	JWTRefreshToken  string
	JWTExpire        string
}

// LoadEnv membaca file .env dan mengisi nilai ke AppConfig
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}

	AppConfig = &Config{
		AppPort:          getEnv("PORT", "3030"),
		DBHost:           getEnv("DB_HOST", "localhost"),
		DBPort:           getEnv("DB_PORT", "5432"),
		DBUser:           getEnv("DB_USER", "postgres"),
		DBPassword:       getEnv("DB_PASSWORD", "password"),
		DBName:           getEnv("DB_NAME", "project_management"),
		JWTSecret:        getEnv("JWT_SECRET", "rahasia"),
		JWTRefreshToken:  getEnv("REFRESH_TOKEN_EXPIRED", "24h"),
		JWTExpire:        getEnv("JWT_EXPIRED", "2h"),
	}
}

// getEnv mengambil nilai variabel environment dengan fallback default
func getEnv(key string, fallback string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return fallback
}

// ConnectDB membuat koneksi ke PostgreSQL menggunakan GORM
func ConnectDB() {
	cfg := AppConfig

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance: ", err)
	}

	// Konfigurasi connection pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db
}
