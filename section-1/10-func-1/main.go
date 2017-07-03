package main

import "fmt"

func main() {
	a, b := 1, 2
	r := add(a, b)
	fmt.Println(r)
}

func add(x, y int) int {
	p := 1
	return x + y + p
}
