package main

import (
	"fmt"
	"sync"
)

type singleton struct {
	value map[string]string
}
var instance *singleton
var once sync.Once

func GetSingleton() *singleton {
		once.Do(func() {
		instance = &singleton{value:make(map[string]string)}
	})
	return instance
}
func main() {
	s1 := GetSingleton()
	s1.value["foo"] = "bar"
	fmt.Println("This is ", s1.value["foo"])

	s2 := GetSingleton()
	fmt.Println("This is ", s2.value["foo"])
	s2.value["foo"] = "Fun"

	fmt.Println("This is ", s1.value["foo"])
}
