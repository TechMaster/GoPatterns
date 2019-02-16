package main

import (
	"fmt"
)

type Weekday int

const (
	Sunday    Weekday = iota + 1 // value: 1, type: Weekday
	Monday                       // value: 2, type: Weekday
	Tuesday                      // value: 3, type: Weekday
	Wednesday                    // value: 4, type: Weekday
	Thursday                     // value: 5, type: Weekday
	Friday                       // value: 6, type: Weekday
	Saturday                     // value: 7, type: Weekday
)

func main() {
	fmt.Println(Sunday)
	fmt.Printf(
		"Sunday=%[1]d (Type=%[1]T)\n"+
			"Monday=%[2]d (Type=%[2]T)\n"+
			"Tuesday=%[3]d (Type=%[3]T)\n"+
			"Wednesday=%[4]d (Type=%[4]T)\n"+
			"Thursday=%[5]d (Type=%[5]T)\n"+
			"Friday=%[6]d (Type=%[6]T)\n"+
			"Saturday=%[7]d (Type=%[7]T)\n",
		Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday,
	)
}
