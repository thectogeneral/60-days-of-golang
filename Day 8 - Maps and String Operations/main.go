package main

import "fmt"

func updateGrade(students map[string]int, name string, newGrade int) {
	students[name] = newGrade
}

func main() {
	students := map[string]int{
		"Alice":   85,
		"Bob":     90,
		"Charlie": 78,
	}

	fmt.Println("Original Grades:", students)
	updateGrade(students, "Alice", 92)
	fmt.Println("Updated Grades:", students)
}
