package main

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func main() {
	//fmt.Println("enter the expression you want to compute")
	//expression := flag.String("math", "", "the expression you want to compute goes here")
	//if len(os.Args) < 2 {
	//	fmt.Println("expected 'math' subcommands")
	//	os.Exit(1)
	//}
	//flag.Parse()

	expression := "3 *  ( 32 + 3 )"

	onlyParenthese := regexp.MustCompile(`[^\(\)\{\}\[\]]`).ReplaceAllString(expression, "")

	if !isValid(onlyParenthese) {
		fmt.Println("Mismatch in parentheses1")
		os.Exit(1)
	}

	tokens := regexp.MustCompile("[ ]+").Split(expression, -1)

	for _, token := range tokens {
		fmt.Println(token)
	}
	RPN := *ShuntingYard(tokens)
	for RPN.Len() > 0 {
		v := RPN.Dequeue()
		fmt.Print(v)
	}

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
		if tokenIsNumber(o1) {
			output.Enqueue(o1)
		}
		if tokenIsFunction(o1) {
			operator.Push(o1)
		}
		if tokenIsOperator(o1) {
			if operator.Len() > 0 {
				for o2 := operator.Peek().(string); o2 != "(" && (precedence[o2] > precedence[o1] || (precedence[o2] == precedence[o1] && isLeftAssossiative(o1))); {
					output.Enqueue(operator.Pop())
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
			if operator.Len() > 0 && tokenIsFunction(operator.Peek().(string)) {
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

func isLeftAssossiative(o1 string) bool {
	leftAssossiative := make(map[string]bool)
	leftAssossiative["/"] = true
	leftAssossiative["-"] = true
	leftAssossiative["+"] = true
	leftAssossiative["*"] = true
	// not dealing with ^ which is right associative because we will implement pow
	if ok := leftAssossiative[o1]; ok {
		return true
	}
	return false

}

func tokenIsOperator(c string) bool {
	operators := make(map[string]bool)
	operators["+"] = true
	operators["*"] = true
	operators["/"] = true
	operators["-"] = true
	if ok := operators[c]; ok {
		return true
	}
	return false
}

func tokenIsFunction(c string) bool {
	functions := make(map[string]bool)
	functions["cos"] = true
	functions["acos"] = true
	functions["sin"] = true
	functions["asin"] = true
	functions["tan"] = true
	functions["atan"] = true
	functions["sqrt"] = true
	functions["pow"] = true
	if ok := functions[c]; ok {
		return true
	}
	return false
}

func tokenIsNumber(c string) bool {
	_, err := strconv.ParseFloat(c, 64)
	if err == nil {
		return true
	} else if unicode.IsDigit(rune(c[0])) {
		return true
	}
	return false
}

func RPN_Evaluator(stack *stack.Stack) float64 {

	return 0
}

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
func atan(number string) float64 {
	return 0
}
func sqrt(number string) float64 {
	return 0
}
func pow(number string) float64 {
	return 0
}
