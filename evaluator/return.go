package evaluator

import (
	"github.com/looplanguage/loop/models/ast"
	"github.com/looplanguage/loop/models/object"
)

func evalReturn(value ast.Expression) object.Object {
	r := &object.Return{Value: value}

	return r
}
