package main

import (
	SY "MathExpressionEvaluator_/ShuntingYardAlgorithm"
	"flag"
	"fmt"
	"os"
	"regexp"
)

func main() {
	expression := readExpression()
	answer := evaluationExpression(expression)
	fmt.Println(answer)
}

func evaluationExpression(expression *string) float64 {
	tokens := regexp.MustCompile("[ ]+").Split(*expression, -1)
	rpn := *SY.ShuntingYard(tokens)
	answer := SY.RpnEvaluator(rpn)
	return answer
}

func readExpression() *string {
	fmt.Println("Your answer is :")
	expression := flag.String("math", "", "the expression you want to compute goes here")
	if len(os.Args) < 2 {
		fmt.Println("expected 'math' subcommands")
		os.Exit(1)
	}
	flag.Parse()
	onlyParenthese := regexp.MustCompile(`[^\(\)\{\}\[\]]`).ReplaceAllString(*expression, "")
	if !SY.IsValid(onlyParenthese) {
		fmt.Println("Mismatch in parentheses1")
		os.Exit(1)
	}
	return expression
}
