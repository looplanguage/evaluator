package evaluator

import (
	"github.com/looplanguage/loop/models/object"
)

func evalString(string string) object.Object {
	return &object.String{Value: string}
}
