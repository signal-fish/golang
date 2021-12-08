package main

import "fmt"

/*
1. In Go, an interface is a set of method signatures.
2. An interface that has zero methods is called an empty interface. It is represented as interface{}.
*/
func main() {
	pemp1 := Permanent{
		empID:    1,
		basicpay: 5000,
		pf:       20,
	}
	pemp2 := Permanent{
		empID:    2,
		basicpay: 6000,
		pf:       30,
	}
	cemp1 := Contract{
		empID:    3,
		basicpay: 3000,
	}
	freelancer1 := Freelancer{
		empID:       4,
		ratePerHour: 70,
		totalHours:  120,
	}
	freelancer2 := Freelancer{
		empID:       5,
		ratePerHour: 100,
		totalHours:  100,
	}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1, freelancer1, freelancer2}
	totalExpense(employees)
}

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empID    int
	basicpay int
	pf       int
}

type Contract struct {
	empID    int
	basicpay int
}

type Freelancer struct {
	empID       int
	ratePerHour int
	totalHours  int
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

func (f Freelancer) CalculateSalary() int {
	return f.ratePerHour * f.totalHours
}

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}

	fmt.Printf("Total Expense Per Month $%d", expense)
}
