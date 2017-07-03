package main

import "fmt"

func main() {
	m := make(map[interface{}]interface{})
	m[1] = "Hello"
	m["name"] = "Gopher"

	if x, ok := m["name"].(string); ok {
		fmt.Println("name:", x)
	}

	if x, ok := m["name"].(int); ok {
		fmt.Println("name:", x)
	}
}
