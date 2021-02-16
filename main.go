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
let map = fn(arr, f) {
  let iter = fn(arr, accumulated) {
    if (len(arr) == 0){
      accumulated
    } else {
      iter(rest(arr), push(accumulated, f(first(arr))));
    }
  }
  iter(arr, []);
}

let reduce = fn(arr, initial, f) {
  let iter = fn(arr, result) {
    if (len(arr) == 0){
      result
    } else {
      iter(rest(arr), f(result, first(arr)));
    }
  }
  iter(arr, initial);
}

let a = [1,2,3,4];
let double = fn(x) { x * 2 };
map(a, double);

let sum = fn(arr) {
  reduce(arr, 0, fn(initial, el){initial + el});
}

sum([1, 2, 3, 4, 5]);

let people = [{"name":"Alice", "age": 24}, {"name":"Anna", "age": 28}];
people[0]["name"];
people[1]["age"];
people[0]["age"] + people[1]["age"];

let getName = fn(person) { person["name"]; }
getName(people[0])
getName(people[1])
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
