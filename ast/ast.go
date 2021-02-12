package ast

import (
	"fmt"
)

type Node interface {
	TokenLiteral() string
	fmt.Stringer
}
