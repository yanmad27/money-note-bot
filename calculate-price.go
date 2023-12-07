package main

import "github.com/vjeantet/govaluate"

func calculatePrice(price string) int {
	expression, err := govaluate.NewEvaluableExpression(price)
	if err != nil {
		return 0
	}

	result, err := expression.Evaluate(nil)

	return int(result.(float64))
}
