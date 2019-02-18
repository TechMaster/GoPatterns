package main

import (
	"fmt"
	"reflect"
)

type iSwim interface {
	Swim()
	Dive()
}

type iFly interface {
	Fly()
}


type WildDuck struct {
	name string
	age int

}

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

func (duck WildDuck) Say(text string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(text)
	}
}


func main() {
	wildDuck := WildDuck{}

	//Liệt kê các thuộc tính
	st := reflect.TypeOf(wildDuck)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i).Name
		fmt.Println(field)
	}

	//Liệt kê các method chú ý lần này phải dùng reflect.TypeOf(&wildDuck)
	st2 := reflect.TypeOf(&wildDuck)

	for j := 0; j < st2.NumMethod(); j++ {
		method := st2.Method(j).Name
		fmt.Println(method)

	}


	//Gọi method bằng tên. Phải dùng ValueOf chứ không phải TypeOf
	reflect.ValueOf(&wildDuck).MethodByName("Fly").Call([]reflect.Value{})


	Invoke(&wildDuck, "Say", "Quack Quack", 5)



	var somethingCanSwim iSwim =  &wildDuck
	Invoke(somethingCanSwim, "Swim")
	Invoke(somethingCanSwim, "Fly") // Vẫn chạy tốt mặc dù somethingCanSwim kiểu iSwim

}

func Invoke(any interface{}, name string, args... interface{}) {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	reflect.ValueOf(any).MethodByName(name).Call(inputs)
}