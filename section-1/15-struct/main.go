package main

import "fmt"

type person struct {
	Name     string
	NickName string
}

func main() {
	p1 := person{
		Name:     "TT",
		NickName: "Acoshift",
	}
	fmt.Println(p1)

	p2 := person{"TT", "Acoshift"}
	fmt.Println(p2)

	var p3 person
	p3.Name = "TT"
	p3.NickName = "Acoshift"
	fmt.Println(p3)
}
