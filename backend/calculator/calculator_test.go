package calculator

import (
	"fmt"
	"testing"
)

func TestSimpleOperations(t *testing.T) {
	tests := []struct {
		expression string
		expected   string
	}{
		{"2+2", "4"},
		{"5-3", "2"},
		{"4*3", "12"},
		{"8/4", "2"},
		{"10^2", "100"},
	}

	for _, test := range tests {
		result, err := Calculate(test.expression)
		if err != nil {
			t.Errorf("Expected no error for expression %s, got %v", test.expression, err)
		}
		if result != test.expected {
			t.Errorf("For expression %s, expected %s, got %s", test.expression, test.expected, result)
		}
	}
}

func TestCalculateDivideByZero(t *testing.T) {
	result, err := Calculate("5/0")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := "+Inf"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestCalculateSquareRoot(t *testing.T) {
	result, err := Calculate("√16")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := "4"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestCalculateInvalidExpression(t *testing.T) {
	_, err := Calculate("3+4+")
	if err == nil {
		t.Errorf("Expected an error for invalid expression, got nil")
	}
}

func TestTrimTrailingZeros(t *testing.T) {
	input := "12.340000"
	expected := "12.34"
	result := TrimTrailingZeros(input)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestGetFirstNRunes(t *testing.T) {
	input := "Hello, World!"
	n := 5
	expected := "Hello"
	result := GetFirstNRunes(input, n)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestGetFirstNRunesNegativeN(t *testing.T) {
	input := "Hello, World!"
	n := -3
	expected := ""
	result := GetFirstNRunes(input, n)
	if result != expected {
		t.Errorf("Expected empty string, got %s", result)
	}
}

func TestCalculateLargeExpression(t *testing.T) {
	expression := "1+2+3+4+5+6+7+8+9+10"
	result, err := Calculate(expression)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := "55"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestCalculateWithParentheses(t *testing.T) {
	expression := "(2+3)*4"
	result, err := Calculate(expression)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := "20"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestCalculateWithDecimals(t *testing.T) {
	expression := "5.5+2.3"
	result, err := Calculate(expression)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := "7.8"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestCalculateWithMultipleSquareRoots(t *testing.T) {
	expression := "√16+√9"
	result, err := Calculate(expression)
	fmt.Print(result)
	fmt.Print(err)
	if err != nil {
		t.Errorf("Expected no error, got √")
	}
	expected := "7"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestCalculateSquareRootWithParentheses(t *testing.T) {
	expression := "√(25+25)"
	result, err := Calculate(expression)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expected := "7.07106781"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
