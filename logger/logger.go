package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
)

type Level int

const (
	ErrorLevel Level = 0
	WarnLevel  Level = 1
	InfoLevel  Level = 2
	DebugLevel Level = 3
)

type Logger struct {
	out       io.WriteCloser
	Level     Level
	Logger    *log.Logger
	requestID string
}

func levelToString(level Level) string {
	switch level {
	case InfoLevel:
		return "info"
	case ErrorLevel:
		return "error"
	case WarnLevel:
		return "warn"
	case DebugLevel:
		return "debug"
	}
	return ""
}

func levelToFormatString(level Level) string {
	switch level {
	case InfoLevel:
		return "[" + levelToString(level) + "]"
	case WarnLevel:
		return "[" + levelToString(level) + "]"
	case ErrorLevel:
		return "[" + levelToString(level) + "]"
	case DebugLevel:
		return "[" + levelToString(level) + "]"
	}
	return ""
}

var logFlags = log.Ldate | log.Ltime | log.Lmicroseconds

func NewLogger(path string, logLevel Level) Logger {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("日志文件无法打开，请检查， err:" + err.Error())
	}
	logger := Logger{
		out:    file,
		Level:  logLevel,
		Logger: log.New(file, "", logFlags),
	}
	return logger
}

func (l *Logger) SetRequestID(requestID string) {
	l.requestID = requestID
}

func (l *Logger) Close() error {
	return l.out.Close()
}

func (l *Logger) Debug(args ...interface{}) {
	if l.Level < DebugLevel {
		return
	}

	preFixArray := make([]interface{}, 0, 10)
	preFixArray = append(preFixArray, getCaller())
	preFixArray = append(preFixArray, l.requestID)
	preFixArray = append(preFixArray, levelToFormatString(DebugLevel))

	args = append(preFixArray, args...)
	l.Logger.Print(args...)
}

func (l Logger) INFO(args ...interface{}) {
	if l.Level < InfoLevel {
		return
	}

	preFixArray := make([]interface{}, 0, 10)
	preFixArray = append(preFixArray, getCaller())
	preFixArray = append(preFixArray, l.requestID)
	preFixArray = append(preFixArray, levelToFormatString(InfoLevel))

	args = append(preFixArray, args...)
	l.Logger.Print(args...)
}

func (l Logger) Warn(args ...interface{}) {
	if l.Level < WarnLevel {
		return
	}

	preFixArray := make([]interface{}, 0, 10)
	preFixArray = append(preFixArray, getCaller())
	preFixArray = append(preFixArray, l.requestID)
	preFixArray = append(preFixArray, levelToFormatString(WarnLevel))

	args = append(preFixArray, args...)
	l.Logger.Println(args...)
}

func (l Logger) Error(args ...interface{}) {
	if l.Level < ErrorLevel {
		return
	}

	preFixArray := make([]interface{}, 0, 10)
	preFixArray = append(preFixArray, getCaller())
	preFixArray = append(preFixArray, l.requestID)
	preFixArray = append(preFixArray, levelToFormatString(ErrorLevel))

	args = append(preFixArray, args...)
	l.Logger.Println(args...)
}

func getCaller() string {
	pc, fullPath, line, ok := runtime.Caller(2)
	if !ok {
		return ""
	}

	pcName := runtime.FuncForPC(pc).Name()
	file := fullPath

	pcNameParts := strings.Split(pcName, "/")
	pcName = pcNameParts[len(pcNameParts)-1]
	pcNameParts = strings.Split(pcName, ".")
	pcName = pcNameParts[len(pcNameParts)-1]

	return fmt.Sprintf("%s:%s:%d", file, pcName, line)
}

var GLog Logger
