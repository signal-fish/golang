package main

import "fmt"

/*
1. A slice does not own any data of its own. It is just a representation of the underlying array. Any modifications done to the slice will be reflected in the underlying array.
2. The length of the slice is the number of elements in the slice. The capacity of the slice is the number of elements in the underlying array starting from the index from which the slice is created.
*/

func main() {
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars))

	fmt.Println("====================")

	var names []string
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "Signal", "Fish", "John")
		fmt.Println("names contents:", names)
	}
	fmt.Println("====================")
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)
}
