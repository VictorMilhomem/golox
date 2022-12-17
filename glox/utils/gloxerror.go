package utils

import (
	"log"
	"strconv"
)

type LoxError struct {
	Line      int64
	msg       string
	errorCode int64
}

func (e *LoxError) Error() string {
	return "[line" + strconv.FormatInt(e.Line, 10) + "]" + "Error: " + e.msg + "Exit Code: " + strconv.FormatInt(e.errorCode, 10)
}

func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
