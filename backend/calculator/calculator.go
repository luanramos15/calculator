package calculator

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/expr-lang/expr"
)

func Calculate(expression string) (string, error) {
	if strings.Contains(expression, "√") {
		// Parentheses first
		reParen := regexp.MustCompile(`√\(([^()]+)\)`)
		expression = reParen.ReplaceAllString(expression, "($1)^(1/2)")

		// Then simple numbers/vars
		reSimple := regexp.MustCompile(`√([a-zA-Z0-9]+)`)
		expression = reSimple.ReplaceAllString(expression, "($1)^(1/2)")
	}

	program, err := expr.Compile(expression)
	if err != nil {
		return "", err
	}

	output, err := expr.Run(program, expr.AsFloat64())
	if err != nil {
		return "", err
	}

	var value string

	switch v := output.(type) {
	case int:
		value = strconv.Itoa(v)
	case float64:
		value =
			TrimTrailingZeros(
				GetFirstNRunes(
					strconv.FormatFloat(
						v,
						'f', 8, 64,
					),
					15,
				),
			)
	default:
		return "", err
	}

	return value, nil
}

func TrimTrailingZeros(s string) string {
	re := regexp.MustCompile(`\.?0+$`)
	return re.ReplaceAllString(s, "")
}

func GetFirstNRunes(s string, n int) string {
	if n < 0 {
		return ""
	}

	runes := []rune(s)
	if n >= len(runes) {
		return s
	}
	return string(runes[:n])
}
