package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello, Go!")

	// Example usage of doTask function
	err := doTask("")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Task executed successfully!")
	}

	// Testing the divide function
	testDivide(10, 2) // Expected: 5, nil
	testDivide(5, 0)  // Expected: 0, "division by zero" error
	testDivide(9, 3)  // Expected: 3, nil
}

func doTask(task string) error {
	if task == "" {
		return errors.New("task cannot be empty")
	}
	return nil
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func testDivide(a, b int) {
	result, err := divide(a, b)
	if err != nil {
		fmt.Printf("Error dividing %d by %d: %v\n", a, b, err)
	} else {
		fmt.Printf("Result of dividing %d by %d: %d\n", a, b, result)
	}
}
