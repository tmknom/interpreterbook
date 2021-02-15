package evaluator

import (
	"fmt"
	"monkey/object"
)

var builtins = map[string]*object.Builtin{
	"len": object.NewBuiltin(func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return object.NewError(fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args)))
		}

		switch arg := args[0].(type) {
		case *object.String:
			return object.NewInteger(int64(len(arg.Value)))
		default:
			return object.NewError(fmt.Sprintf("argument to `len` not supported, got %s", args[0].Type()))
		}
	}),
}
