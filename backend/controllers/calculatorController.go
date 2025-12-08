package controllers

import (
	calculator "Calculator/calculator"
	"encoding/json"
	"net/http"
	"strings"
)

const allowedChars string = "0123456789/*-+.^()√"

func CalculatorHandler(w http.ResponseWriter, r *http.Request) {
	e, error := parseRequest(r)
	if error != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if !validate(e.MathExpression) {
		http.Error(w, "Invalid characters in expression", http.StatusBadRequest)
		return
	}

	result, error := calculator.Calculate(e.MathExpression) // Didn't filter by the divide by 0 case because + or - inf is a valid result for it
	if error != nil {
		http.Error(w, "Invalid result", http.StatusBadRequest)
		return
	}

	output := calculator.Result{Result: result}

	jsonBytes, err := json.Marshal(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func parseRequest(r *http.Request) (calculator.Expression, error) {
	var e calculator.Expression
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
