package main

import "fmt"

type Car struct {
	Model string
}

func main() {
	//c := Car{
	//	Model: "BMW",
	//}

	c := new(Car) // pointer is created using new
	c.Model = "Ferrari"

	fmt.Println(c)

	fmt.Println("=========")

	showCar(c)

	fmt.Println("=========")

	fmt.Println(c)

}

func showCar(c *Car) {
	c.Model = "Mercedes"
	fmt.Println(c)
}
