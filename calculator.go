// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64, rest ...float64) float64 {
	result := a + b
	for _, item := range rest {
		result += item
	}

	return result
}

// Subtract takes 2+ numbers and returns the result of subtracting the second
// from the first.  For variadic args, evaluation is (a - b) - c ...
func Subtract(first float64, second float64, rest ...float64) float64 {
	result := first - second

	for _, item := range rest {
		result -= item
	}

	return result

}

// Multiply multiplies 2+ numbers.  For variadic args, evaluation is (a * b) * c ...
func Multiply(first float64, second float64, rest ...float64) float64 {
	result := first * second

	for _, item := range rest {
		result *= item
	}

	return result

}

// Divide divides 2+ numbers or returns an error if it would divide by 0.
// For variadic args, evaluation is (a / b) / c ...
func Divide(first float64, second float64, rest ...float64) (float64, error) {
	if second == 0 {
		return 0, errors.New("can't divide by zero")
	}
	result := first / second

	for _, item := range rest {
		if item == 0 {
			return 0, errors.New("can't divide by zero")
		}
		result /= item
	}

	return result, nil
}

// Sqrt takes the square root of a nonnegative number or returns an error.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("can't take the square root of a negative number")
	}
	return math.Sqrt(a), nil
}

var validExprRe = regexp.MustCompile(`\d*(\.\d)?\s*[-+\/*]\s*\d*(\.\d)?$`)
var operatorRe = regexp.MustCompile(`[-+\/*]`)

// EvalExpr evaluates simple arithmetic expressions of "float64 operator float64"
func EvalExpr(in string) (float64, error) {

	if !validExprRe.MatchString(in) {
		return 0, errors.New("could not parse input expression")
	}

	operatorMatch := operatorRe.FindAllString(in, -1)
	if len(operatorMatch) < 1 {
		return 0, errors.New("we couldn't get an operator")
	}
	operator := operatorMatch[0]
	parts := strings.Split(in, operator)

	left, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	if err != nil {
		return 0, err
	}
	right, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if err != nil {
		return 0, err
	}

	switch operator {
	case `+`:
		return left + right, nil
	case `-`:
		return left - right, nil
	case `*`:
		return left * right, nil
	case `/`:
		return left / right, nil
	default:
		return 0, errors.New("Could not match the given operator") // How could I mock to test this line?
	}
}
