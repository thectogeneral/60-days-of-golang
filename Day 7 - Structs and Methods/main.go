package main

import "fmt"

func main() {
	myCar := Car{
		Model: "Toyota Camry",
		Year:  2021,
	}

	// Call the Start method
	myCar.Start()

	// Call the Stop method
	myCar.Stop()

	// Update the Year using the UpdateYear method
	fmt.Println("Before update, Year:", myCar.Year)
	myCar.UpdateYear(2023)
}

type Car struct {
	Model string
	Year  int
}

func (c Car) Start() {
	fmt.Println(c.Model, "is starting")
}

func (c Car) Stop() {
	fmt.Println(c.Model, "is stoping")
}

func (c *Car) UpdateYear(year int) {
	c.Year = year
}
