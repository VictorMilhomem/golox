package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/VictorMilhomem/glox/glox/lexer"
	"github.com/VictorMilhomem/glox/glox/utils"
	"github.com/chzyer/readline"
)

func runFile(fpath string) {
	str := strings.NewReader(path.Base(fpath))
	bytes, err := io.ReadAll(str)
	utils.Check(err)
	run(string(bytes[:]))
}

func runPrompt() {
	rl, err := readline.New("glox > ")
	utils.Check(err)
	defer rl.Close()

	for {
		line, err := rl.Readline()
		utils.Check(err)
		run(line)
	}
}

func run(source string) *utils.LoxError {
	scanner := lexer.NewScanner(source)
	tokens := scanner.ScanTokens()

	for _, token := range tokens {
		fmt.Println(token.ToString())
	}

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
