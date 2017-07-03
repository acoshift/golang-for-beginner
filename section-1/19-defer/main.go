package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		diff := time.Now().Sub(start)
		fmt.Printf("time: %dms\n", diff/time.Millisecond)
	}()
	time.Sleep(time.Second)
}
