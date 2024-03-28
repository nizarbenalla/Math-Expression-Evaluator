package ShuntingYardAlgorithm

import (
	"MathExpressionEvaluator_/trig"
	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
	"strconv"
	"unicode"
)

func RPN_Evaluator(queue queue.Queue) float64 {
	if queue.Len() == 0 {
		return 0
	}
	stack := stack.New()
	for queue.Len() > 0 {
		if TokenIsNumber(queue.Peek().(string)) {
			stack.Push(queue.Dequeue())
		} else if TokenIsFunction(queue.Peek().(string)) {
			value, _ := strconv.ParseFloat(stack.Pop().(string), 64)

			switch queue.Peek().(string) {
			case "cos":
				stack.Push(trig.Cos(value))
			case "sin":
				stack.Push(trig.Sin(value))
			case "tan":
				stack.Push(trig.Tan(value))
			case "atan":
				stack.Push(trig.Atan(value))
			case "sqrt":
				stack.Push(trig.Sqrt(value))
			case "pow":
				exponent, _ := strconv.Atoi(stack.Pop().(string))
				stack.Push(trig.Pow(value, exponent))

			}
			queue.Dequeue()

		} else if TokenIsOperator(queue.Peek().(string)) {
			value1, _ := strconv.ParseFloat(stack.Pop().(string), 64)
			value2, _ := strconv.ParseFloat(stack.Pop().(string), 64)
			switch queue.Peek().(string) {
			case "+":
				stack.Push(value1 + value2)
			case "-":
				stack.Push(value2 - value1)
			case "/":
				stack.Push(value2 / value1)
			case "*":
				stack.Push(value1 * value2)
			}
			queue.Dequeue()

		}

	}
	return stack.Pop().(float64)
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
