package evaluator

import (
	"github.com/looplanguage/loop/models/ast"
	"github.com/looplanguage/loop/models/object"
)

func evalVariableDeclaration(key *ast.Identifier, value ast.Expression, env *object.Environment) object.Object {
	v := Eval(value, env)

	set := env.Set(key.Value, v)

	if _, ok := v.(*object.Function); ok {
		return &object.Null{}
	}

	return set
}
