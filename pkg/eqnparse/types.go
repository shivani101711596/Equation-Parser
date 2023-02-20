package eqnparse

// DO NOT MODIFY THIS FILE

// Operator represents mathematical binary operators
type Operator rune

// These are the valid mathemematical operators
const (
	Addition       Operator = '+'
	Subtraction    Operator = '-'
	Multiplication Operator = '*'
	Division       Operator = '/'
)

// ValidOperators is the list of valid operators for an expression
var ValidOperators = []Operator{Addition, Subtraction, Multiplication, Division}

// Expression represents a mathematical expression
type Expression struct {
	Numbers   []int
	Operators []Operator
}

// Equation represents a mathematical equation
type Equation struct {
	lhs, rhs Expression
}
