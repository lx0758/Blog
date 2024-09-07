package logger

import (
	"io"
	"log"
	"os"
)

var logger = log.New(Writer(), "", log.LstdFlags)

type loggerWriter struct{}

func (_ loggerWriter) Write(bytes []byte) (n int, err error) {
	logger.Print(bytes)
	return 0, err
}

func Writer() io.Writer {
	return os.Stdout
}

func Logger() *log.Logger {
	return logger
}

func Print(args ...any) {
	logger.Print(args)
}

func Printf(format string, args ...any) {
	logger.Printf(format, args)
}

func Println(args ...any) {
	logger.Println(args)
}

func Panic(args ...any) {
	logger.Panic(args)
}

func Panicf(format string, args ...any) {
	logger.Panicf(format, args)
}

func Panicln(args ...any) {
	logger.Panicln(args)
}
