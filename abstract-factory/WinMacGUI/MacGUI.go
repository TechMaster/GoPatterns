package main

import (
	"fmt"
)
type MacGUIFactory struct {

}


func (macGUI *MacGUIFactory) CreateButton() IButton {
	return new(MacButton)
}

func (macGUI *MacGUIFactory) CreateCheckBox() ICheckBox {
	return new(MacCheckBox)
}

type MacButton struct {

}

func (macButton *MacButton) OnClick(){
	fmt.Println("Mac Button OnClick")
}

type MacCheckBox struct {

}
func (macCheckBox *MacCheckBox) OnCheck() {
	fmt.Println("Mac CheckBox OnCheck")
}