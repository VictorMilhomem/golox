package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		log.Fatal("Usage: gen <outputdir>")
	}
	dir := args[0]
	if err := os.Mkdir(dir, os.ModePerm); err != nil {
		fmt.Println("")
	}
	var types []string
	defineAst(dir, "Expr", types)
}

func defineAst(dir string, basename string, types []string) {
	path := dir + "/" + basename + ".go"

	file, err := os.Create(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	emitLine(file, "package "+dir)
	emitLine(file, "type "+basename+" interface {")
	emitLine(file, "}")
}

func emit(file *os.File, code string) {
	_, err := file.WriteString(code)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
}

func emitLine(file *os.File, code string) {
	_, err := file.WriteString(code + "\n")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
}
