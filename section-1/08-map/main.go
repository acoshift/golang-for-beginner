package main

import "fmt"

func main() {
	a := make(map[string]string)
	a["hello"] = "gopher"
	a["name"] = "acoshift"

	fmt.Println(a)
	fmt.Println(a["hello"])

	fmt.Println(a["x"])
	if x, ok := a["x"]; ok {
		fmt.Println(x)
	} else {
		fmt.Println("x not exists")
	}
	if name, ok := a["name"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("name not exists")
	}

	for key, value := range a {
		fmt.Println(key, ":", value)
	}

	delete(a, "hello")
	fmt.Println(a)

	b := map[string]string{
		"hello": "gopher",
		"name":  "acoshift",
	}
	fmt.Println(b)
}
