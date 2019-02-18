package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type IButton interface {
	OnClick()
}

type ICheckBox interface {
	OnCheck()
}

type IGuiFactory interface {
	CreateButton() IButton
	CreateCheckBox() ICheckBox
}

func main() {
	fmt.Println("Abstract Factory Demo")
	reader := bufio.NewReader(os.Stdin)
	var GUIFactory IGuiFactory
	for {
		fmt.Println("type [w] to create Win GUI, [m] to create Mac GUI")
		text, _ := reader.ReadString('\n')

		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		if strings.Compare("w", text) == 0 {
			GUIFactory = new(WinGUIFactory)
		} else if strings.Compare("m", text) == 0 {
			GUIFactory = new(MacGUIFactory)
		}
		var aButton IButton
		var aCheckBox ICheckBox
		aButton = GUIFactory.CreateButton()
		aCheckBox = GUIFactory.CreateCheckBox()
		aButton.OnClick()
		aCheckBox.OnCheck()

	}

}
