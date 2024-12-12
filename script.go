package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Create an instance of Person
	p := Person{Name: "Alice", Age: 25}

	// Marshal the struct into JSON
	data, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("JSON:", string(data))

	// Unmarshal the JSON back into a struct
	var p2 Person
	err = json.Unmarshal(data, &p2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Struct: %+v\n", p2)
}
