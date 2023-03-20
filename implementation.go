package lab2

import (
	"fmt"
	"strings"
)

// PrefixToPostfix converts
func PrefixToPostfix(input string) (string, error) {
	var parts = strings.Split(input, " ")
	var stack []string

	for i := len(parts) - 1; i >= 0; i-- {
		if isOperator(parts[i]) {
			if len(stack) < 2 {
				return "", fmt.Errorf("incorrect prefix expression")
			}

			var firstOperand = stack[len(stack)-1]
			var secondOperand = stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result = firstOperand + " " + secondOperand + " " + parts[i]
			stack = append(stack, result)
		} else {
			stack = append(stack, parts[i])
		}
	}
	return strings.Join(stack, " "), nil
}

// check is symbol is oeprator
func isOperator(value string) bool {
	switch value {
	case "+", "-", "*", "/":
		return true
	default:
		return false
	}
}
