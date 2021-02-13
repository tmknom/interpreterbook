package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
)

type Object interface {
	Type() ObjectType
	Inspect() string
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
