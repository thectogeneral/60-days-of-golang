package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Student struct {
	Person
	School string
}

type Employee struct {
	Person
	Company string
}

func main() {
	student := Student{Person: Person{Name: "Alice", Age: 20}, School: "Go University"}
	employee := Employee{Person: Person{Name: "Bob", Age: 30}, Company: "Go Corp"}

	fmt.Println("Student:", student)
	fmt.Println("Employee:", employee)
}
