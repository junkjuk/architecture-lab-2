package lab2

import (
	"fmt"
	"strconv"
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
		} else if isNumber(parts[i]) {
			stack = append(stack, parts[i])
		} else {
			return "", fmt.Errorf("incorrect prefix expression")
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

func isNumber(value string) bool {
	_, err := strconv.Atoi(value)
	return err == nil
}
