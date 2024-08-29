package main

import (
	"fmt"
	"log"
	"os"
	"time"

	appconfig "github.com/lenna-ai/azureOneSmile.git/config/appConfig"
	seeds "github.com/lenna-ai/azureOneSmile.git/db/seeds/userSeeds"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// Ganti dengan DSN MySQL Anda

	date := time.Now().Format("01-02-2006")
    logLumberJack := &lumberjack.Logger{
        Filename:   fmt.Sprintf("storage/logs/database/%v.log",date),
        MaxSize:    10,  // Maximum size in megabytes before log is rotated
        MaxBackups: 7,   // Maximum number of old log files to keep
        MaxAge:     1,   // Maximum number of days to retain old log files
        Compress:   true, // Compress old log files
        LocalTime:  true,                        // Use local time for log rotation
    }
	loggerNew := logrus.New()
    loggerNew.SetOutput(logLumberJack)
    loggerNew.SetFormatter(&logrus.JSONFormatter{})
	newLogger := logger.New(
        loggerNew, // Use logrus as the GORM logger output
        logger.Config{
            SlowThreshold: time.Second,   // Slow SQL query threshold
            LogLevel:      logger.Info,   // Log level (Info, Warn, Error)
            IgnoreRecordNotFoundError: true, // Ignore ErrRecordNotFound error for logger
            Colorful:      false,         // Disable color output (logrus handles formatting)
        },
    )
	appconfig.InitApplication()
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"),os.Getenv("DB_PASSWORD"),os.Getenv("DB_HOST"),os.Getenv("DB_PORT"),os.Getenv("DB_NAME"))
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger: newLogger, // Set the custom GORM logger
    })
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}


	// Jalankan seeder untuk menambahkan data palsu
	if err := seeds.SeedUsers(db, 10); err != nil {
		log.Fatalf("failed to seed users: %v", err)
	}

	log.Println("Seeding completed successfully")
}
