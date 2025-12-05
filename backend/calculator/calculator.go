package calculator

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/expr-lang/expr"
)

const allowedChars = "0123456789/*-+.^()√"

func Calculate(expression string) (string, error) {
	if strings.Contains(expression, "√") {
		var re = regexp.MustCompile(`(?m)√\(?([^\)]*)\)?`)

		for _, match := range re.FindAllStringSubmatch(expression, -1) {
			expression = strings.Replace(expression, match[0], "("+match[1]+")^(1/2)", 1) //√ is the same as power of 1/2
		}
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

func Handler(w http.ResponseWriter, r *http.Request) {
	e, error := parseRequest(r)
	if error != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if !validate(e.MathExpression) {
		http.Error(w, "Invalid characters in expression", http.StatusBadRequest)
		return
	}

	result, error := Calculate(e.MathExpression) // Didn't filter by the divide by 0 case because + or - inf is a valid result for it
	if error != nil {
		http.Error(w, "Invalid result", http.StatusBadRequest)
		return
	}

	output := Result{Result: result}

	jsonBytes, err := json.Marshal(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func parseRequest(r *http.Request) (Expression, error) {
	var e Expression
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		return e, err
	}

	return e, nil
}

func validate(mathExpression string) bool {
	for _, r := range mathExpression {
		if !strings.ContainsRune(allowedChars, r) {
			return false
		}
	}

	return true
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
