package config

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Database() *gorm.DB {
	createDirStorageLogsDatabase()
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

    // Set up GORM logger
    newLogger := logger.New(
        loggerNew, // Use logrus as the GORM logger output
        logger.Config{
            SlowThreshold: time.Second,   // Slow SQL query threshold
            LogLevel:      logger.Info,   // Log level (Info, Warn, Error)
            IgnoreRecordNotFoundError: true, // Ignore ErrRecordNotFound error for logger
            Colorful:      false,         // Disable color output (logrus handles formatting)
        },
    )

    // Connect to the database using GORM
    // migrate -path db/migrations -database "mysql://arifin:Arifin123\!@tcp(10.217.18.4:3306)/lennadb" down
    // migrate -path db/migrations -database "mysql://arifin:Arifin123\!@tcp(10.217.18.4:3306)/lennadb" down
    dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"),os.Getenv("DB_PASSWORD"),os.Getenv("DB_HOST"),os.Getenv("DB_PORT"),os.Getenv("DB_NAME"))
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger: newLogger, // Set the custom GORM logger
    })

    if err != nil {
        loggerNew.Error("Failed to connect to the database:", err)
        return db
    }

    loggerNew.Info("Connected to the database successfully")

	return db
}

func createDirStorageLogsDatabase() {
	dir := "./storage/logs/database"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0744)
		if err != nil {
			fmt.Println(dir, "can't created directory")
		}
		fmt.Println("success created directory", dir)
	} else {
		fmt.Println("The provided directory named", dir, "exists")
	}
}