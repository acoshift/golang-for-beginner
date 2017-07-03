package main

import "fmt"

type person struct {
	Name     string
	NickName string
}

func (p person) Mutate1() {
	p.Name = "Hacker1"
	fmt.Println("mutate1:", p)
}

func (p *person) Mutate2() {
	p.Name = "Hacker2"
	fmt.Println("mutate2:", p)
}

func mutatePerson1(p person) {
	p.Name = "Hacker1"
	fmt.Println("mutatePerson1:", p)
}

func mutatePerson2(p *person) {
	p.Name = "Hacker2"
	fmt.Println("mutatePerson2:", p)
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

	p1.Mutate1()
	fmt.Println(p1)

	p1.Mutate2()
	fmt.Println(p1)

	mutatePerson1(p2)
	fmt.Println(p2)
	mutatePerson2(&p2)
	fmt.Println(p2)
}
