package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Abstract Factory")
	fmt.Println("-----------------")
	var DoorFactory IFactory
	var aDoor IDoor
	var aWorker IWorker

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("type [w] to create WoodenDoorFactory, [s] to create SteelDoorFactory")
		text, _ := reader.ReadString('\n')

		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		if strings.Compare("w", text) == 0 {
			DoorFactory = NewWoodenDoorFactory()
		} else if strings.Compare("s", text) == 0 {
			DoorFactory = NewSteelDoorFactory()
		}

		aDoor = DoorFactory.GetDoor()
		aWorker = DoorFactory.GetWorker()
		fmt.Println(aDoor.getMaterial())
		fmt.Println(aDoor.getType())
		fmt.Println(aWorker.getWorker())
	}
}



