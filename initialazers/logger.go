package initialazers

import (
	"log"
	"os"
)

var (
	Logger *log.Logger
)

func InitLogger() {
	var logPath = "../log/log.log"
	var logFile, err = os.OpenFile(logPath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println("logger file error", err)
	}
	Logger = log.New(logFile, "ERROR :", log.Ldate|log.Ltime|log.Lshortfile)
}
