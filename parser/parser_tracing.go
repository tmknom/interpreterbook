package parser

import (
	"fmt"
	"strings"
)

var Debug = false
var traceLevel int = 0

const traceIdentPlaceholder string = "    "

func identLevel() string {
	return strings.Repeat(traceIdentPlaceholder, traceLevel-1)
}

func tracePrint(fs string) {
	if Debug {
		fmt.Printf("%s%s\n", identLevel(), fs)
	}
}

func incIdent() { traceLevel = traceLevel + 1 }
func decIdent() { traceLevel = traceLevel - 1 }

func trace(msg string) string {
	incIdent()
	tracePrint("BEGIN " + msg)
	return msg
}

func untrace(msg string) {
	tracePrint("END " + msg)
	decIdent()
}

func traceDetail(msg string) {
	tracePrint("")
	tracePrint(traceIdentPlaceholder + "DETAIL " + msg)
	tracePrint("")
}
