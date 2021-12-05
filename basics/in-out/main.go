package main

import "fmt"

func main() {
	var name string
	var age byte
	var salary float32
	var isPassExam bool

	/*
		fmt.Println("fmt.Scanln()")
		fmt.Println("Please Enter Your Name:")
		fmt.Scanln(&name)
		fmt.Println("Please Enter Your Age:")
		fmt.Scanln(&age)
		fmt.Println("Please Enter Your Salary:")
		fmt.Scanln(&salary)
		fmt.Println("Did you pass the exam:")
		fmt.Scanln(&isPassExam)

		fmt.Printf(" name: %v\n age: %v\n salary: %v\n isPassExam: %v\n", name, age, salary, isPassExam)
	*/

	fmt.Println("fmt.Scanf()")
	fmt.Scanf("%s %d %f %t", &name, &age, &salary, &isPassExam)
	fmt.Printf(" name: %v\n age: %v\n salary: %v\n isPassExam: %v\n", name, age, salary, isPassExam)
}
