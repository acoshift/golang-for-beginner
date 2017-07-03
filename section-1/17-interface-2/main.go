package main

import "fmt"

func main() {
	checkType("Gopher")
	checkType(10)
	checkType(true)
}

func checkType(v interface{}) {
	switch p := v.(type) {
	case string:
		fmt.Println("String:", "hello "+p)
	case int:
		fmt.Println("Int:", p+10)
	case bool:
		fmt.Println("Bool:", !p)
	}
}
