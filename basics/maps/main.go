package main

import "fmt"

func main() {
	employeeSalary := make(map[string]int)
	employeeSalary["Signal"] = 10000
	employeeSalary["Fish"] = 9000
	employeeSalary["Marsha"] = 8000
	emp := "Signal"
	v, ok := employeeSalary[emp]
	if ok {
		fmt.Printf("Salary of %v is %v\n", emp, v)
	} else {
		fmt.Print("%v not found")
	}
	for k, v := range employeeSalary {
		fmt.Printf("employeeSalary[%s] = %d\n", k, v)
	}

	fmt.Println("Before deleting, employeeSalary =", employeeSalary)
	delete(employeeSalary, "Fish")
	fmt.Println("After deleting, employeeSalary =", employeeSalary)
}
