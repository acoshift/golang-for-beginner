package main

import "fmt"

func main() {
	a := [5]int{}
	fmt.Println(a)
	mutateArray(a)
	fmt.Println(a)
	mutateArray2(&a)
	fmt.Println(a)
	mutateSlice(a[:])
	fmt.Println(a)
}

func mutateArray(a [5]int) {
	a[0] = 10
}

func mutateArray2(a *[5]int) {
	(*a)[0] = 10
}

func mutateSlice(a []int) {
	a[0] = 20
}
