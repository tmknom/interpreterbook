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
		case *object.Array:
			return object.NewInteger(int64(len(arg.Elements)))
		case *object.String:
			return object.NewInteger(int64(len(arg.Value)))
		default:
			return object.NewError(fmt.Sprintf("argument to `len` not supported, got %s", args[0].Type()))
		}
	}),
	"first": object.NewBuiltin(func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return object.NewError(fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args)))
		}

		if args[0].Type() != object.ARRAY_OBJ {
			return object.NewError(fmt.Sprintf("argument to `first` must be ARRAY, got %s", args[0].Type()))
		}

		array := args[0].(*object.Array)
		if len(array.Elements) > 0 {
			return array.Elements[0]
		}
		return object.NULL
	}),
	"last": object.NewBuiltin(func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return object.NewError(fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args)))
		}

		if args[0].Type() != object.ARRAY_OBJ {
			return object.NewError(fmt.Sprintf("argument to `last` must be ARRAY, got %s", args[0].Type()))
		}

		array := args[0].(*object.Array)
		length := len(array.Elements)
		if length > 0 {
			return array.Elements[length-1]
		}
		return object.NULL
	}),
	"rest": object.NewBuiltin(func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return object.NewError(fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args)))
		}

		if args[0].Type() != object.ARRAY_OBJ {
			return object.NewError(fmt.Sprintf("argument to `last` must be ARRAY, got %s", args[0].Type()))
		}

		array := args[0].(*object.Array)
		length := len(array.Elements)
		if length > 0 {
			newElements := make([]object.Object, length-1, length-1)
			copy(newElements, array.Elements[1:length])
			return object.NewArray(newElements)
		}
		return object.NULL
	}),
	"push": object.NewBuiltin(func(args ...object.Object) object.Object {
		if len(args) != 2 {
			return object.NewError(fmt.Sprintf("wrong number of arguments. got=%d, want=2", len(args)))
		}

		if args[0].Type() != object.ARRAY_OBJ {
			return object.NewError(fmt.Sprintf("argument to `push` must be ARRAY, got %s", args[0].Type()))
		}

		array := args[0].(*object.Array)
		length := len(array.Elements)

		newElements := make([]object.Object, length+1, length+1)
		copy(newElements, array.Elements)
		newElements[length] = args[1]

		return object.NewArray(newElements)
	}),
}
