package config

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Logger(app *fiber.App) {
	createDirStorageLogs()
	log := &lumberjack.Logger{
        Filename:   fmt.Sprintf("storage/logs/%v.log",time.Now().Format("2017-09-07")),
        MaxSize:    1000000,  // Maximum size in megabytes before log is rotated
        MaxBackups: 7,   // Maximum number of old log files to keep
        MaxAge:     7,   // Maximum number of days to retain old log files
        Compress:   true, // Compress old log files
    }

    loggerNew := logrus.New()
    loggerNew.SetOutput(log)
    loggerNew.SetFormatter(&logrus.JSONFormatter{})

    app.Use(logger.New(logger.Config{
        Output: loggerNew.Writer(), // Direct Fiber logs to the custom logger
        Format: "${time} ${status} - ${method} ${path} ${body}\n", // Customize log format (optional)
    }))
}
func createDirStorageLogs() {
	dirs := []string{
		"./storage/logs",
	}
	for _, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err := os.MkdirAll(dir, 0777)
			if err != nil {
				fmt.Println(dir, "can't create directory")
			}
			fmt.Println("success created directory", dir)
		} else {
			fmt.Println("The provided directory named", dir, "exists")
		}
	}
}