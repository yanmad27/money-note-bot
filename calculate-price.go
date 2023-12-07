package main

import "github.com/Pramod-Devireddy/go-exprtk"

func calculatePrice(price string) int {
	exprtkObj := exprtk.NewExprtk()

	exprtkObj.SetExpression(price)

	err := exprtkObj.CompileExpression()
	if err != nil {
		return 0
	}

	exprtkObj.SetDoubleVariableValue("x", 8)
	return int(exprtkObj.GetEvaluatedValue())

}
