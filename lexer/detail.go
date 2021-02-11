package lexer

type DebugTracer struct {
	Line       string
	LineNumber int
}

func newDebugTracer() *DebugTracer {
	d := &DebugTracer{LineNumber: 1}
	d.resetLine()
	return d
}

func (d *DebugTracer) columnNumber() int {
	return len(d.Line)
}

func (d *DebugTracer) incrementLine() {
	d.LineNumber += 1
}

func (d *DebugTracer) appendChar(ch byte) {
	if ch != '\n' {
		d.Line += string(ch)
	}
}

func (d *DebugTracer) resetLine() {
	d.Line = ""
}
