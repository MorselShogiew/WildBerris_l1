package main

import "fmt"

func (hum Human) Name() string {
	return hum.name
}

type Human struct {
	name    string
	surname string
}

type Action struct {
	Human
	num string
}

func main() {

	pers := Human{"Marsel", "Shagiew"}
	person := Action{pers, "123"}
	fmt.Println("Name of person is", person.Name())
}
