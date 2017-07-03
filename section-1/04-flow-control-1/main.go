package main

import "fmt"

func main() {
	fmt.Print("Input a fruit: ")
	var fruit string
	fmt.Scanln(&fruit)

	if fruit == "" {
		fmt.Println("meh! 😑")
		return
	}

	switch fruit {
	case "apple":
		fmt.Println("🍎")
	case "banana":
		fmt.Println("🍌")
	case "lemon":
		fmt.Println("🍋")
	case "pineapple":
		fmt.Println("🍍")
	default:
		fmt.Println("💩")
	}
}
