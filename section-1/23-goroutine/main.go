package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")
	go talk("hello")
	go talk("hi")
	fmt.Println("Waiting")
	time.Sleep(5 * time.Second)
	fmt.Println("End")
}

func talk(prefix string) {
	for i := 0; i < 10; i++ {
		fmt.Println(prefix, i)
		time.Sleep(time.Second)
	}
}
