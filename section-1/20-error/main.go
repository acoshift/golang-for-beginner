package main

import "fmt"

func validateAge1(age int) bool {
	if age < 18 {
		return false
	} else if age > 60 {
		return false
	}
	return true
}

func validateAge2(age int) error {
	if age < 18 {
		return fmt.Errorf("age too low")
	} else if age > 60 {
		return fmt.Errorf("age too high")
	}
	return nil
}

type user struct {
	ID   int
	Name string
}

func getUser(id int) (*user, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid id")
	}
	return &user{ID: id, Name: "Test"}, nil
}

func main() {
	fmt.Println(validateAge1(10))
	fmt.Println(validateAge2(10))

	u, err := getUser(0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u)
	}

	u, err = getUser(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u)
	}
}
