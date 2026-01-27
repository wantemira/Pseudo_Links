// Package logger provides logging configuration
package logger

import (
	"log"
	"os"
	"time"
)

// NewLogger initializes application logger
func NewLogger() {
	file, err := os.OpenFile("/logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Failed to open log file: %v", err)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("error with file close %v", err)
			return
		}
	}()

	logger := log.New(file, "APP: ", log.LstdFlags)

	go func() {
		for {
			logger.Println("Hello from Dockerized Go app")
			logger.Printf("Current time: %s", time.Now().Format(time.RFC3339))
			time.Sleep(5 * time.Second)
		}
	}()
}
