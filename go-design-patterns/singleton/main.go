package main

import (
	"fmt"
	"sync"
)

func main() {

	for i := 1; i < 10; i++ {
		go getLoggerInstance()
	}
}

type MyLogger struct {
	loglevel int
}

func (l *MyLogger) Log(s string) {
	fmt.Println(l.loglevel, ":", s)
}

func (l *MyLogger) SetLogLevel(level int) {
	l.loglevel = level
}

var logger *MyLogger
var once sync.Once

func getLoggerInstance() *MyLogger {

	once.Do(func() {
		fmt.Println("Creating Logger Instance")
		logger = &MyLogger{}
	})
	fmt.Println("Return Logger instance")
	return logger
}
