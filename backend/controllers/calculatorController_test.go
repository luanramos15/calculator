package controllers

import (
	"testing"
)

func TestValidate(t *testing.T) {
	validExpressions := []string{
		"2+2",
		"5-3*4/2",
		"√16+(3.5^2)",
		"10.0/2.5-1",
	}

	invalidExpressions := []string{
		"2+2a",
		"5-3*$ 2",
		"√16+(3.5^2)!",
		"10.0/2.5-@1",
	}

	for _, expr := range validExpressions {
		if !validate(expr) {
			t.Errorf("Expected expression %s to be valid", expr)
		}
	}

	for _, expr := range invalidExpressions {
		if validate(expr) {
			t.Errorf("Expected expression %s to be invalid", expr)
		}
	}
}
