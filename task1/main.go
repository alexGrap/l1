package main

import "fmt"

// Human struct is parent structure with method "Presentation" and fields "name" and "surname"
type Human struct {
	name    string
	surname string
}

func (human *Human) Presentation() {
	fmt.Printf("Bonjour, my name %s %s", human.name, human.surname)
}

// Action struct is child structure that "inherits" the fields and methods of the parent structure
type Action struct {
	Human
}

// main has initialization of Action structure and calls their methods
func main() {
	init := Action{Human{name: "Aleksey", surname: "Pargasoov"}}
	init.Presentation()
}
