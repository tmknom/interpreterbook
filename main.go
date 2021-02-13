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
	runDebugger()
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
	input := "1 * (2 + 3);"
	p := parser.NewParser(lexer.NewLexer(input))
	p.ParseProgram()
}
