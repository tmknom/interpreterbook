package parser

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`
	cases := []struct {
		want []*ast.LetStatement
	}{
		{
			want: []*ast.LetStatement{
				ast.NewLetStatementByName("x"),
				ast.NewLetStatementByName("y"),
				ast.NewLetStatementByName("foobar"),
			},
		},
	}

	p := NewParser(lexer.NewLexer(input))
	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	for i, tc := range cases {
		got := program.Statements[i].(*ast.LetStatement)
		opt := cmpopts.IgnoreUnexported(*got.Token)
		if diff := cmp.Diff(got, tc.want[i], opt); diff != "" {
			t.Errorf("failed statement: diff (-got +want):\n%s", diff)
		}
	}
}
