package main


import (
	"fmt"
)
type WinGUIFactory struct {

}

func (WinGUI *WinGUIFactory) CreateButton() IButton {
	return new(WinButton)
}

func (WinGUI *WinGUIFactory) CreateCheckBox() ICheckBox {
	return new(WinCheckBox)
}

type WinButton struct {

}

func (winButton *WinButton) OnClick(){
	fmt.Println("Win Button OnClick")
}

type WinCheckBox struct {

}
func (WinCheckBox *WinCheckBox) OnCheck() {
	fmt.Println("Win CheckBox OnCheck")
}