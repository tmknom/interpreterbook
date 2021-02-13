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
				newIdentifierExpressionStatement("foobar"),
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

func TestParsingPrefixExpressions(t *testing.T) {
	cases := []struct {
		input string
		want  *ast.PrefixExpression
	}{
		{
			input: "!5;",
			want:  newPrefixExpression(bangToken, ast.NewIntegerLiteralByValue(5)),
		},
		{
			input: "-15;",
			want:  newPrefixExpression(minusToken, ast.NewIntegerLiteralByValue(15)),
		},
		{
			input: "!true;",
			want:  newPrefixExpression(bangToken, ast.NewBooleanByValue("true")),
		},
		{
			input: "!false;",
			want:  newPrefixExpression(bangToken, ast.NewBooleanByValue("false")),
		},
	}

	for _, tc := range cases {
		p := parser.NewParser(lexer.NewLexer(tc.input))
		program := p.ParseProgram()
		checkParserError(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d",
				len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Errorf("program.Statements[0] not *ast.ExpressionStatement: %+v", stmt)
			continue
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Errorf("stmt.Expression not *ast.PrefixExpression: %+v", exp)
			continue
		}

		opt := cmpopts.IgnoreUnexported(*exp.Token)
		if diff := cmp.Diff(exp, tc.want, opt); diff != "" {
			t.Errorf("failed statement %q, diff (-got +want):\n%s", tc.input, diff)
		}
	}
}

func TestParsingInfixExpressions(t *testing.T) {
	cases := []struct {
		input string
		want  *ast.InfixExpression
	}{
		{
			input: "5 + 5;",
			want: newInfixExpression(
				ast.NewIntegerLiteralByValue(5),
				plusToken,
				ast.NewIntegerLiteralByValue(5),
			),
		},
		{
			input: "5 - 5;",
			want: newInfixExpression(
				ast.NewIntegerLiteralByValue(5),
				minusToken,
				ast.NewIntegerLiteralByValue(5),
			),
		},
		{
			input: "5 * 5;",
			want: newInfixExpression(
				ast.NewIntegerLiteralByValue(5),
				asteriskToken,
				ast.NewIntegerLiteralByValue(5),
			),
		},
		{
			input: "5 / 5;",
			want: newInfixExpression(
				ast.NewIntegerLiteralByValue(5),
				slashToken,
				ast.NewIntegerLiteralByValue(5),
			),
		},
		{
			input: "5 > 5;",
			want: newInfixExpression(
				ast.NewIntegerLiteralByValue(5),
				gtToken,
				ast.NewIntegerLiteralByValue(5),
			),
		},
		{
			input: "5 < 5;",
			want: newInfixExpression(
				ast.NewIntegerLiteralByValue(5),
				ltToken,
				ast.NewIntegerLiteralByValue(5),
			),
		},
		{
			input: "5 == 5;",
			want: newInfixExpression(
				ast.NewIntegerLiteralByValue(5),
				eqToken,
				ast.NewIntegerLiteralByValue(5),
			),
		},
		{
			input: "5 != 5;",
			want: newInfixExpression(
				ast.NewIntegerLiteralByValue(5),
				notEqToken,
				ast.NewIntegerLiteralByValue(5),
			),
		},
		{
			input: "true == true;",
			want: newInfixExpression(
				ast.NewBooleanByValue("true"),
				eqToken,
				ast.NewBooleanByValue("true"),
			),
		},
		{
			input: "true != false;",
			want: newInfixExpression(
				ast.NewBooleanByValue("true"),
				notEqToken,
				ast.NewBooleanByValue("false"),
			),
		},
		{
			input: "false == false;",
			want: newInfixExpression(
				ast.NewBooleanByValue("false"),
				eqToken,
				ast.NewBooleanByValue("false"),
			),
		},
	}

	for _, tc := range cases {
		p := parser.NewParser(lexer.NewLexer(tc.input))
		program := p.ParseProgram()
		checkParserError(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d",
				len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Errorf("program.Statements[0] not *ast.ExpressionStatement: %+v", stmt)
			continue
		}

		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Errorf("stmt.Expression not *ast.InfixExpression: %+v", exp)
			continue
		}

		opt := cmpopts.IgnoreUnexported(*exp.Token)
		if diff := cmp.Diff(exp, tc.want, opt); diff != "" {
			t.Errorf("failed statement %q, diff (-got +want):\n%s", tc.input, diff)
		}
	}
}

func TestIfExpression(t *testing.T) {
	cases := []struct {
		input string
		want  *ast.IfExpression
	}{
		{
			input: "if (x < y) { x }",
			want: &ast.IfExpression{
				Token: token.NewIdentifierToken("if"),
				Condition: newInfixExpression(
					ast.NewIdentifierByName("x"),
					ltToken,
					ast.NewIdentifierByName("y"),
				),
				Consequence: &ast.BlockStatement{
					Token: token.NewToken(token.LBRACE, "{"),
					Statements: []ast.Statement{
						newIdentifierExpressionStatement("x"),
					},
				},
			},
		},
	}

	for _, tc := range cases {
		p := parser.NewParser(lexer.NewLexer(tc.input))
		program := p.ParseProgram()
		checkParserError(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d",
				len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Errorf("program.Statements[0] not *ast.ExpressionStatement: %+v", stmt)
			continue
		}

		exp, ok := stmt.Expression.(*ast.IfExpression)
		if !ok {
			t.Errorf("stmt.Expression not *ast.IfExpression: %+v", exp)
			continue
		}

		opt := cmpopts.IgnoreUnexported(*exp.Token)
		if diff := cmp.Diff(exp, tc.want, opt); diff != "" {
			t.Errorf("failed statement %q, diff (-got +want):\n%s", tc.input, diff)
		}
	}
}

func TestIfElseExpression(t *testing.T) {
	cases := []struct {
		input string
		want  *ast.IfExpression
	}{
		{
			input: "if (x < y) { x } else { y }",
			want: &ast.IfExpression{
				Token: token.NewIdentifierToken("if"),
				Condition: newInfixExpression(
					ast.NewIdentifierByName("x"),
					ltToken,
					ast.NewIdentifierByName("y"),
				),
				Consequence: &ast.BlockStatement{
					Token: token.NewToken(token.LBRACE, "{"),
					Statements: []ast.Statement{
						newIdentifierExpressionStatement("x"),
					},
				},
				Alternative: &ast.BlockStatement{
					Token: token.NewToken(token.LBRACE, "{"),
					Statements: []ast.Statement{
						newIdentifierExpressionStatement("y"),
					},
				},
			},
		},
	}

	for _, tc := range cases {
		p := parser.NewParser(lexer.NewLexer(tc.input))
		program := p.ParseProgram()
		checkParserError(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d",
				len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Errorf("program.Statements[0] not *ast.ExpressionStatement: %+v", stmt)
			continue
		}

		exp, ok := stmt.Expression.(*ast.IfExpression)
		if !ok {
			t.Errorf("stmt.Expression not *ast.IfExpression: %+v", exp)
			continue
		}

		opt := cmpopts.IgnoreUnexported(*exp.Token)
		if diff := cmp.Diff(exp, tc.want, opt); diff != "" {
			t.Errorf("failed statement %q, diff (-got +want):\n%s", tc.input, diff)
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"true",
			"true",
		},
		{
			"false",
			"false",
		},
		{
			"3 > 5 == false",
			"((3 > 5) == false)",
		},
		{
			"3 < 5 == true",
			"((3 < 5) == true)",
		},
		{
			"1 + (2 + 3) + 4",
			"((1 + (2 + 3)) + 4)",
		},
		{
			"(5 + 5) * 2",
			"((5 + 5) * 2)",
		},
		{
			"2 / (5 + 5)",
			"(2 / (5 + 5))",
		},
		{
			"(5 + 5) * 2 * (5 + 5)",
			"(((5 + 5) * 2) * (5 + 5))",
		},
		{
			"-(5 + 5)",
			"(-(5 + 5))",
		},
		{
			"!(true == true)",
			"(!(true == true))",
		},
	}

	for _, tc := range cases {
		p := parser.NewParser(lexer.NewLexer(tc.input))
		program := p.ParseProgram()
		checkParserError(t, p)

		actual := program.String()
		if actual != tc.want {
			t.Errorf("want=%q, got=%q", tc.want, actual)
		}
	}
}

func newInfixExpression(left ast.Expression, t *token.Token, right ast.Expression) *ast.InfixExpression {
	return &ast.InfixExpression{
		Token:    t,
		Left:     left,
		Operator: t.Literal,
		Right:    right,
	}
}

func newPrefixExpression(t *token.Token, exp ast.Expression) *ast.PrefixExpression {
	return &ast.PrefixExpression{
		Token:    t,
		Operator: t.Literal,
		Right:    exp,
	}
}

func newIdentifierExpressionStatement(identifier string) *ast.ExpressionStatement {
	return &ast.ExpressionStatement{
		Token:      token.NewIdentifierToken(identifier),
		Expression: ast.NewIdentifierByName(identifier),
	}
}

var (
	bangToken     = token.NewToken(token.BANG, "!")
	minusToken    = token.NewToken(token.MINUS, "-")
	plusToken     = token.NewToken(token.PLUS, "+")
	asteriskToken = token.NewToken(token.ASTERISK, "*")
	slashToken    = token.NewToken(token.SLASH, "/")
	gtToken       = token.NewToken(token.GT, ">")
	ltToken       = token.NewToken(token.LT, "<")
	eqToken       = token.NewToken(token.EQ, "==")
	notEqToken    = token.NewToken(token.NOT_EQ, "!=")
)

func checkParserError(t *testing.T, p *parser.Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors, input = %q", len(errors), p.Input())
	for _, err := range errors {
		t.Errorf("parser error: %s", err)
	}
	t.FailNow()
}
