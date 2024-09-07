package logger

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", log.LstdFlags)

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

func Fatal(args ...any) {
	logger.Fatal(args)
}

func Fatalf(format string, args ...any) {
	logger.Fatalf(format, args)
}

func Fatalln(args ...any) {
	logger.Fatalln(args)
}
