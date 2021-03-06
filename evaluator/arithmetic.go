package evaluator

import (
	"fmt"
	"github.com/looplanguage/loop/models/ast"
	"github.com/looplanguage/loop/models/object"
)

func evalSuffixExpression(left object.Object, right object.Object, operator string) object.Object {
	if left, ok := left.(*object.Integer); ok {
		if right, ok := right.(*object.Integer); ok {
			switch operator {
			case "+":
				return &object.Integer{Value: left.Value + right.Value}
			case "*":
				return &object.Integer{Value: left.Value * right.Value}
			case "/":
				return &object.Integer{Value: left.Value / right.Value}
			case "-":
				return &object.Integer{Value: left.Value - right.Value}
			case ">":
				return &object.Boolean{Value: left.Value > right.Value}
			case "<":
				return &object.Boolean{Value: left.Value < right.Value}
			case ">=":
				return &object.Boolean{Value: left.Value >= right.Value}
			case "<=":
				return &object.Boolean{Value: left.Value <= right.Value}
			}
		}
	}

	if operator == "==" {
		if left.Type() == right.Type() && left.Inspect() == right.Inspect() {
			return &TRUE
		}

		return &FALSE
	}

	if operator == "+" {
		if left, ok := left.(*object.String); ok {
			if right, ok := right.(*object.String); ok {
				return &object.String{Value: fmt.Sprintf("%s%s", left.Value, right.Value)}
			}

			if right, ok := right.(*object.Integer); ok {
				return &object.String{Value: fmt.Sprintf("%s%d", left.Value, right.Value)}
			}
		}
	}

	return &object.Error{Message: fmt.Sprintf("invalid operator. got=%q", operator)}
}

func evalConditionalStatement(condition ast.Expression, ElseCondition *ast.BlockStatement, ElseStatement *ast.ConditionalStatement, Body *ast.BlockStatement, env *object.Environment) object.Object {
	run := Eval(condition, env)

	if run, ok := run.(*object.Boolean); ok {
		if run.Value {
			var lastEval object.Object

			for _, stmt := range Body.Statements {
				lastEval = Eval(stmt, env)
			}

			if lastEval == nil {
				return &object.Null{}
			}

			return lastEval
		} else {
			if ElseStatement != nil {
				return evalConditionalStatement(ElseStatement.Condition, ElseStatement.ElseCondition, ElseStatement.ElseStatement, ElseStatement.Body, env)
			} else if ElseCondition != nil && len(ElseCondition.Statements) > 0 {
				var lastEval object.Object

				for _, stmt := range ElseCondition.Statements {
					lastEval = Eval(stmt, env)
				}

				return lastEval
			}
		}
	} else {
		return &object.Error{Message: fmt.Sprintf("condition is of invalid type. expected=%q. got=%q", "BOOLEAN", run.Type())}
	}

	return &object.Null{}
}
