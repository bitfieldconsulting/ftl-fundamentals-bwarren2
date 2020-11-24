// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply multiplies numbers
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide divides two numbers or returns an error
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	return a / b, nil
}

// Sqrt takes the square root of a number
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("Can't take the square root of a negative number")
	}
	return math.Pow(a, .5), nil
}
