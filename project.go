package main

import (
	SY "MathExpressionEvaluator_/ShuntingYardAlgorithm"
	trig "MathExpressionEvaluator_/trig"
	"fmt"
	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
	"os"
	"regexp"
)

func main() {
	//fmt.Println("enter the expression you want to compute")
	//expression := flag.String("math", "", "the expression you want to compute goes here")
	//if len(os.Args) < 2 {
	//	fmt.Println("expected 'math' subcommands")
	//	os.Exit(1)
	//}
	//flag.Parse()

	expression := "cos(1)"
	onlyParenthese := regexp.MustCompile(`[^\(\)\{\}\[\]]`).ReplaceAllString(expression, "")
	if !SY.IsValid(onlyParenthese) {
		fmt.Println("Mismatch in parentheses1")
		os.Exit(1)
	}

	tokens := regexp.MustCompile("[ ]+").Split(expression, -1)
	//for _, token := range tokens {
	//	fmt.Println(token)
	//}
	RPN := *ShuntingYard(tokens)
	//for RPN.Len() > 0 {
	//	println(RPN.Dequeue().(string))
	//}

	x := SY.RPN_Evaluator(RPN)
	fmt.Println(x)
	println(trig.Cos(100))
	//
	//if os.Args[1] == "math" {
	//	fmt.Println(*expression)
	//}

}

func ShuntingYard(tokens []string) *queue.Queue {
	precedence := make(map[string]int)
	precedence["+"] = 1
	precedence["-"] = 1
	precedence["*"] = 2
	precedence["/"] = 2
	operator := stack.New()
	output := queue.New()
	for _, o1 := range tokens {
		if SY.TokenIsNumber(o1) {
			output.Enqueue(o1)
		}
		if SY.TokenIsFunction(o1) {
			operator.Push(o1)
		}
		if SY.TokenIsOperator(o1) {
			for operator.Len() > 0 {
				if operator.Peek().(string) != "(" && (precedence[operator.Peek().(string)] > precedence[o1] || (precedence[operator.Peek().(string)] == precedence[o1] && SY.IsLeftAssossiative(o1))) {
					output.Enqueue(operator.Pop())
				} else {
					break
				}
			}
			operator.Push(o1)
		}
		if o1 == "(" {
			operator.Push(o1)
		}
		if o1 == ")" {
			for operator.Peek() != "(" {
				if operator.Len() > 0 {
					output.Enqueue(operator.Pop())
				} else {
					fmt.Println("Mismatch in parentheses2")
				}
			}
			operator.Pop()
			if operator.Len() > 0 && SY.TokenIsFunction(operator.Peek().(string)) {
				output.Enqueue(operator.Pop())
			}
		}
	}

	for operator.Len() > 0 {
		if operator.Peek() == "(" {
			fmt.Println("Mismatch in parentheses3")
		}
		output.Enqueue(operator.Pop())
	}

	return output
}
