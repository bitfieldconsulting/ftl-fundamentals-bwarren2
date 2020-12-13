package calculator_test

import (
	"calculator"
	"math/rand"
	"reflect"
	"runtime"
	"testing"
	"time"
)

func TestRandomly(t *testing.T) {
	t.Parallel()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		a, b := rand.Float64(), rand.Float64()
		want := a + b
		got := calculator.Add(a, b)
		if want != got {
			t.Errorf("Want %f, got %f, using %f + %f", want, got, a, b)
		}
	}
	for i := 0; i < 100; i++ {

		a, b := rand.Float64(), rand.Float64()
		want := a - b
		got := calculator.Subtract(a, b)
		if want != got {
			t.Errorf("Want %v, got %v, using %v - %v", want, got, a, b)
		}
	}

}

func TestBasicOps(t *testing.T) {
	tcs := []struct {
		name       string
		fn         func(a float64, b float64, rest ...float64) float64
		a, b       float64
		rest       []float64
		want       float64
		expectsErr bool
	}{
		{name: "Add two units", fn: calculator.Add, a: 0, b: 0, want: 0},
		{name: "Add two nonunits", fn: calculator.Add, a: 1, b: 2, want: 3},
		{name: "Add unit and nonunit", fn: calculator.Add, a: 0, b: 2, want: 2},
		{name: "Add negative and inverse", fn: calculator.Add, a: -2, b: 2, want: 0},
		{name: "Add negative and positive", fn: calculator.Add, a: -2, b: 3, want: 1},
		{name: "Add two numbers", fn: calculator.Add, a: 1, b: 2, want: 3},
		{name: "Add three numbers", fn: calculator.Add, a: 1, b: 2, rest: []float64{3}, want: 6},
		{name: "Add four numbers", fn: calculator.Add, a: 1, b: 2, rest: []float64{3, 4}, want: 10},
		{name: "Add five numbers", fn: calculator.Add, a: -1, b: 2, rest: []float64{-3, 4, -5}, want: -3},
		{name: "Subtract two positives", fn: calculator.Subtract, a: 4, b: 2, want: 2},
		{name: "Subtract positive and zero", fn: calculator.Subtract, a: 2, b: 0, want: 2},
		{name: "Subtract zero and positive", fn: calculator.Subtract, a: 0, b: 2, want: -2},
		{name: "Subtract two negatives", fn: calculator.Subtract, a: -1, b: -2, want: 1},
		{name: "Subtract two numbers", fn: calculator.Subtract, a: 1, b: 2, want: -1},
		{name: "Subtract three numbers", fn: calculator.Subtract, a: 1, b: 2, rest: []float64{3}, want: -4},
		{name: "Subtract four numbers", fn: calculator.Subtract, a: 1, b: 2, rest: []float64{3, 4}, want: -8},
		{name: "Subtract five numbers", fn: calculator.Subtract, a: -1, b: 2, rest: []float64{-3, 4, -5}, want: 1},
		{name: "Multiply two positives", fn: calculator.Multiply, a: 2, b: 2, want: 4},
		{name: "Multiply positive and unit", fn: calculator.Multiply, a: 2, b: 1, want: 2},
		{name: "Multiply positive and zero", fn: calculator.Multiply, a: 2, b: 0, want: 0},
		{name: "Multiply zero and zero", fn: calculator.Multiply, a: 0, b: 0, want: 0},
		{name: "Multiply zero and negative", fn: calculator.Multiply, a: 0, b: -1, want: 0},
		{name: "Multiply positive and negative units", fn: calculator.Multiply, a: 1, b: -1, want: -1},
		{name: "Multiply positive and negative unit", fn: calculator.Multiply, a: 2, b: -1, want: -2},
		{name: "Multiply two negative units", fn: calculator.Multiply, a: -1, b: -1, want: 1},
		{name: "Multiply negative and negative unit", fn: calculator.Multiply, a: -2, b: -1, want: 2},
		{name: "Multiply two negatives", fn: calculator.Multiply, a: -2, b: -2, want: 4},
		{name: "Multiply two numbers", fn: calculator.Multiply, a: 1, b: 2, want: 2},
		{name: "Multiply three numbers", fn: calculator.Multiply, a: 1, b: 2, rest: []float64{3}, want: 6},
		{name: "Multiply four numbers", fn: calculator.Multiply, a: 1, b: 2, rest: []float64{3, 4}, want: 24},
		{name: "Multiply five numbers", fn: calculator.Multiply, a: -1, b: 2, rest: []float64{-3, 4, -5}, want: -120},
	}
	for _, tc := range tcs {
		got := tc.fn(tc.a, tc.b, tc.rest...)
		if tc.want != got {
			t.Errorf("%v(%f, %f, %v): want %v, got %v", runtime.FuncForPC(reflect.ValueOf(tc.fn).Pointer()).Name(), tc.a, tc.b, tc.rest, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	tcs := []struct {
		a, b       float64
		rest       []float64
		want       float64
		expectsErr bool
		name       string
	}{
		{name: "Divide ones", a: 1, b: 1, want: 1},
		{name: "Divide 2 by 1", a: 2, b: 1, want: 2},
		{name: "Divide 1 by 2", a: 1, b: 2, want: 0.5},
		{name: "Divide positive by zero", a: 1, b: 0, expectsErr: true},
		{name: "Divide positive by zero", a: -1, b: 0, expectsErr: true},
		{name: "Divide negative by zero", a: 0, b: 0, expectsErr: true},
		{name: "Divide negative by zero", a: -1, b: -1, want: 1},
		{name: "Divide positive by negative", a: 2, b: -1, want: -2},
		{name: "Divide 2 numbers", a: 1, b: 2, want: .5},
		{name: "Divide 3 numbers", a: 1, b: 2, rest: []float64{.5}, want: 1},
		{name: "Divide 3 numbers, 3rd zero", a: 1, b: 2, rest: []float64{0}, expectsErr: true},
		{name: "Given example", a: 12, b: 4, rest: []float64{3}, want: 1},
		{name: "Given example", a: 12, b: 4, rest: []float64{0}, expectsErr: true},
	}
	for _, tc := range tcs {
		got, err := calculator.Divide(tc.a, tc.b, tc.rest...)
		errorReceived := err != nil
		if tc.expectsErr != errorReceived {
			t.Fatalf("unexpected error status: err %v in %v", err, tc.name)
		}
		if !tc.expectsErr && tc.want != got {
			t.Errorf("Wanted %v, got %v in %v", tc.want, got, tc.name)
		}
	}
}

func TestSqrt(t *testing.T) {
	tcs := []struct {
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
	for _, tc := range tcs {
		got, err := calculator.Sqrt(tc.in)
		errorReceived := err != nil
		if tc.expectsErr != errorReceived {
			t.Fatalf("unexpected error status: err %v in %v", err, tc.name)
		}
		if !tc.expectsErr && tc.want != got {
			t.Errorf("Wanted %v, got %v in %v", tc.want, got, tc.name)
		}
	}
}

func TestEval(t *testing.T) {
	tcs := []struct {
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

	for _, tc := range tcs {
		got, err := calculator.EvalExpr(tc.in)
		errorReceived := err != nil
		if tc.expectsErr != errorReceived {
			t.Fatalf("unexpected error status: err %v given %v", err, tc.in)
		}
		if !tc.expectsErr && tc.want != got {
			t.Errorf("Wanted %v, got %v in %v", tc.want, got, tc.in)
		}

	}
}
