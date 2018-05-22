package logger

import (
	"runtime"
	"log"
	"fmt"
)

func init(){
	log.SetFlags(log.Ldate | log.Ltime )
}

func Debug(msg... interface{}) {
	_, filename, line, _ := runtime.Caller(1) 
	log.Println(fmt.Sprintf("[%s] %s:%d ", "DEBG", filename, line), msg)
}

func Info(msg... interface{}) {
	_, filename, line, _ := runtime.Caller(1) 
	log.Println(fmt.Sprintf("[%s] %s:%d ", "INFO", filename, line), msg)
}

func Warn(msg... interface{}) {
	_, filename, line, _ := runtime.Caller(1) 
	log.Println(fmt.Sprintf("[%s] %s:%d ", "WARN", filename, line), msg)
}

func Error(msg... interface{}) {
	_, filename, line, _ := runtime.Caller(1) 
	log.Println(fmt.Sprintf("[%s] %s:%d ", "ERRO", filename, line), msg)
}

