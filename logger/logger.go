package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
)


type Logger interface {
	Warn(message string)
	Info(message string)
	Error(message string)
}

type logger struct {
	warnLog *log.Logger
	infoLog *log.Logger
	errorLog *log.Logger
}

func (lg *logger) Warn(message string) {
	lg.warnLog.Println(message)
}
func (lg *logger) Info(message string) {
	lg.infoLog.Println(message)
}
func (lg *logger) Error(message string) {
	lg.infoLog.Println(message)
}

func init() {
	fmt.Println("logger package init called")
}

var (
	logInstance Logger
	// once sync.Once
	mutex sync.Mutex
)

func customSigleton( fun func() ) {
	mutex.Lock()
	defer mutex.Unlock()

	if logInstance != nil {
		return
	}
	fun()
}

func GetLogger() Logger {
	// once.Do(func() {
	// 	logFile, err := os.OpenFile("logFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	logInstance = &logger{
	// 		warnLog: log.New(logFile, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile),
	// 		infoLog: log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
	// 		errorLog: log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	// 	}
	// })

	customSigleton(func() {
		logFile, err := os.OpenFile("logFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		logInstance = &logger{
			warnLog: log.New(logFile, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile),
			infoLog: log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
			errorLog: log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		}
	})

	return logInstance
}