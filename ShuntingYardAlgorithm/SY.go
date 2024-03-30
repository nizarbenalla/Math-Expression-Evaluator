package ShuntingYardAlgorithm

import (
	"MathExpressionEvaluator_/trig"
	"fmt"
	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
	"strconv"
	"unicode"
)

func ShuntingYard(tokens []string) *queue.Queue {
	precedence := make(map[string]int)
	precedence["+"] = 1
	precedence["-"] = 1
	precedence["*"] = 2
	precedence["/"] = 2
	operator := stack.New()
	output := queue.New()
	for _, o1 := range tokens {
		if TokenIsNumber(o1) {
			output.Enqueue(o1)
		}
		if TokenIsFunction(o1) {
			operator.Push(o1)
		}
		if TokenIsOperator(o1) {
			for operator.Len() > 0 {
				if operator.Peek().(string) != "(" && (precedence[operator.Peek().(string)] > precedence[o1] || (precedence[operator.Peek().(string)] == precedence[o1] && IsLeftAssossiative(o1))) {
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
			if operator.Len() > 0 && TokenIsFunction(operator.Peek().(string)) {
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

func RpnEvaluator(queue queue.Queue) float64 {
	if queue.Len() == 0 {
		return 0
	}
	numbers := stack.New()
	for queue.Len() > 0 {
		if TokenIsNumber(queue.Peek().(string)) {
			x, _ := strconv.ParseFloat(queue.Dequeue().(string), 64)
			numbers.Push(x)
		} else if TokenIsFunction(queue.Peek().(string)) {
			value, _ := strconv.ParseFloat(numbers.Pop().(string), 64)

			switch queue.Peek().(string) {
			case "cos":
				numbers.Push(trig.Cos(value))
			case "acos":
				numbers.Push(trig.Acos(value))
			case "sin":
				numbers.Push(trig.Sin(value))
			case "asin":
				numbers.Push(trig.Asin(value))
			case "tan":
				numbers.Push(trig.Tan(value))
			case "atan":
				numbers.Push(trig.Atan(value))
			case "sqrt":
				numbers.Push(trig.Sqrt(value))
			case "pow":
				exponent, _ := strconv.Atoi(numbers.Pop().(string))
				numbers.Push(trig.Pow(value, exponent))

			}
			queue.Dequeue()

		} else if TokenIsOperator(queue.Peek().(string)) {

			value1 := numbers.Pop().(float64)
			value2 := numbers.Pop().(float64)
			switch queue.Peek().(string) {
			case "+":
				numbers.Push(value1 + value2)
			case "-":
				numbers.Push(value2 - value1)
			case "/":
				numbers.Push(value2 / value1)
			case "*":
				numbers.Push(value1 * value2)
			}
			queue.Dequeue()

		}

	}
	return numbers.Pop().(float64)
}
func IsLeftAssossiative(o1 string) bool {
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

func TokenIsOperator(c string) bool {
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

func TokenIsFunction(c string) bool {
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

func TokenIsNumber(c string) bool {
	_, err := strconv.ParseFloat(c, 64)
	if err == nil {
		return true
	} else if unicode.IsDigit(rune(c[0])) {
		return true
	}
	return false
}

func IsValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2 == 1 {
		return false
	}
	matchingPair := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	var pile []rune
	for _, r := range s {
		if _, ok := matchingPair[r]; ok {
			pile = append(pile, r)
		} else if len(pile) == 0 || matchingPair[pile[len(pile)-1]] != r {
			return false
		} else {
			pile = pile[:len(pile)-1]
		}
	}
	return len(pile) == 0
}
