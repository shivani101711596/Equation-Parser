package eqnparse

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseEquation parses the given equation and returns it as an Equation struct.
// If the equation is invalid or cannot be parsed, an error is returned.
func ParseEquation(equation string) (*Equation, error) {
	// Remove whitespace from the equation
	equation = strings.ReplaceAll(equation, " ", "")

	// Split the equation into the LHS and RHS
	parts := strings.Split(equation, "=")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid equation format")
	}

	// Parse the LHS and RHS expressions
	lhs, err := parseExpression(parts[0])
	if err != nil {
		return nil, fmt.Errorf("invalid equation format: %v", err)
	}

	rhs, err := parseExpression(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid equation format: %v", err)
	}

	// Create the Equation struct
	return &Equation{
		lhs: lhs,
		rhs: rhs,
	}, nil
}

// parseExpression parses a mathematical expression.
func parseExpression(expr string) (Expression, error) {
	var numbers []int
	var operators []Operator

	// Iterate over the string and build the expression
	var numStr string
	for _, c := range expr {
		if isOperator(c) {
			// Parse the number and add it to the list
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return Expression{}, fmt.Errorf("invalid number: %s", numStr)
			}
			numbers = append(numbers, num)

			// Add the operator to the list
			op := Operator(c)
			if !isValidOperator(op) {
				return Expression{}, fmt.Errorf("invalid operator: %c", c)
			}
			operators = append(operators, op)

			// Reset the number string
			numStr = ""
		} else {
			// Add the character to the number string
			numStr += string(c)
		}
	}

	// Parse the last number in the expression
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return Expression{}, fmt.Errorf("invalid number: %s", numStr)
	}
	numbers = append(numbers, num)

	// Create the expression struct
	return Expression{
		Numbers:   numbers,
		Operators: operators,
	}, nil
}

// isOperator returns true if the given rune is a valid operator.
func isOperator(c rune) bool {
	return c == '+' || c == '-' || c == '*' || c == '/'
}

// isValidOperator returns true if the given Operator is valid.
func isValidOperator(op Operator) bool {
	for _, validOp := range ValidOperators {
		if op == validOp {
			return true
		}
	}
	return false
}
