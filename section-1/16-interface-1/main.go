package main

import "fmt"

type person struct {
	Name string
}

func (p person) Talk() {
	fmt.Println("Hello, I'm", p.Name)
}

type cat struct{}

func (cat) Talk() {
	fmt.Println("Nyaa nyaaa")
}

type dog struct{}

func (dog) Talk() {
	fmt.Println("Wan wan")
}

type talkable interface {
	Talk()
}

func talkWith(t talkable) {
	t.Talk()
}

func main() {
	p := person{"Gopher"}
	talkWith(p)

	talkWith(cat{})
	talkWith(&dog{})
}
