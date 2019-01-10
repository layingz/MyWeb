package models

import (
	"os"
	"log"
)

type LogOp interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
	Close()
}

type Logger struct {
	logPtr *log.Logger
	filePtr *os.File
}

func newLogFile(path string) *os.File{
	f, _ := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	return f
}

func(l *Logger)Println(v ...interface{}){
	l.logPtr.Println(v ...)
}

func(l *Logger)Printf(format string, v ...interface{}){
	l.logPtr.Printf(format, v ...)
}

func (l *Logger)Close(){
	l.filePtr.Close()
}

func NewLog(logname string) LogOp{
	path := "/var/log/" + logname + ".log"
	f := newLogFile(path)
	return &Logger{logPtr: log.New(f, "", log.Ldate|log.Ltime), filePtr: f}
}

var WebLog LogOp

func WebLogInit(){
	WebLog = NewLog("myweb")
}