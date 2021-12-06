package main

import (
	"encoding/json"
	"fmt"
)

type monster struct {
	Name     string
	Age      int
	Birthday string
	Salary   float64
	Skill    string
}

func testStruct() {
	monster := monster{
		Name:     "Signal",
		Age:      24,
		Birthday: "2000-01-01",
		Salary:   10000,
		Skill:    "React",
	}

	fmt.Println("Before serialization:", monster)
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("Marshal Error occurred:", err)
	}
	fmt.Println("After Serialization, monster =", string(data))
}

func testMap() {
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["name"] = "Signal"
	m["age"] = 24

	fmt.Println("Before serialization, m =", m)
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Marshal Error Occurred:", err)
	}
	fmt.Println("After Serialization, m =", string(data))
}

func testSlice() {
	var slice []map[string]interface{}

	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "Signal"
	m1["age"] = 24
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "Marsha"
	m2["age"] = 23
	slice = append(slice, m2)

	fmt.Println("Before serialization, slice =", slice)
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("Marshal Error Occurred:", err)
	}
	fmt.Println("After serialization, slice =", data)
}

func testFloat64() {
	var num1 float64 = 1.23

	fmt.Println("Before serialization, num1 =", num1)
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Println("Marshal Error Occurred:", err)
	}
	fmt.Println("After serialization, num1 =", data)
}

func main() {
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}
