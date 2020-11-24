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
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(nums ...float64) (float64, error) {
	if len(nums) < 2 {
		return 0, errors.New("Please call with at least 2 args")
	}
	if len(nums) == 2 {
		return nums[0] - nums[1], nil
	}
	result, err := Subtract(nums[:len(nums)-1]...)
	if err != nil {
		return 0, err // How can I mock to test this?
	}
	return result - nums[len(nums)-1], nil
}

// Multiply multiplies numbers
func Multiply(nums ...float64) (float64, error) {
	if len(nums) < 2 {
		return 0, errors.New("Please call with at least 2 args")
	}
	if len(nums) == 2 {
		return nums[0] * nums[1], nil
	}
	result, err := Multiply(nums[:len(nums)-1]...)
	if err != nil {
		return 0, err // How can I mock to test this?
	}
	return result * nums[len(nums)-1], nil

}

// Divide divides 2+ numbers or returns an error
func Divide(nums ...float64) (float64, error) {
	if len(nums) < 2 {
		return 0, errors.New("We need more operands")
	}
	if nums[len(nums)-1] == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	if len(nums) == 2 {
		return nums[0] / nums[1], nil
	}
	result, err := Divide(nums[:len(nums)-1]...)
	if err != nil {
		return 0, err // How can I mock to get coverage on this line?
	}
	return result / nums[len(nums)-1], nil
}

// Sqrt takes the square root of a number
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("Can't take the square root of a negative number")
	}
	return math.Pow(a, .5), nil
}

// EvalExpr evaluates simple arithmetic expressions of "float64 operator float64"
func EvalExpr(in string) (float64, error) {
	re := regexp.MustCompile(`\d*(\.\d)?\s*[-+\/*]\s*\d*(\.\d)?$`)
	if !re.MatchString(in) {
		return 0, errors.New("Could not parse input expression")
	}
	operatorRe := regexp.MustCompile(`[-+\/*]`)
	operator := operatorRe.FindAllString(in, -1)[0]
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
