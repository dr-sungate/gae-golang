package logger

import (
	"runtime"
	"log"
	"fmt"
)
func init(){
	log.SetFlags(log.Ldate | log.Ltime )
}

func Println (msg... interface{}) {
	_, filename, line, _ := runtime.Caller(1) 
	log.Println(fmt.Sprintf("%s:%d", filename, line), msg)
}