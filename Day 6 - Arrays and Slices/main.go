package main

import "fmt"

func main() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	slice = append(slice, 6)
	fmt.Println(slice)

}
