package object

import (
	"bytes"
	"fmt"
	"monkey/ast"
	"strings"
)

type ObjectType string
type BuiltinFunction func(args ...Object) Object

const (
	STRING_OBJ       = "STRING"
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	FUNCTION_OBJ     = "FUNCTION"
	BUILTIN_OBJ      = "BUILTIN"
	ERROR_OBJ        = "ERROR"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type String struct {
	Value string
}

func NewString(value string) *String {
	return &String{Value: value}
}

var _ Object = (*String)(nil)

func (s String) Type() ObjectType {
	return STRING_OBJ
}

func (s String) Inspect() string {
	return s.Value
}

type Integer struct {
	Value int64
}

func NewInteger(value int64) *Integer {
	return &Integer{Value: value}
}

var _ Object = (*Integer)(nil)

func (i Integer) Type() ObjectType {
	return INTEGER_OBJ
}

func (i Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

type Boolean struct {
	Value bool
}

var (
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
)

func NewBoolean(value bool) *Boolean {
	if value {
		return TRUE
	}
	return FALSE
}

var _ Object = (*Boolean)(nil)

func (b Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

func (b Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

type Null struct{}

var NULL = &Null{}

var _ Object = (*Null)(nil)

func (n Null) Type() ObjectType {
	return NULL_OBJ
}

func (n Null) Inspect() string {
	return "null"
}

type ReturnValue struct {
	Value Object
}

func NewReturnValue(value Object) *ReturnValue {
	return &ReturnValue{Value: value}
}

var _ Object = (*ReturnValue)(nil)

func (r ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}

func (r ReturnValue) Inspect() string {
	return r.Value.Inspect()
}

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func NewFunction(parameters []*ast.Identifier, body *ast.BlockStatement, env *Environment) *Function {
	return &Function{
		Parameters: parameters,
		Body:       body,
		Env:        env,
	}
}

var _ Object = (*Function)(nil)

func (f Function) Type() ObjectType {
	return FUNCTION_OBJ
}

func (f Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, parameter := range f.Parameters {
		params = append(params, parameter.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

type Builtin struct {
	Fn BuiltinFunction
}

func NewBuiltin(fn BuiltinFunction) *Builtin {
	return &Builtin{Fn: fn}
}

var _ Object = (*Builtin)(nil)

func (b Builtin) Type() ObjectType {
	return BUILTIN_OBJ
}

func (b Builtin) Inspect() string {
	return "builtin function"
}

type Error struct {
	Message string
}

var _ Object = (*Error)(nil)

func NewError(message string) *Error {
	return &Error{Message: message}
}

func (e Error) Type() ObjectType {
	return ERROR_OBJ
}

func (e Error) Inspect() string {
	return "ERROR: " + e.Message
}
