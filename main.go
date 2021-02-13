package main

import (
	"fmt"
	"monkey/lexer"
	"monkey/parser"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	runRepl()
}

func runRepl() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

func runDebugger() {
	parser.Debug = true
	input := "add(1, 2 * 3, 4 + 5);"
	p := parser.NewParser(lexer.NewLexer(input))
	p.ParseProgram()

	if len(p.Errors()) > 0 {
		fmt.Printf("\ninput = %s\n\n", input)
		fmt.Printf("%+v\n", p.Errors())
	}
}
