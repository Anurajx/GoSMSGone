package main

import (
	"fmt"
	"errors"
)

type divideError struct {
	dividend float64
}

func (d divideError) Error() string{
	return fmt.Sprintf("cannot divide %v by zero",d.dividend)
}

// ?

// func divide(dividend, divisor float64) (float64, error) {
// 	if divisor == 0 {
// 		return 0, divideError{dividend: dividend}
// 	}
// 	return dividend / divisor, nil
// }

func divide(x, y float64) (float64, error) {
	if y == 0 {
		// ?
		var err error = errors.New("dividing by zero is forbidden")
		return 0.0,err
	}
	return x / y, nil
}

