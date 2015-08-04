package logger

import (
	"fmt"
	"io"
	"time"
)

type LogLevel string

const (
	LOG_DEBUG     LogLevel = "debug"
	LOG_INFO      LogLevel = "info"
	LOG_NOTICE    LogLevel = "notice"
	LOG_WARNING   LogLevel = "warning"
	LOG_ERROR     LogLevel = "error"
	LOG_CRITIAL   LogLevel = "critial"
	LOG_ALERT     LogLevel = "alert"
	LOG_EMERGENCY LogLevel = "emergency"
)

type Logger struct {
	out  io.Writer
	name string
}

func (self *Logger) Output(level LogLevel, msg string) {
	now := time.Now()
	dateTime := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d.%03d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond()/1e6)
	self.out.Write([]byte(self.name + " -- " + string(level) + " -- [" + dateTime + "] -- " + msg))
}

func (self *Logger) Debug(msg string) {
	self.Output(LOG_DEBUG, msg)
}

func (self *Logger) Info(msg string) {
	self.Output(LOG_INFO, msg)
}

func (self *Logger) Notice(msg string) {
	self.Output(LOG_NOTICE, msg)
}

func (self *Logger) Warning(msg string) {
	self.Output(LOG_WARNING, msg)
}

func (self *Logger) Error(msg string) {
	self.Output(LOG_ERROR, msg)
}

var logger *Logger = &Logger{}

func InitGlobal(name string, out io.Writer) {
	logger.name = name
	logger.out = out
}

func New(name string, out io.Writer) (logger *Logger) {
	logger = new(Logger)
	logger.name = name
	logger.out = out

	return
}

func Debug(msg string) {
	logger.Debug(msg)
}

func Info(msg string) {
	logger.Info(msg)
}

func Notice(msg string) {
	logger.Notice(msg)
}

func Warning(msg string) {
	logger.Warning(msg)
}

func Error(msg string) {
	logger.Error(msg)
}
