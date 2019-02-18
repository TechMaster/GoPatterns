package main

import (
	"fmt"
)

type Radio struct {

}

func (radio *Radio) TurnOn() {
	fmt.Println("Turn on Radio")
}

func (radio *Radio) SwitchOff() {
	fmt.Println("Switch off Radio")
}