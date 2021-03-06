package evaluator

import (
	"fmt"
	"github.com/looplanguage/loop/models/object"
	"strings"
)

var BuiltinFunctions = map[string]*object.BuiltinFunction{
	"len": {
		Function: func(args []object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: fmt.Sprintf("not enough arguments to builtin function \"len\". expected=1. got=%d", len(args))}
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			}

			return &object.Error{Message: fmt.Sprintf("wrong type as argument for \"len\". got=%T", args[0])}
		},
	},
	"print": {
		Function: func(args []object.Object) object.Object {
			return Print(args, "")
		},
	},
}

func Print(args []object.Object, suffix string) object.Object {
	values := ""

	for _, arg := range args {
		if str, ok := arg.(*object.String); ok {
			value := strings.Replace(str.Value, `\n`, "\n", -1)

			values += value
			continue
		}

		if integer, ok := arg.(*object.Integer); ok {
			value := fmt.Sprintf("%d", integer.Value)

			values += value
		}
	}

	values += suffix
	fmt.Print(values)

	return &object.Null{}
}
