package utils

import (
	"log"
	"strconv"
)

type LoxError struct {
	Line int
	Msg  string
}

func (e *LoxError) Error() string {
	return "[line" + strconv.Itoa(e.Line) + "]" + "Error: " + e.Msg
}

func Check(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}
