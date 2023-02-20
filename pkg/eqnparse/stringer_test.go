package eqnparse

import (
	"fmt"
	"math/rand"
	"testing"
)

// DO NOT MODIFY THIS FILE

func ExampleEquation() {
	eqn := Equation{
		lhs: Expression{[]int{7, 2, 3}, []Operator{Addition, Subtraction}},
		rhs: Expression{Numbers: []int{6}},
	}
	fmt.Println(eqn)
	// Output: 7+2-3=6
}

func TestExpression_String(t *testing.T) {
	tests := []struct {
		name  string
		expr  Expression
		wantS string
	}{
		{
			"basic",
			Expression{[]int{7, 2, 3}, []Operator{Addition, Subtraction}},
			"7+2-3",
		},
		{
			"int",
			Expression{[]int{7}, []Operator{}},
			"7",
		},
		{
			"multiple-divide",
			Expression{[]int{7, 2, 3}, []Operator{Multiplication, Division}},
			"7*2/3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS := tt.expr.String()
			if gotS != tt.wantS {
				t.Errorf("Expected %s == %s", gotS, tt.wantS)
			}
		})
	}
}

func TestEquation_String(t *testing.T) {
	tests := []struct {
		name  string
		eqn   Equation
		wantS string
	}{
		{
			"basic",
			Equation{
				lhs: Expression{[]int{7, 2, 3}, []Operator{Addition, Subtraction}},
				rhs: Expression{[]int{6}, nil},
			},
			"7+2-3=6",
		},

		{
			"complex",
			Equation{
				lhs: Expression{[]int{9, 3, 4}, []Operator{Multiplication, Division}},
				rhs: Expression{[]int{5, 2}, []Operator{Addition}},
			},
			"9*3/4=5+2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS := tt.eqn.String()
			if gotS != tt.wantS {
				t.Errorf("Expected %s == %s", gotS, tt.wantS)
			}
		})
	}
}

func BenchmarkExpression_String(b *testing.B) {
	b.StopTimer()

	// generate an Expression
	size := 10
	expr := Expression{
		Numbers:   make([]int, size),
		Operators: make([]Operator, size-1),
	}
	for i := range expr.Numbers {
		expr.Numbers[i] = rand.Intn(10000)
	}
	for i := range expr.Operators {
		expr.Operators[i] = ValidOperators[rand.Intn(len(ValidOperators))]
	}

	for i := 0; i < b.N; i++ {
		b.StartTimer()
		s := expr.String()
		b.StopTimer()

		if s == "" {
			b.Errorf("Expected an equation but go nil")
		}
		b.SetBytes(int64(len(s)))
	}
}
