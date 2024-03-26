package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
	"os"
	"regexp"
	"strings"
)

func main() {
	//fmt.Println("enter the expression you want to compute")
	//expression := flag.String("math", "", "the expression you want to compute goes here")
	//if len(os.Args) < 2 {
	//	fmt.Println("expected 'math' subcommands")
	//	os.Exit(1)
	//}
	//flag.Parse()

	expression := "3 *  (2+3)"
	onlyParenthese := regexp.MustCompile(`[^\(\)\{\}\[\]]`).ReplaceAllString(expression, "")
	if !isValid(onlyParenthese) {
		fmt.Println("wrong parentheses")
		os.Exit(1)
	}

	ShuntingYard(expression)

	//
	//if os.Args[1] == "math" {
	//	fmt.Println(*expression)
	//}

}

func ShuntingYard(expression string) *stack.Stack {
	return nil
}

func handleMaths(math *string) int {
	var tokens []string = strings.Split(*math, "")

	for i, s := range tokens {
		fmt.Println(i)
		fmt.Println(s)
	}
	return 0
}
func RPN_Evaluator(stack *stack.Stack) float64 {

	return 0
}

//Shunting yard algorithm

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}
	matchingPair := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := []rune{}
	for _, r := range s {
		if _, ok := matchingPair[r]; ok {
			stack = append(stack, r)
		} else if len(stack) == 0 || matchingPair[stack[len(stack)-1]] != r {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func cos(number string) float64 {
	return 0
}
func acos(number string) float64 {
	return 0
}

func sin(number string) float64 {
	return 0
}
func asin(number string) float64 {
	return 0
}
func tan(number string) float64 {
	return 0
}
func sqrt(number string) float64 {
	return 0
}
func pow(number string) float64 {
	return 0
}
