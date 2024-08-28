package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func setupLogger(filename string) *logrus.Logger {
	// numberInt,_ := strconv.Atoi(os.Getenv("TIME_STORAGE_DAY"))

	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Failed to open log file: %v", err)
    }
    defer logFile.Close()

	// Set file permissions to 777
    err = os.Chmod(filename, 0777)
    if err != nil {
        log.Fatalf("Failed to change file permissions: %v", err)
    }

	lumberjackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10000,
		MaxBackups: 7,
		MaxAge:    1,
		Compress:   true,
		LocalTime:  true,                        // Use local time for log rotation
	}

	log := logrus.New()
	log.SetOutput(lumberjackLogger)
	log.SetFormatter(&logrus.JSONFormatter{})
	return log
}

var GeneralLogger *logrus.Logger


func Logger(app *fiber.App) {
	createDirStorageLogs()

	app.Use(func(c *fiber.Ctx) error {

		date := time.Now().Format("01-02-2006")
		generalLogFile := fmt.Sprintf("./storage/logs/general_log/%s.log", date)
		dailyLogFile := fmt.Sprintf("./storage/logs/%s.log", date)

		GeneralLogger = setupLogger(generalLogFile)
		dailyLogger := setupLogger(dailyLogFile)


		start := time.Now()
		err := c.Next()
		latency := time.Since(start)
		latencyStr := fmt.Sprintf("%dms", latency.Milliseconds())

		body := c.Body()
		var bodyJSON interface{}
		if err := json.Unmarshal(body, &bodyJSON); err == nil {
			compactBody, _ := json.Marshal(bodyJSON)
			c.Locals("body", string(compactBody))
		} else {
			c.Locals("body", string(body))
		}

		GeneralLogger.WithFields(logrus.Fields{
			"body":          c.Locals("body"),
			"queryParams":   c.OriginalURL(),
			"reqHeaders":    c.GetReqHeaders(),
			"time":          time.Now().Format("15:04:05"),
			"date":          date,
			"status":        c.Response().StatusCode(),
			"ip":            c.IP(),
			"method":        c.Method(),
			"url":           c.OriginalURL(),
			"path":          c.Path(),
			"route":         c.Route().Path,
			"error":         err,
			"resBody":       string(c.Response().Body()),
			"responseTime":  latencyStr,
		}).Info("Request logged")

		customLogEntry := fmt.Sprintf(
			"body : %s | queryParams : %s | reqHeaders : %v | time : %s | date : %s | status : %d | ip : %s | method : %s | url : %s | path : %s | route : %s | error : %v | resBody : %s | responseTime : %s",
			c.Locals("body"),
			c.OriginalURL(),
			c.GetReqHeaders(),
			time.Now().Format("15:04:05"),
			date,
			c.Response().StatusCode(),
			c.IP(),
			c.Method(),
			c.OriginalURL(),
			c.Path(),
			c.Route().Path,
			err,
			string(c.Response().Body()),
			latencyStr,
		)
	
		dailyLogger.Out.Write([]byte(customLogEntry + "\n"))

		return nil
	})
}

func createDirStorageLogs() {
	dirs := []string{
		"./storage/logs/general_log",
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
