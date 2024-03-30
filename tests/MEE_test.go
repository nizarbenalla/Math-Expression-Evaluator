package MEE_test

import (
	"MathExpressionEvaluator_/helper"
	"testing"
)

func TestSimpleAddition(t *testing.T) {
	input := "( 1.5 + 1 )"
	result := helper.EvaluationExpression(&input)
	expected := 2.5
	if result != expected {
		t.Errorf("Expected: %f, Got: %f", expected, result)
	}
}
