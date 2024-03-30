package helper

import (
	SY "MathExpressionEvaluator_/ShuntingYardAlgorithm"
	"flag"
	"fmt"
	"os"
	"regexp"
)

func EvaluationExpression(expression *string) float64 {
	tokens := regexp.MustCompile("[ ]+").Split(*expression, -1)
	rpn := *SY.ShuntingYard(tokens)
	answer := SY.RpnEvaluator(rpn)
	return answer
}

func ReadExpression() *string {
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
