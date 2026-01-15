package log

import (
	"fmt"
	"log"
)

type Logger struct{}

func NewLogger() Logger {
	return Logger{}
}

func (Logger) Println(in ...any) {
	log.Println(in...)
}

func (Logger) Printf(format string, v ...any) {
	log.Printf(format, v...)
}

func (Logger) Panic(in ...any) {
	log.Panic(in...)
}

func (Logger) Blue(in ...any) {
	log.Println(blue + fmt.Sprint(in...) + reset)
}
func (Logger) Yellow(in ...any) {
	log.Println(yellow + fmt.Sprint(in...) + reset)
}
func (Logger) Green(in ...any) {
	log.Println(green + fmt.Sprint(in...) + reset)
}
func (Logger) Red(in ...any) {
	log.Println(red + fmt.Sprint(in...) + reset)
}

// Info level log (blue)
func (Logger) Info(op string, msg string) {
	log.Println(blue + fmt.Sprintf("%s:%s", op, msg) + reset)
}

// Success level log (green)
func (Logger) Success(op string, msg string) {
	log.Println(green + fmt.Sprintf("%s:%s", op, msg) + reset)
}

// Warn level log (yellow)
func (Logger) Warn(op string, msg string, err error) {
	if err == nil {
		log.Println(yellow + fmt.Sprintf("%s:%s", op, msg) + reset)
	} else {
		log.Println(yellow + fmt.Sprintf("%s:%s:%s", op, msg, err.Error()) + reset)
	}
}

// Error level log (red)
func (Logger) Error(op string, msg string, err error) {
	if err == nil {
		log.Println(red + fmt.Sprintf("%s:%s", op, msg) + reset)
	} else {
		log.Println(red + fmt.Sprintf("%s:%s:%s", op, msg, err.Error()) + reset)
	}
}
