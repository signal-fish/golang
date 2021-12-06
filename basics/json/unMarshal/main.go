package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name   string
	Age    int
	Salary float64
}

func unmarshalStruct() {
	str := `{"Name": "Signal", "Age": 24, "Salary": 10000}`

	fmt.Println("Before unserialization, str =", str)
	var signal user
	err := json.Unmarshal([]byte(str), &signal)
	if err != nil {
		fmt.Println("Unmarshal Error Occurred:", err)
	}
	fmt.Println("After unserialization, signal =", signal)
}

func unmarshalMap() {
	str := `{"name": "Signal", "age": 24, "salary": 10000}`

	fmt.Println("Before unserialization, str =", str)
	var m map[string]interface{}
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Println("Unmarshal Error Occurred:", err)
	}
	fmt.Println("After UnSerialization, m =", m)
}

func unmarshalSlice() {
	str := `[{"name": "Signal", "age": 24, "salary": 10000}, {"name": "Marshal", "age": 23, "salary": 20000}]`
	fmt.Println("Before unserialization, str =", str)

	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println("Unmarshal Error Occurred:", err)
	}
	fmt.Println("After Serialization, slice =", slice)
}

func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}
