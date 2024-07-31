package main

import (
	"errors"
	"fmt"
	"runtime"
	"time"
)

type Logger struct {

}

func (logger *Logger) log(handler func(any),message any) {
	pc,file,line,_ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc)
	fmt.Println(file,line,"in function",funcName.Name())
	handler(message)
	fmt.Println(time.Now())
}

func errLog(err any) {
	_,ok := err.(error)
	if ok {
		fmt.Printf("error occured: ")
		fmt.Println(err)
	}
} 

const (
	textBlack = iota + 30
	textRed
	textGreen
	textYellow
	textBlue
	textPurple
	textCyan
	textWhite
)

type Config struct {
	color int
	message string
}

func colorfulLog(config any) {
	_config,ok := config.(Config)
	if ok {
		fmt.Printf("\x1b[0;%dm%s\x1b[0m\n", _config.color, _config.message)
	}
}

func main() {
	
	logger := Logger{}
	logger.log(errLog,errors.New("test error"))
	logger.log(colorfulLog,Config{color: textCyan,message: "test colorful message"})
}