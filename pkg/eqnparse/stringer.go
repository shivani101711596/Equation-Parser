package eqnparse

import (
	"strconv"
	"strings"
)

// STUDENTS: MODIFY THE IMPLEMENTATIONS TO IMPROVE THE BENCHMARK

func (expr Expression) String() string {
	var sb strings.Builder
	for i, n := range expr.Numbers {
		if i > 0 {
			sb.WriteRune(rune(expr.Operators[i-1]))
		}
		sb.WriteString(strconv.Itoa(n))
	}
	return sb.String()
}

func (eqn Equation) String() string {
	var sb strings.Builder
	sb.WriteString(eqn.lhs.String())
	sb.WriteRune('=')
	sb.WriteString(eqn.rhs.String())
	return sb.String()
}
