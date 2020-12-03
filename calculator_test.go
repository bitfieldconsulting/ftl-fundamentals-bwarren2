package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
)

func TestAddMulti(t *testing.T) {
	testcases := []struct {
		a, b float64
		rest []float64
		want float64
		name string
	}{
		{name: "Add two units", a: 0, b: 0, rest: []float64{}, want: 0},
		{name: "Add two nonunits", a: 1, b: 2, rest: []float64{}, want: 3},
		{name: "Add unit and nonunit", a: 0, b: 2, rest: []float64{}, want: 2},
		{name: "Add negative and inverse", a: -2, b: 2, rest: []float64{}, want: 0},
		{name: "Add negative and positive", a: -2, b: 3, rest: []float64{}, want: 1},
		{name: "Add two numbers", a: 1, b: 2, rest: []float64{}, want: 3},
		{name: "Add three numbers", a: 1, b: 2, rest: []float64{3}, want: 6},
		{name: "Add four numbers", a: 1, b: 2, rest: []float64{3, 4}, want: 10},
		{name: "Add five numbers", a: -1, b: 2, rest: []float64{-3, 4, -5}, want: -3},
	}
	for _, testcase := range testcases {
		got := calculator.Add(testcase.a, testcase.b, testcase.rest...)
		if testcase.want != got {
			t.Errorf("Wanted %v, got %v in %v", testcase.want, got, testcase.name)
		}
	}
}
func TestAddRandom(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		a, b := rand.Float64(), rand.Float64()
		want := a + b
		got := calculator.Add(a, b)
		if want != got {
			t.Errorf("Want %f, got %f, using %f + %f", want, got, a, b)
		}
	}
}

func TestSubtractMulti(t *testing.T) {
	testcases := []struct {
		a, b       float64
		rest       []float64
		want       float64
		expectsErr bool
		name       string
	}{
		{name: "Subtract two positives", a: 4, b: 2, rest: []float64{}, want: 2},
		{name: "Subtract positive and zero", a: 2, b: 0, rest: []float64{}, want: 2},
		{name: "Subtract zero and positive", a: 0, b: 2, rest: []float64{}, want: -2},
		{name: "Subtract two negatives", a: -1, b: -2, rest: []float64{}, want: 1},
		{name: "Subtract two numbers", a: 1, b: 2, rest: []float64{}, want: -1},
		{name: "Subtract three numbers", a: 1, b: 2, rest: []float64{3}, want: -4},
		{name: "Subtract four numbers", a: 1, b: 2, rest: []float64{3, 4}, want: -8},
		{name: "Subtract five numbers", a: -1, b: 2, rest: []float64{-3, 4, -5}, want: 1},
	}
	for _, testcase := range testcases {
		got := calculator.Subtract(testcase.a, testcase.b, testcase.rest...)
		if testcase.want != got {
			t.Errorf("Wanted %v, got %v in %v", testcase.want, got, testcase.name)
		}
	}
}

func TestSubtractRandom(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		a, b := rand.Float64(), rand.Float64()
		want := a - b
		got := calculator.Subtract(a, b)
		if want != got {
			t.Errorf("Want %v, got %v, using %v - %v", want, got, a, b)
		}
	}
}

func TestMultiplyMulti(t *testing.T) {
	testcases := []struct {
		a, b, want float64
		rest       []float64
		expectsErr bool
		name       string
	}{
		{name: "Multiply two positives", a: 2, b: 2, rest: []float64{}, want: 4},
		{name: "Multiply positive and unit", a: 2, b: 1, rest: []float64{}, want: 2},
		{name: "Multiply positive and zero", a: 2, b: 0, rest: []float64{}, want: 0},
		{name: "Multiply zero and zero", a: 0, b: 0, rest: []float64{}, want: 0},
		{name: "Multiply zero and negative", a: 0, b: -1, rest: []float64{}, want: 0},
		{name: "Multiply positive and negative units", a: 1, b: -1, rest: []float64{}, want: -1},
		{name: "Multiply positive and negative unit", a: 2, b: -1, rest: []float64{}, want: -2},
		{name: "Multiply two negative units", a: -1, b: -1, rest: []float64{}, want: 1},
		{name: "Multiply negative and negative unit", a: -2, b: -1, rest: []float64{}, want: 2},
		{name: "Multiply two negatives", a: -2, b: -2, rest: []float64{}, want: 4},
		{name: "Multiply two numbers", a: 1, b: 2, rest: []float64{}, want: 2},
		{name: "Multiply three numbers", a: 1, b: 2, rest: []float64{3}, want: 6},
		{name: "Multiply four numbers", a: 1, b: 2, rest: []float64{3, 4}, want: 24},
		{name: "Multiply five numbers", a: -1, b: 2, rest: []float64{-3, 4, -5}, want: -120},
	}
	for _, testcase := range testcases {
		got := calculator.Multiply(testcase.a, testcase.b, testcase.rest...)
		if testcase.want != got {
			t.Errorf("Wanted %v, got %v in %v", testcase.want, got, testcase.name)
		}
	}
}

func TestDivideMulti(t *testing.T) {
	testcases := []struct {
		a, b       float64
		rest       []float64
		want       float64
		expectsErr bool
		name       string
	}{
		{name: "Divide ones", a: 1, b: 1, rest: []float64{}, want: 1},
		{name: "Divide 2 by 1", a: 2, b: 1, rest: []float64{}, want: 2},
		{name: "Divide 1 by 2", a: 1, b: 2, rest: []float64{}, want: 0.5},
		{name: "Divide positive by zero", a: 1, b: 0, expectsErr: true},
		{name: "Divide positive by zero", a: -1, b: 0, expectsErr: true},
		{name: "Divide negative by zero", a: 0, b: 0, expectsErr: true},
		{name: "Divide negative by zero", a: -1, b: -1, rest: []float64{}, want: 1},
		{name: "Divide positive by negative", a: 2, b: -1, rest: []float64{}, want: -2},
		{name: "Divide 2 numbers", a: 1, b: 2, rest: []float64{}, want: .5},
		{name: "Divide 3 numbers", a: 1, b: 2, rest: []float64{.5}, want: 1},
		{name: "Given example", a: 12, b: 4, rest: []float64{3}, want: 1},
		{name: "Given example", a: 12, b: 4, rest: []float64{0}, expectsErr: true},
	}
	for _, testcase := range testcases {
		got, err := calculator.Divide(testcase.a, testcase.b, testcase.rest...)
		errorReceived := err != nil
		if testcase.expectsErr != errorReceived {
			t.Fatalf("unexpected error status: err %v in %v", err, testcase.name)
		}
		if !testcase.expectsErr && testcase.want != got {
			t.Errorf("Wanted %v, got %v in %v", testcase.want, got, testcase.name)
		}
	}
}

func TestSqrt(t *testing.T) {
	testcases := []struct {
		in, want   float64
		expectsErr bool
		name       string
	}{
		{name: "Sqrt 4", in: 4, want: 2},
		{name: "Sqrt 0", in: 0, want: 0},
		{name: "Sqrt 1", in: 1, want: 1},
		{name: "Sqrt .25", in: .25, want: .5},
		{name: "Sqrt small negative number", in: -0.000000001, expectsErr: true},
		{name: "Sqrt negative number", in: -1, expectsErr: true},
	}
	for _, testcase := range testcases {
		got, err := calculator.Sqrt(testcase.in)
		errorReceived := err != nil
		if testcase.expectsErr != errorReceived {
			t.Fatalf("unexpected error status: err %v in %v", err, testcase.name)
		}
		if !testcase.expectsErr && testcase.want != got {
			t.Errorf("Wanted %v, got %v in %v", testcase.want, got, testcase.name)
		}
	}
}

func TestEval(t *testing.T) {
	testcases := []struct {
		in         string
		want       float64
		expectsErr bool
	}{
		{in: "1+2", want: 3},
		{in: "2+2", want: 4},
		{in: "2-3", want: -1},
		{in: "2*3", want: 6},
		{in: "3/2", want: 1.5},
		{in: "1  +    2", want: 3},
		{in: "2 + 2", want: 4},
		{in: "2- 3", want: -1},
		{in: "2     *      3", want: 6},
		{in: "3/   2", want: 1.5},
		{in: "3/", expectsErr: true},
		{in: "3*", expectsErr: true},
		{in: "3+", expectsErr: true},
		{in: "3-", expectsErr: true},
		{in: "/", expectsErr: true},
		{in: "*", expectsErr: true},
		{in: "+", expectsErr: true},
		{in: "-", expectsErr: true},
		{in: " /", expectsErr: true},
		{in: "* ", expectsErr: true},
		{in: " + ", expectsErr: true},
		{in: "  -  ", expectsErr: true},
		{in: "2^2", expectsErr: true},
		{in: "2**3", expectsErr: true},
	}

	for _, testcase := range testcases {
		got, err := calculator.EvalExpr(testcase.in)
		errorReceived := err != nil
		if testcase.expectsErr != errorReceived {
			t.Fatalf("unexpected error status: err %v given %v", err, testcase.in)
		}
		if !testcase.expectsErr && testcase.want != got {
			t.Errorf("Wanted %v, got %v in %v", testcase.want, got, testcase.in)
		}

	}
}
