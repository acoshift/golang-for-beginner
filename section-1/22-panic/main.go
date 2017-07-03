package main

import "fmt"

func main() {
	fmt.Println("start")
	// doFailWork()
	doSafeWork()
	fmt.Println("done")
}

func doFailWork() {
	panic("fail")
}

func doSafeWork() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("work fail:", r)
		}
	}()
	doFailWork()
	fmt.Println("work success")
}
