package initializers

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"sync"
	"time"
)

type DatabaseConnection struct {
}

var (
	Db      *gorm.DB
	OauthDb *gorm.DB
	err     error
	once    sync.Once
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func DbConnection() error {
	//once.Do(func() {
	loadEnv()
	// Read the environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:                              true,
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	sqlDB, err := Db.DB()
	if err != nil {
		log.Fatalf("failed to access the underlying database: %v", err)
	}

	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Minute * 10)
	sqlDB.SetConnMaxIdleTime(0)
	//})

	return nil
}

func OauthDatabaseConnection() error {
	//once.Do(func() {
	dbHost := viper.GetString("oauth.dbHost")
	dbPort := viper.GetString("oauth.dbPort")
	dbUser := viper.GetString("oauth.dbUsername")
	dbPassword := viper.GetString("oauth.dbPassword")
	dbName := viper.Get("oauth.dbName")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	OauthDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt:                              true,
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	sqlDB, err := OauthDb.DB()
	if err != nil {
		log.Fatalf("failed to access the underlying database: %v", err)
	}

	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Minute * 10)
	sqlDB.SetConnMaxIdleTime(0)
	//})
	return nil
}
