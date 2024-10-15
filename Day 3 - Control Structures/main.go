package main

import "fmt"

func main() {
	age := 30
	day := "Monday"

	// If else statement
	if age > 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}

	// Switch statement
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("Weekend is near")
	default:
		fmt.Println("Another day")
	}

	// For loop
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	//Check if a number is Odd or even
	if age%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
