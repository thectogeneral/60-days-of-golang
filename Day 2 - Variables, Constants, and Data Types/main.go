package main

import "fmt"

func main() {
	var name string = "Go Developer"
	fmt.Println(name)

	//Go also supports shorthand declaration:
	age := 30
	fmt.Println(age)

	//Declare constants:
	const pi = 3.14
	fmt.Println(pi)

	var score int = 95
	fmt.Println(score)

	var average float64 = 88.9
	fmt.Println(average)

	var status bool = true
	fmt.Println(status)

	var x = 42
	fmt.Println(x)

	var a int = 5
	var b float64 = float64(a)
	fmt.Println(b)

}
