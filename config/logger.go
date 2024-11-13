package config

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
)

var (
	once      = sync.Once{}
	singleton *Logger
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	error   *log.Logger
	fatal   *log.Logger
	writer  io.Writer
}

func NewLogger() *Logger {
	writer := io.Writer(os.Stdout)
	return &Logger{
		debug:   log.New(writer, fmt.Sprintf("%sDEBUG: %s", colorBlue, colorReset), log.LstdFlags),
		info:    log.New(writer, fmt.Sprintf("%sINFO: %s", colorGreen, colorReset), log.LstdFlags),
		warning: log.New(writer, fmt.Sprintf("%sWARNING: %s", colorYellow, colorReset), log.LstdFlags),
		error:   log.New(writer, fmt.Sprintf("%sERROR: %s", colorRed, colorReset), log.LstdFlags),
		fatal:   log.New(writer, fmt.Sprintf("%sFATAL ERROR: %s", colorRed, colorReset), log.LstdFlags),
		writer:  writer,
	}
}
func NewSingletonLogger() *Logger {
	once.Do(func() {
		singleton = NewLogger()
	})

	return singleton
}
func NewFileLogger(path string) (*Logger, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	return &Logger{
		debug:   log.New(file, "DEBUG: ", log.LstdFlags),
		info:    log.New(file, "INFO: ", log.LstdFlags),
		warning: log.New(file, "WARNING: ", log.LstdFlags),
		error:   log.New(file, "ERROR: ", log.LstdFlags),
		fatal:   log.New(file, "FATAL ERROR: ", log.LstdFlags),
		writer:  file,
	}, err
}

func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Warning(v ...interface{}) {
	l.warning.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.error.Println(v...)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.fatal.Fatal(v...)
}
