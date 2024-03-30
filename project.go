package main

import (
	"MathExpressionEvaluator_/helper"
	"fmt"
)

func main() {
	expression := helper.ReadExpression()
	answer := helper.EvaluationExpression(expression)
	fmt.Println(answer)
}
