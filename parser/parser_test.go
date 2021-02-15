package parser_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"monkey/ast"
	"monkey/lexer"
	"monkey/parser"
	"monkey/token"
	"strconv"
	"testing"
)

func TestLetStatements(t *testing.T) {
	cases := []struct {
		input string
		want  *ast.LetStatement
	}{
		{
			input: "let x = 5;",
			want: &ast.LetStatement{
				Token: token.NewIdentifierToken("let"),
				Name:  ast.NewIdentifierByName("x"),
				Value: ast.NewIntegerLiteralByValue(5),
			},
		},
		{
			input: "let y = true;",
			want: &ast.LetStatement{
				Token: token.NewIdentifierToken("let"),
				Name:  ast.NewIdentifierByName("y"),
				Value: ast.NewBooleanByValue("true"),
			},
		},
		{
			input: "let foobar = y;",
			want: &ast.LetStatement{
				Token: token.NewIdentifierToken("let"),
				Name:  ast.NewIdentifierByName("foobar"),
				Value: ast.NewIdentifierByName("y"),
			},
		},
		{
			input: "let foobar = x + y;",
			want: &ast.LetStatement{
				Token: token.NewIdentifierToken("let"),
				Name:  ast.NewIdentifierByName("foobar"),
				Value: newInfixExpression(
					ast.NewIdentifierByName("x"),
					plusToken,
					ast.NewIdentifierByName("y"),
				),
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

		stmt, ok := program.Statements[0].(*ast.LetStatement)
		if !ok {
			t.Errorf("program.Statements[0] not *ast.LetStatement: %+v", stmt)
			continue
		}

		opt := cmpopts.IgnoreUnexported(*stmt.Token)
		if diff := cmp.Diff(stmt, tc.want, opt); diff != "" {
			t.Errorf("failed statement %q, diff (-got +want):\n%s", tc.input, diff)
		}
	}
}

func TestRetStatements(t *testing.T) {
	cases := []struct {
		input string
		want  *ast.ReturnStatement
	}{
		{
			input: "return 5;",
			want: &ast.ReturnStatement{
				Token:       token.NewIdentifierToken("return"),
				ReturnValue: ast.NewIntegerLiteralByValue(5),
			},
		},
		{
			input: "return true;",
			want: &ast.ReturnStatement{
				Token:       token.NewIdentifierToken("return"),
				ReturnValue: ast.NewBooleanByValue("true"),
			},
		},
		{
			input: "return foobar;",
			want: &ast.ReturnStatement{
				Token:       token.NewIdentifierToken("return"),
				ReturnValue: ast.NewIdentifierByName("foobar"),
			},
		},
		{
			input: "return x + y;",
			want: &ast.ReturnStatement{
				Token: token.NewIdentifierToken("return"),
				ReturnValue: newInfixExpression(
					ast.NewIdentifierByName("x"),
					plusToken,
					ast.NewIdentifierByName("y"),
				),
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

		stmt, ok := program.Statements[0].(*ast.ReturnStatement)
		if !ok {
			t.Errorf("program.Statements[0] not *ast.ReturnStatement: %+v", stmt)
			continue
		}

		opt := cmpopts.IgnoreUnexported(*stmt.Token)
		if diff := cmp.Diff(stmt, tc.want, opt); diff != "" {
			t.Errorf("failed statement %q, diff (-got +want):\n%s", tc.input, diff)
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
				newIntegerLiteralExpressionStatement(5),
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

func TestFunctionLiteral(t *testing.T) {
	cases := []struct {
		input string
		want  *ast.FunctionLiteral
	}{
		{
			input: "fn(x, y) { x + y; }",
			want: &ast.FunctionLiteral{
				Token: token.NewIdentifierToken("fn"),
				Parameters: []*ast.Identifier{
					ast.NewIdentifierByName("x"),
					ast.NewIdentifierByName("y"),
				},
				Body: &ast.BlockStatement{
					Token: token.NewToken(token.LBRACE, "{"),
					Statements: []ast.Statement{
						&ast.ExpressionStatement{
							Token: token.NewIdentifierToken("x"),
							Expression: newInfixExpression(
								ast.NewIdentifierByName("x"),
								plusToken,
								ast.NewIdentifierByName("y"),
							),
						},
					},
				},
			},
		},
		{
			input: "fn() {}",
			want: &ast.FunctionLiteral{
				Token:      token.NewIdentifierToken("fn"),
				Parameters: []*ast.Identifier{},
				Body: &ast.BlockStatement{
					Token:      token.NewToken(token.LBRACE, "{"),
					Statements: []ast.Statement{},
				},
			},
		},
		{
			input: "fn(x) {}",
			want: &ast.FunctionLiteral{
				Token: token.NewIdentifierToken("fn"),
				Parameters: []*ast.Identifier{
					ast.NewIdentifierByName("x"),
				},
				Body: &ast.BlockStatement{
					Token:      token.NewToken(token.LBRACE, "{"),
					Statements: []ast.Statement{},
				},
			},
		},
		{
			input: "fn(x, y) {}",
			want: &ast.FunctionLiteral{
				Token: token.NewIdentifierToken("fn"),
				Parameters: []*ast.Identifier{
					ast.NewIdentifierByName("x"),
					ast.NewIdentifierByName("y"),
				},
				Body: &ast.BlockStatement{
					Token:      token.NewToken(token.LBRACE, "{"),
					Statements: []ast.Statement{},
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

		exp, ok := stmt.Expression.(*ast.FunctionLiteral)
		if !ok {
			t.Errorf("stmt.Expression not *ast.FunctionLiteral: %+v", exp)
			continue
		}

		opt := cmpopts.IgnoreUnexported(*exp.Token)
		if diff := cmp.Diff(exp, tc.want, opt); diff != "" {
			t.Errorf("failed statement %q, diff (-got +want):\n%s", tc.input, diff)
		}
	}
}

func TestCallExpression(t *testing.T) {
	cases := []struct {
		input string
		want  *ast.CallExpression
	}{
		{
			input: "add(1, 2 * 3);",
			want: &ast.CallExpression{
				Token: token.NewToken(token.LPAREN, "("),
				Arguments: []ast.Expression{
					ast.NewIntegerLiteralByValue(1),
					newInfixExpression(
						ast.NewIntegerLiteralByValue(2),
						asteriskToken,
						ast.NewIntegerLiteralByValue(3),
					),
				},
				Function: ast.NewIdentifierByName("add"),
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

		exp, ok := stmt.Expression.(*ast.CallExpression)
		if !ok {
			t.Errorf("stmt.Expression not *ast.CallExpression: %+v", exp)
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
		{
			"a + add(b * c) + d",
			"((a + add((b * c))) + d)",
		},
		{
			"add(a, b, 1, 2 * 3, 4 + 5, add(6, 7 * 8))",
			"add(a, b, 1, (2 * 3), (4 + 5), add(6, (7 * 8)))",
		},
		{
			"add(a + b + c * d / f + g)",
			"add((((a + b) + ((c * d) / f)) + g))",
		},
		{
			"a * [1, 2, 3, 4][b * c] * d",
			"((a * ([1, 2, 3, 4][(b * c)])) * d)",
		},
		{
			"add(a * b[2], b[1], 2 * [1, 2][1])",
			"add((a * (b[2])), (b[1]), (2 * ([1, 2][1])))",
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

func TestStringLiteralExpression(t *testing.T) {
	input := `"hello world";`

	p := parser.NewParser(lexer.NewLexer(input))
	program := p.ParseProgram()
	checkParserError(t, p)

	stmt := program.Statements[0].(*ast.ExpressionStatement)
	literal, ok := stmt.Expression.(*ast.StringLiteral)
	if !ok {
		t.Fatalf("exp not *ast.StringLiteral. got=%T", stmt.Expression)
	}

	if literal.Value != "hello world" {
		t.Errorf("literal.Value not %q. got=%q", "hello world", literal.Value)
	}
}

func TestParsingEmptyArrayLiterals(t *testing.T) {
	input := "[]"

	p := parser.NewParser(lexer.NewLexer(input))
	program := p.ParseProgram()
	checkParserError(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	array, ok := stmt.Expression.(*ast.ArrayLiteral)
	if !ok {
		t.Fatalf("exp not ast.ArrayLiteral. got=%T", stmt.Expression)
	}

	if len(array.Elements) != 0 {
		t.Errorf("len(array.Elements) not 0. got=%d", len(array.Elements))
	}
}

func TestParsingArrayLiterals(t *testing.T) {
	cases := []struct {
		input string
		want  []ast.Expression
	}{
		{
			input: "[1, 2 * 2, 3 + 3]",
			want: []ast.Expression{
				ast.NewIntegerLiteralByValue(1),
				newInfixExpression(
					ast.NewIntegerLiteralByValue(2),
					asteriskToken,
					ast.NewIntegerLiteralByValue(2),
				),
				newInfixExpression(
					ast.NewIntegerLiteralByValue(3),
					plusToken,
					ast.NewIntegerLiteralByValue(3),
				),
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
		array, ok := stmt.Expression.(*ast.ArrayLiteral)
		if !ok {
			t.Fatalf("exp not ast.ArrayLiteral. got=%T", stmt.Expression)
		}

		if len(array.Elements) != 3 {
			t.Fatalf("len(array.Elements) not 3. got=%d", len(array.Elements))
		}

		for i, element := range array.Elements {
			var opt cmp.Option
			switch element := element.(type) {
			case *ast.IntegerLiteral:
				opt = cmpopts.IgnoreUnexported(*element.Token)
			case *ast.InfixExpression:
				opt = cmpopts.IgnoreUnexported(*element.Token)
			}
			if diff := cmp.Diff(element, tc.want[i], opt); diff != "" {
				t.Errorf("failed statement %q, diff (-got +want):\n%s", tc.input, diff)
			}
		}
	}
}

func TestParsingIndexExpressions(t *testing.T) {
	cases := []struct {
		input string
		want  ast.IndexExpression
	}{
		{
			input: "myArray[1 + 1]",
			want: ast.IndexExpression{
				Token: token.NewToken(token.LBRACKET, "["),
				Left:  ast.NewIdentifierByName("myArray"),
				Index: newInfixExpression(
					ast.NewIntegerLiteralByValue(1),
					plusToken,
					ast.NewIntegerLiteralByValue(1),
				),
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
		index, ok := stmt.Expression.(*ast.IndexExpression)
		if !ok {
			t.Fatalf("exp not ast.IndexExpression. got=%T", stmt.Expression)
		}

		opt := cmpopts.IgnoreUnexported(*index.Token)
		if diff := cmp.Diff(index.String(), tc.want.String(), opt); diff != "" {
			t.Errorf("failed statement %q, diff (-got +want):\n%s", tc.input, diff)
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

func newIntegerLiteralExpressionStatement(value int64) *ast.ExpressionStatement {
	return &ast.ExpressionStatement{
		Token:      token.NewIntegerToken(strconv.FormatInt(value, 10)),
		Expression: ast.NewIntegerLiteralByValue(value),
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
