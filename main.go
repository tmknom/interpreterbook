package main

import (
	"fmt"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
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
	input := `
if (10 > 1) {
  if (10 > 1) {
    return 10;
  }

  return 1;
}
`

	parser.Debug = false
	p := parser.NewParser(lexer.NewLexer(input))
	program := p.ParseProgram()
	env := object.NewEnvironment()

	if len(p.Errors()) > 0 {
		fmt.Printf("\ninput = %s\n\n", input)
		fmt.Printf("%+v\n", p.Errors())
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		fmt.Printf("%s\n", evaluated.Inspect())
	}
}
