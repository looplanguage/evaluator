package evaluator

import (
	"fmt"
	"github.com/looplanguage/loop/models/ast"
	"github.com/looplanguage/loop/models/object"
)

func evalWhile(condition ast.Expression, block *ast.BlockStatement, env *object.Environment) object.Object {
	exp := Eval(condition, env)

	bool, ok := exp.(*object.Boolean)

	if !ok {
		return &object.Error{Message: fmt.Sprintf("wrong expression value. expected=\"BOOLEAN\". got=%q", exp.Type())}
	}

	var rValue object.Object

	for bool.Value {
		bool = Eval(condition, env).(*object.Boolean)

		rValue = Eval(block, env)

		if _, ok := rValue.(*object.Return); ok {
			break
		}
	}

	if rValue == nil {
		rValue = &object.Null{}
	}

	return rValue
}
