package main

import "fmt"

func updateValue(x *int) {
	*x = *x + 10
}

func main() {
	num := 50
	fmt.Println("Before:", num)
	updateValue(&num)
	fmt.Println("After:", num)
}
