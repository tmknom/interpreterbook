package parser_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"monkey/ast"
	"monkey/lexer"
	"monkey/parser"
	"monkey/token"
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

	p := parser.NewParser(lexer.NewLexer(input))
	program := p.ParseProgram()
	checkParserError(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	for i, tc := range cases {
		stmt := program.Statements[i]
		got, ok := stmt.(*ast.LetStatement)
		if !ok {
			t.Errorf("stmt not *ast.LetStatement: %+v", stmt)
			continue
		}

		opt := cmpopts.IgnoreUnexported(*got.Token)
		if diff := cmp.Diff(got, tc.want[i], opt); diff != "" {
			t.Errorf("failed statement: diff (-got +want):\n%s", diff)
		}
	}
}

func TestRetStatements(t *testing.T) {
	input := `
return 5;
return 10;
return 993322;
`
	cases := []struct {
		want []*ast.ReturnStatement
	}{
		{
			want: []*ast.ReturnStatement{
				ast.NewReturnStatement(),
				ast.NewReturnStatement(),
				ast.NewReturnStatement(),
			},
		},
	}

	p := parser.NewParser(lexer.NewLexer(input))
	program := p.ParseProgram()
	checkParserError(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	for i, tc := range cases {
		stmt := program.Statements[i]
		got, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement: %+v", stmt)
			continue
		}

		opt := cmpopts.IgnoreUnexported(*got.Token)
		if diff := cmp.Diff(got, tc.want[i], opt); diff != "" {
			t.Errorf("failed statement: diff (-got +want):\n%s", diff)
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := `
foobar;
`
	cases := []struct {
		want []*ast.ExpressionStatement
	}{
		{
			want: []*ast.ExpressionStatement{
				&ast.ExpressionStatement{
					Token:      token.NewIdentifierToken("foobar"),
					Expression: ast.NewIdentifierByName("foobar"),
				},
			},
		},
	}

	p := parser.NewParser(lexer.NewLexer(input))
	program := p.ParseProgram()
	checkParserError(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d",
			len(program.Statements))
	}

	for i, tc := range cases {
		stmt := program.Statements[i]
		got, ok := stmt.(*ast.ExpressionStatement)
		if !ok {
			t.Errorf("stmt not *ast.ExpressionStatement: %+v", stmt)
			continue
		}

		opt := cmpopts.IgnoreUnexported(*got.Token)
		if diff := cmp.Diff(got, tc.want[i], opt); diff != "" {
			t.Errorf("failed statement: diff (-got +want):\n%s", diff)
		}
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := `
5;
`
	cases := []struct {
		want []*ast.ExpressionStatement
	}{
		{
			want: []*ast.ExpressionStatement{
				&ast.ExpressionStatement{
					Token:      token.NewIntegerToken("5"),
					Expression: ast.NewIntegerLiteralByValue(5),
				},
			},
		},
	}

	p := parser.NewParser(lexer.NewLexer(input))
	program := p.ParseProgram()
	checkParserError(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d",
			len(program.Statements))
	}

	for i, tc := range cases {
		stmt := program.Statements[i]
		got, ok := stmt.(*ast.ExpressionStatement)
		if !ok {
			t.Errorf("stmt not *ast.ExpressionStatement: %+v", stmt)
			continue
		}

		opt := cmpopts.IgnoreUnexported(*got.Token)
		if diff := cmp.Diff(got, tc.want[i], opt); diff != "" {
			t.Errorf("failed statement: diff (-got +want):\n%s", diff)
		}
	}
}

func checkParserError(t *testing.T, p *parser.Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, err := range errors {
		t.Errorf("parser error: %s", err)
	}
	t.FailNow()
}
