package calculator_test

import (
	"calculator"
	"testing"
)

type testCase struct {
	a, b        float64
	want        float64
	message     string
	errExpected bool
}

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 4, message: "two positive number adds results to a positive number"},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s Add(%f, %f): want %f, got %f",
				tc.message, tc.a, tc.b, tc.want, got)

		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 0},
		{a: 1, b: 1, want: 0},
		{a: 5, b: 0, want: 5},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}
	}

}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 1},
		{a: 5, b: 0, want: 0},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}
	}

}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name        string
		a, b        float64
		want        float64
		errExpected bool
	}
	testCases := []testCase{
		{name: "N divided by N is 1", a: 2, b: 2, want: 1, errExpected: false},
		{name: "One divided by two is one-half", a: 1, b: 2, want: 0.5, errExpected: false},
		{name: "Division by zero is not allowed", a: 5, b: 0, want: 999, errExpected: true},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := err != nil
		if errReceived != tc.errExpected {
			t.Fatalf("%s: Divide(%f, %f): unexpected error status: %v", tc.name, tc.a, tc.b, err)
		}
		if !tc.errExpected && tc.want != got {
			t.Errorf("%s: Divide(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}

}

func TestSqrt(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name        string
		input       float64
		want        float64
		errExpected bool
	}
	testCases := []testCase{
		{name: "Square root of 4 is 2", input: 4, want: 2, errExpected: false},
		{name: "Square root of one is one", input: 1, want: 1, errExpected: false},
		{name: "Square root of a negative number is not allowed (unless you're a physicist)", input: -1, want: 999, errExpected: true},
	}
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.input)
		errReceived := err != nil
		if errReceived != tc.errExpected {
			t.Fatalf("%s: Sqrt(%f): unexpected error status: %v", tc.name, tc.input, err)
		}
		if !tc.errExpected && tc.want != got {
			t.Errorf("%s: Sqrt(%f): want %f, got %f", tc.name, tc.input, tc.want, got)
		}
	}
}
