package evaluator

import (
	"fmt"
	"github.com/looplanguage/loop/models/ast"
	"github.com/looplanguage/loop/models/object"
)

func evalIdentifier(identifier *ast.Identifier, env *object.Environment) object.Object {
	value := env.Get(identifier.Value)

	if value != nil {
		return value
	}

	if builtin, ok := BuiltinFunctions[identifier.Value]; ok {
		return builtin
	}

	return &object.Error{Message: fmt.Sprintf("identifier %q not defined", identifier.Value), Line: identifier.Token.Line, Column: identifier.Token.Column}
}
