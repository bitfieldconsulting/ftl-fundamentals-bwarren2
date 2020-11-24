package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
)

type BinaryOpTestCase struct {
	in1, in2, want float64
	expectsErr     bool
	name           string
}

type UnaryOpTestCase struct {
	in, want   float64
	expectsErr bool
	name       string
}

type MultiOpTestCase struct {
	in         []float64
	want       float64
	expectsErr bool
	name       string
}

func TestAdd(t *testing.T) {
	t.Parallel()
	tests := []BinaryOpTestCase{
		{in1: 0, in2: 0, want: 0, name: "Add two units"},
		{in1: 1, in2: 2, want: 3, name: "Add two nonunits"},
		{in1: 0, in2: 2, want: 2, name: "Add unit and nonunit"},
		{in1: -2, in2: 2, want: 0, name: "Add negative and inverse"},
		{in1: -2, in2: 3, want: 1, name: "Add negative and positive"},
	}
	for _, testcase := range tests {
		got := calculator.Add(testcase.in1, testcase.in2)
		if testcase.want != got {
			t.Errorf("want %f, got %f in %v", testcase.want, got, testcase.name)
		}

	}
}

func TestAddRandom(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		in1, in2 := rand.Float64(), rand.Float64()
		want := in1 + in2
		got := calculator.Add(in1, in2)
		if want != got {
			t.Errorf("Want %f, got %f, using %f + %f", want, got, in1, in2)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	tests := []BinaryOpTestCase{
		{in1: 4, in2: 2, want: 2, name: "Subtract two positives"},
		{in1: 2, in2: 0, want: 2, name: "Subtract positive and zero"},
		{in1: 0, in2: 2, want: -2, name: "Subtract zero and positive"},
		{in1: -1, in2: -2, want: 1, name: "Subtract two negatives"},
	}
	for _, testcase := range tests {

		got, err := calculator.Subtract(testcase.in1, testcase.in2)
		if err != nil {
			t.Errorf("Got an unexpected error using %v - %v", testcase.in1, testcase.in2)
		}
		if testcase.want != got {
			t.Errorf("want %f, got %f in %v", testcase.want, got, testcase.name)
		}
	}
}

func TestSubtractMulti(t *testing.T) {
	testcases := []MultiOpTestCase{
		{in: []float64{1}, expectsErr: true, name: "Need 2+ operands"},
		{in: []float64{1, 2}, want: -1, name: "Subtract two numbers"},
		{in: []float64{1, 2, 3}, want: -4, name: "Subtract three numbers"},
		{in: []float64{1, 2, 3, 4}, want: -8, name: "Subtract four numbers"},
		{in: []float64{-1, 2, -3, 4, -5}, want: 1, name: "Subtract five numbers"},
	}
	for _, testcase := range testcases {
		got, err := calculator.Subtract(testcase.in...)
		switch {
		case err != nil && !testcase.expectsErr:
			t.Errorf("Did not expect an error but got one in %v", testcase.name)
		case err == nil && testcase.expectsErr:
			t.Errorf("Expected an error and did NOT get one in %v", testcase.name)
		case err == nil && !testcase.expectsErr:
			if testcase.want != got {
				t.Errorf("Wanted %v, got %v in %v", testcase.want, got, testcase.name)
			}
		}
	}
}
func TestSubtractRandom(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		in1, in2 := rand.Float64(), rand.Float64()
		want := in1 - in2
		got, err := calculator.Subtract(in1, in2)
		if err != nil {
			t.Errorf("Got an unexpected error using %v - %v", in1, in2)
		}
		if want != got {
			t.Errorf("Want %v, got %v, using %v - %v", want, got, in1, in2)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	tests := []BinaryOpTestCase{
		{in1: 2, in2: 2, want: 4, name: "Multiply two positives"},
		{in1: 2, in2: 1, want: 2, name: "Multiply positive and unit"},
		{in1: 2, in2: 0, want: 0, name: "Multiply positive and zero"},
		{in1: 0, in2: 0, want: 0, name: "Multiply zero and zero"},
		{in1: 0, in2: -1, want: 0, name: "Multiply zero and negative"},
		{in1: 1, in2: -1, want: -1, name: "Multiply positive and negative units"},
		{in1: 2, in2: -1, want: -2, name: "Multiply positive and negative unit"},
		{in1: -1, in2: -1, want: 1, name: "Multiply two negative units"},
		{in1: -2, in2: -1, want: 2, name: "Multiply negative and negative unit"},
		{in1: -2, in2: -2, want: 4, name: "Multiply two negatives"},
	}
	for _, testcase := range tests {
		got, err := calculator.Multiply(testcase.in1, testcase.in2)
		switch {
		case err != nil && !testcase.expectsErr:
			t.Errorf("Did not expect an error but got one in %v", testcase.name)
		case err == nil && testcase.expectsErr:
			t.Errorf("Expected an error and did NOT get one in %v", testcase.name)
		case err == nil && !testcase.expectsErr:
			if testcase.want != got {
				t.Errorf("Wanted %v, got %v in %v", testcase.want, got, testcase.name)
			}
		}
	}
}

func TestMultiplyMulti(t *testing.T) {
	testcases := []MultiOpTestCase{
		{in: []float64{1}, expectsErr: true, name: "Need 2+ operands"},
		{in: []float64{1, 2}, want: 2, name: "Multiply two numbers"},
		{in: []float64{1, 2, 3}, want: 6, name: "Multiply three numbers"},
		{in: []float64{1, 2, 3, 4}, want: 24, name: "Multiply four numbers"},
		{in: []float64{-1, 2, -3, 4, -5}, want: -120, name: "Multiply five numbers"},
	}
	for _, testcase := range testcases {
		got, err := calculator.Multiply(testcase.in...)
		switch {
		case err != nil && !testcase.expectsErr:
			t.Errorf("Did not expect an error but got one in %v", testcase.name)
		case err == nil && testcase.expectsErr:
			t.Errorf("Expected an error and did NOT get one in %v", testcase.name)
		case err == nil && !testcase.expectsErr:
			if testcase.want != got {
				t.Errorf("Wanted %v, got %v in %v", testcase.want, got, testcase.name)
			}
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	tests := []BinaryOpTestCase{
		{in1: 1, in2: 1, want: 1, name: "Divide ones"},
		{in1: 2, in2: 1, want: 2, name: "Divide 2 by 1"},
		{in1: 1, in2: 2, want: 0.5, name: "Divide 1 by 2"},
		{in1: 1, in2: 0, expectsErr: true, name: "Divide positive by zero"},
		{in1: -1, in2: 0, expectsErr: true, name: "Divide negative by zero"},
		{in1: 0, in2: 0, expectsErr: true, name: "Divide negative by zero"},
		{in1: -1, in2: -1, want: 1, name: "Divide negative by zero"},
		{in1: 2, in2: -1, want: -2, name: "Divide positive by negative"},
	}
	for _, testcase := range tests {
		got, err := calculator.Divide(testcase.in1, testcase.in2)
		switch {
		case err != nil && !testcase.expectsErr:
			t.Errorf("Did not expect an error but got one in %v", testcase.name)
		case err == nil && testcase.expectsErr:
			t.Errorf("Expected an error and did NOT get one in %v", testcase.name)
		case err == nil && !testcase.expectsErr:
			if testcase.want != got {
				t.Errorf("Wanted %v, got %v in %v", testcase.want, got, testcase.name)
			}
		}
	}
}

func TestDivideMulti(t *testing.T) {
	testcases := []MultiOpTestCase{
		{in: []float64{1, 2}, want: .5, name: "Divide 2 numbers"},
		{in: []float64{1, 2, .5}, want: 1, name: "Divide 3 numbers"},
		{in: []float64{12, 4, 3}, want: 1, name: "Given example"},
		{in: []float64{1}, expectsErr: true, name: "Can't divide a single number"},
	}
	for _, testcase := range testcases {
		got, err := calculator.Divide(testcase.in...)
		switch {
		case err != nil && !testcase.expectsErr:
			t.Errorf("Did not expect an error but got one in %v", testcase.name)
		case err == nil && testcase.expectsErr:
			t.Errorf("Expected an error and did NOT get one in %v", testcase.name)
		case err == nil && !testcase.expectsErr:
			if testcase.want != got {
				t.Errorf("Wanted %v, got %v in %v", testcase.want, got, testcase.name)
			}

		}
	}
}

func TestSqrt(t *testing.T) {
	testcases := []UnaryOpTestCase{
		{in: 4, want: 2, name: "Sqrt 4"},
		{in: 0, want: 0, name: "Sqrt 0"},
		{in: 1, want: 1, name: "Sqrt 1"},
		{in: .25, want: .5, name: "Sqrt .25"},
		{in: -0.000000001, expectsErr: true, name: "Sqrt small negative number"},
		{in: -1, expectsErr: true, name: "Sqrt negative number"},
	}
	for _, testcase := range testcases {
		got, err := calculator.Sqrt(testcase.in)
		switch {
		case err != nil && !testcase.expectsErr:
			t.Errorf("Got an error without expecting one for sqrt %v", testcase.in)
		case err == nil && testcase.expectsErr:
			t.Errorf("Did not get an error but expected one for sqrt %v", testcase.in)
		case err == nil && !testcase.expectsErr:
			if testcase.want != got {
				t.Errorf("Got %v, wanted %v, using sqrt %v", got, testcase.want, testcase.in)
			}
		}
	}
}
