package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	. "github.com/VictorMilhomem/glox/glox/utils"
	"github.com/chzyer/readline"
)

func runFile(fpath string) {
	str := strings.NewReader(path.Base(fpath))
	bytes, err := io.ReadAll(str)
	Check(err)
	err = run(string(bytes[:]))
	Check(err)
}

func runPrompt() {
	rl, err := readline.New("> ")
	Check(err)
	defer rl.Close()

	for {
		line, err := rl.Readline()
		Check(err)
		err = run(line)
		Check(err)
	}
}

func run(source string) *LoxError {
	fmt.Println(source)

	return nil
}

func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println("Usage: glox [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}
}
