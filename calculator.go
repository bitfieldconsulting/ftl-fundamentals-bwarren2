// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"fmt"
	"io"
	"math"
	"regexp"
)

// Add takes 2+ numbers and returns the result of adding them together.
func Add(a, b float64, rest ...float64) float64 {
	result := a + b
	for _, item := range rest {
		result += item
	}
	return result
}

// Subtract takes 2+ numbers and subtracts them together.
func Subtract(a float64, b float64, rest ...float64) float64 {
	result := a - b
	for _, item := range rest {
		result -= item
	}
	return result
}

// Multiply multiplies 2+ numbers.
func Multiply(a float64, b float64, rest ...float64) float64 {
	result := a * b
	for _, item := range rest {
		result *= item
	}
	return result
}

// Divide divides 2+ numbers or returns an error if it would divide by 0.
func Divide(a float64, b float64, rest ...float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("can't divide by zero")
	}
	result := a / b
	for _, item := range rest {
		if item == 0 {
			return 0, errors.New("can't divide by zero")
		}
		result /= item
	}
	return result, nil
}

// Sqrt takes a nonnegative number and returns its square root or returns an error.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("can't take the square root of a negative number: %f", a)
	}
	return math.Sqrt(a), nil
}

var validExprRe = regexp.MustCompile(`\d*(\.\d)?\s*[-+\/*]\s*\d*(\.\d)?$`)
var operatorRe = regexp.MustCompile(`[-+\/*]`)

// EvalExpr evaluates simple arithmetic expressions of "float64 operator float64"
func EvalExpr(in string) (float64, error) {

	var left, right float64
	var operator string

	_, err := fmt.Sscanf(in, "%v%v%v", &left, &operator, &right)
	if err != nil && err != io.EOF {
		fmt.Print(err)
		panic(err)
	}
	switch operator {
	case "+":
		return left + right, nil
	case "-":
		return left - right, nil
	case "*":
		return left * right, nil
	case "/":
		return left / right, nil
	default:
		return 0, fmt.Errorf("Could not match the given operator: %v", operator)
	}
}
