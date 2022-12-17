package utils

import (
	"log"
	"os"
	"strconv"

	"github.com/fatih/color"
)

type LoxError struct {
	Line      int
	Msg       string
	ErrorCode int
}

func (e *LoxError) Error() string {
	return "[line" + strconv.Itoa(e.Line) + "]" + "Error: " + e.Msg + "Exit Code: " + strconv.Itoa(e.ErrorCode)
}

func Check(e error) {
	if e != nil {
		color.Set(color.FgRed, color.Bold)
		log.Println(e)
		color.Unset() // Use it in your function
		os.Exit(-1)
	}
}
