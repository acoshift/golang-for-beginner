package main

import (
	"fmt"
	"time"
)

func main() {
	res, err := doWorkWithLimitTime(1 * time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func doVeryLongWork() int {
	time.Sleep(4 * time.Second)
	return 1
}

func doWorkWithLimitTime(d time.Duration) (int, error) {
	ch := make(chan int)
	go func() {
		ch <- doVeryLongWork()
	}()
	select {
	case r := <-ch:
		return r, nil
	case <-time.After(d):
		return 0, fmt.Errorf("timeout")
	}
}
