package main

import (
	"fmt"
	shape "github.com/TechMaster/GoPatterns/oop2/Rect"
)

func main() {
	//r := shape.Rect{width: 10, height: 5}
	r := shape.NewRect(10,5)
	fmt.Println("area: ", r.Area())
	fmt.Println("perimeter: ", r.Perimeter())

	r.Scale(2)
	fmt.Printf("%p\n", r)
	fmt.Println(r.Info())
}
