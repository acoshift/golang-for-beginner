package main

import "fmt"

func main() {
	a := make([]int, 5)
	a[0] = 10
	a[2] = 30
	a[3] = 40
	fmt.Println(a)
	fmt.Println(len(a))

	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(b)
	fmt.Println(b[2:4])
	fmt.Println(b[:4])
	fmt.Println(b[4:])

	c := [5]int{1, 2, 3, 4, 5}
	fmt.Println(c)
	d := c[:2]
	d[0] = 20
	fmt.Println(c)
	fmt.Println(d)

	d = append(d, 30)
	fmt.Println(c)
	fmt.Println(d)
}
