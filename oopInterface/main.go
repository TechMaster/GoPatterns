package main

import (
	"fmt"
)

type iSwim interface {
	Swim()
	Dive()
}

type iFly interface {
	Fly()
}
type Duck struct {

}

type WildDuck struct {

}
//---- Vịt Nhà không biết bay chỉ bơi và lặn
func (duck *Duck) Swim() {
	fmt.Println("Duck can Swim")
}
/*
func (duck *Duck) Dive() {
	fmt.Println("Duck can Dive")
}*/
//----- Vịt Trời

func (duck *WildDuck) Swim() {
	fmt.Println("Wild Duck can Swim")
}

func (duck *WildDuck) Dive() {
	fmt.Println("Wild Duck can Dive")
}

func (duck *WildDuck) Fly() {
	fmt.Println("Wild Duck can Fly")
}


func main() {
	wildDuck := WildDuck{}
	wildDuck.Fly()
	wildDuck.Swim()


	//Chú ý cách sử dụng biến kiểu iSwim
	var aDuck iSwim = &wildDuck
	aDuck.Swim()


	domesticDuck := Duck{}
	aDuck = &domesticDuck
	aDuck.Dive()

}

