package main

import "fmt"

func main() {
	var a [5]int
	a[0] = 10
	a[2] = 30
	a[3] = 40
	fmt.Println(a)
	fmt.Println(len(a))

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	for i := range a {
		fmt.Println(a[i])
	}
	for _, v := range a {
		fmt.Println(v)
	}

	b := [3]int{1, 2, 3}
	fmt.Println(b)

	var c [2][3]int
	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[i]); j++ {
			c[i][j] = j
		}
	}
	fmt.Println(c)
}
