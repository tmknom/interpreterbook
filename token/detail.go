package token

import "fmt"

type DetailToken struct {
	Line         string
	LineNumber   int
	ColumnNumber int
}

func NewDetailToken(line string, lineNumber int, columnNumber int) *DetailToken {
	return &DetailToken{
		Line:         line,
		LineNumber:   lineNumber,
		ColumnNumber: columnNumber,
	}
}

func (d *DetailToken) String() string {
	return fmt.Sprintf("&DetailToken{line: %d, column: %d, line: '%s'}", d.LineNumber, d.ColumnNumber, d.Line)
}
