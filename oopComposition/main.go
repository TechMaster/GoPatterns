package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Address  //Compose Address inside
}
type Address struct {
	Number string
	Street string
	City   string
	State  string
	Zip    string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}
func (p *Person) Location() {
	fmt.Println("I’m at", p.Address.Number, p.Address.Street, p.Address.City, p.Address.State, p.Address.Zip)
}

func (a *Address) ShowStreet() {
	fmt.Println(a.Street)
}
func main() {
	p := Person{
		Name: "Steve",
		Address: Address{
			Number: "13",
			Street: "Main",
			City:   "Gotham",
			State:  "NY",
			Zip:    "01313",
		},
	}
	p.Talk()
	p.Location()
	p.ShowStreet()
}
