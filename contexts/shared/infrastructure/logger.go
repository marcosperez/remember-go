package infrastructure

import (
	"fmt"
	"time"
)

type ILogger interface {
	log(s string) bool
	error(s string) bool
}

type ConsoleLogger struct {
	ILogger
}

func (l *ConsoleLogger) Log(msg string) bool {
	fmt.Println(msg)

	return true
}

func (l *ConsoleLogger) Error(msg error) bool {
	currentTime := time.Now()

	fmt.Printf("[Error][%s] %s", currentTime.Format("20060102150405"), msg)

	return true
}

var Logger = ConsoleLogger{}
