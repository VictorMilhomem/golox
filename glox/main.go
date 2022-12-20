package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/VictorMilhomem/glox/ast"
	"github.com/VictorMilhomem/glox/glox/lexer"
	"github.com/VictorMilhomem/glox/glox/utils"
	"github.com/chzyer/readline"
)

func runFile(fpath string) {
	bytes, err := ioutil.ReadFile(path.Base(fpath))
	source := string(bytes)
	utils.Check(err)
	run(source)
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
	parser := ast.NewParser(tokens)
	statements := parser.Parse()

	interpreter := ast.NewInterpreter()

	interpreter.Interpret(statements)

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
