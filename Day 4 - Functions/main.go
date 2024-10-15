package main

import "fmt"

func main() {
	sum := add(2, 5)
	fmt.Println("The sum is:", sum)

	division, err := divide(float64(5), float64(2))
	if err != nil {
		fmt.Println("Error in division:", err)
	} else {
		fmt.Println("The division is:", division)
	}

	factorial := factorial(10)
	fmt.Println("The factorial is:", factorial)
}

func add(a int, b int) int {
	return a + b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
