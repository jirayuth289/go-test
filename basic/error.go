package basic

import (
	"errors"
	"fmt"
)

func devide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot devide by zero")
	}

	return a / b, nil
}

func testError() {
	result, err := devide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Result:", result)
}
