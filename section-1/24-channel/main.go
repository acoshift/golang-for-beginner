package main

import (
	"fmt"
	"time"
)

var (
	arr1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 = []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
)

func main() {
	withoutGo()
	withGo()
}

func timer() func() {
	t := time.Now()
	return func() {
		diff := time.Now().Sub(t)
		fmt.Println(diff)
	}
}

func withoutGo() {
	defer timer()()
	fmt.Println("sum of arr1 is", sum(arr1))
	fmt.Println("sum of arr2 is", sum(arr2))
}

func withGo() {
	defer timer()()
	chRes1 := make(chan int)
	chRes2 := make(chan int)
	go func() {
		chRes1 <- sum(arr1)
	}()
	go func() {
		chRes2 <- sum(arr2)
	}()

	fmt.Println("sum of arr1 is", <-chRes1)
	fmt.Println("sum of arr2 is", <-chRes2)
}

func sum(arr []int) int {
	s := 0
	for _, x := range arr {
		s += x
		time.Sleep(time.Millisecond * 200)
	}
	return s
}
