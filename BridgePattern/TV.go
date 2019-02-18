package main

import (
	"fmt"
)

type TV struct {

}

func (tV *TV) TurnOn() {
	fmt.Println("Turn on TV")
}

func (tV *TV) SwitchOff() {
	fmt.Println("Switch off TV")
}