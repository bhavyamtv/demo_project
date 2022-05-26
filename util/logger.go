package util

import (
	"log"
	"os"
	"runtime"
)

var Debug *log.Logger

func LogSet() {

	//Need to be Used for staging

	/*
		file, err := os.OpenFile("/var/log/automatorcustom.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				log.Println("Unable tp ctreate the file")
				return
			}
		Debug = log.New(file, "DEBUG  ", log.Ldate|log.Ltime|log.Lmicroseconds)
	*/

	// Used during Dev

	Debug = log.New(os.Stdout, "DEBUG  ", log.Ldate|log.Ltime|log.Lmicroseconds)

	return
}
func FuncName() {
	pc, _, _, _ := runtime.Caller(1)
	Debug.Println(runtime.FuncForPC(pc).Name())
}
