package eqnparse

import (
	"fmt"
	"reflect"
	"testing"
)

// STUDENTS: YOU MUST IMPLEMENT ALL OF THE BELOW TESTS IN THIS FILE
// ExampleParseEquation
// TestParseEquation
// BenchmarkParseEquation
// FuzzParseEquation

/* Example equations. Note equations do have to be mathematically valid
3-1=2
5*4+2=22
3-1=2
9*2=18
6*2/3=4
3 - 1 = 2
9 + 0 = 9
2 * 3 = 6
4 + 5 = 9
4 + 5 + 6 = 15
2 + 3 + 1 = 6
7+3-3=7* 1
*/

func ExampleParseEquation() {
	equationStr := "7 + 2 /3 = 1*3"

	equation, err := ParseEquation(equationStr)
	if err != nil {
		fmt.Printf("Error parsing equation: %v\n", err)
		return
	}

	fmt.Printf("LHS Numbers: %v\n", equation.lhs.Numbers)
	fmt.Printf("LHS Operators: %v\n", equation.lhs.Operators)
	fmt.Printf("RHS Numbers: %v\n", equation.rhs.Numbers)
	fmt.Printf("RHS Operators: %v\n", equation.rhs.Operators)

	// Output:
	// LHS Numbers: [7 2]
	// LHS Operators: [+ /]
	// RHS Numbers: [1 3]
	// RHS Operators: [*]
}

func TestParseEquation(t *testing.T) {
	testCases := []struct {
		input string
		want  *Equation
		err   bool
	}{
		{
			input: "1+2=3",
			want: &Equation{
				lhs: Expression{
					Numbers:   []int{1, 2},
					Operators: []Operator{Addition},
				},
				rhs: Expression{
					Numbers: []int{3},
				},
			},
			err: false,
		},
		{
			input: "5 * 4 - 3 / 6 = 19",
			want: &Equation{
				lhs: Expression{
					Numbers:   []int{5, 4, 3, 6},
					Operators: []Operator{Multiplication, Subtraction, Division},
				},
				rhs: Expression{
					Numbers: []int{19},
				},
			},
			err: false,
		},
		{
			input: "1 + 2 - 3 *",
			want:  nil,
			err:   true,
		},
		{
			input: "1 + 2 = 3 = 4",
			want:  nil,
			err:   true,
		},
		{
			input: "1.5 + 2.0 = 3.5",
			want:  nil,
			err:   true,
		},
	}

	for _, tc := range testCases {
		got, err := ParseEquation(tc.input)

		if tc.err && err == nil {
			t.Errorf("ParseEquation(%q) expected error, but got nil", tc.input)
		} else if !tc.err && err != nil {
			t.Errorf("ParseEquation(%q) unexpected error: %v", tc.input, err)
		} else if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("ParseEquation(%q) = %v, want %v", tc.input, got, tc.want)
		}
	}
}

func BenchmarkParseEquation(b *testing.B) {
	// Define the equation to be parsed
	equation := "1 + 2 / 3 = 4 * 5"

	// Run the ParseEquation function b.N times
	for n := 0; n < b.N; n++ {
		_, err := ParseEquation(equation)
		if err != nil {
			b.Fatalf("unexpected error: %v", err)
		}
	}
}

func FuzzParseEquation(f *testing.F) {
	f.Fuzz(func(t *testing.T, input string) {
		_, err := ParseEquation(input)
		if err != nil {
			t.Logf("failed to parse equation %q: %v", input, err)
		}
	})
}
