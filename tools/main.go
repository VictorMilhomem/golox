package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
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
	types := []string{
		"Binary   : Expr left, Token operator, Expr right",
		"Grouping : Expr expression",
		"Literal  : Object value",
		"Unary    : Token operator, Expr right",
	}
	defineAst(dir, "Expr", types)
}

func defineAst(dir string, basename string, types []string) {
	path := "./" + dir + "/" + basename + ".go"

	file, err := os.Create(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	boilerHeaders(file, dir, basename)
	var structNames []string
	for _, t := range types {
		structName := strings.Trim(strings.Split(t, ":")[0], " ")
		fields := strings.Trim(strings.Split(t, ":")[1], " ")
		defineType(file, basename, structName, fields)
		structNames = append(structNames, structName)
	}
	emitLine(file, "type "+basename+"Visitor[T Types] interface{")
	for _, s := range structNames {
		defineVisitor(file, basename, s)
	}
	emitLine(file, "}")
}

func boilerHeaders(file *os.File, dir string, basename string) {
	//type ExprS[T Types] struct {
	//	expr *Expr[T]
	//	Types interface{}
	//}

	emitLine(file, "package "+dir)
	emitLine(file, "import(")
	emitLine(file, ". \"github.com/VictorMilhomem/glox/glox/lexer\"")
	emitLine(file, "\"golang.org/x/exp/constraints\"")
	emitLine(file, ")")
	emitLine(file, "type Types interface{")
	emitLine(file, "    constraints.Ordered")
	emitLine(file, "}")

	emitLine(file, "type I"+basename+"[T Types] interface {")
	emitLine(file, "    Accept(visitor "+basename+"Visitor[T]) "+" interface{}")
	emitLine(file, "}")
}

func fieldsSplit(file *os.File, basename string, str string) {
	fields := strings.Split(str, ", ")

	for _, field := range fields {
		t := strings.Split(field, " ")[0]
		value := strings.Split(field, " ")[1]

		if t == basename {
			t = "I" + basename
			t += "[T]"
		}

		emitLine(file, "    "+UpperCaseFirstChar(value)+" "+t)
	}
}

func defineVisitor(file *os.File, basename, structName string) {
	var genVis string = "Visit" + structName + "(" + strings.ToLower(basename) + " " + structName + "[T]) interface{}"
	emitLine(file, "    "+genVis)
}

func defineType(file *os.File, basename, structName, fields string) {
	var vis string = "Visit" + structName + "(*v)"

	emitLine(file, "type "+structName+"[T Types] struct {")
	fieldsSplit(file, basename, fields)
	emitLine(file, "}")

	emitLine(file, "func (v *"+structName+"[T])"+"Accept(visitor "+basename+"Visitor[T]) interface{} {")
	emitLine(file, "    return visitor."+vis)
	emitLine(file, "}")
}

func emit(file *os.File, code string) {
	_, err := file.WriteString(code)
	if err != nil {
		log.Fatalln(err)
	}
}

func emitLine(file *os.File, code string) {
	_, err := file.WriteString(code + "\n")
	if err != nil {
		log.Fatalln(err)
	}
}

func UpperCaseFirstChar(str string) string {
	str = strings.ToLower(str)
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}

	return ""
}
