/*
https://blog.learngoprogramming.com/golang-const-type-enums-iota-bc4befd096d3
 */
package basic

import (
	"fmt"
)
type Weekday int
const (
	Sunday    Weekday = 0
	Monday    Weekday = 1
	Tuesday   Weekday = 2
	Wednesday Weekday = 3
	Thursday  Weekday = 4
	Friday    Weekday = 5
	Saturday  Weekday = 6
)
func declareEnumAsConst() {
	// will display 0
	fmt.Println(Sunday)
	// will display 6
	fmt.Println(Saturday)
}

//Return name of day in week
func (day Weekday) String() string {
	// declare an array of strings
	// ... operator counts how many
	// items in the array (7)
	names := [...]string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday"}
	// → `day`: It's one of the
	// values of Weekday constants.
	// If the constant is Sunday,
	// then day is 0.
	// prevent panicking in case of
	// `day` is out of range of Weekday
	if day < Sunday || day > Saturday {
		return "Unknown"
	}
	// return the name of a Weekday
	// constant from the names array
	// above.
	return names[day]
}
func main() {
	declareEnumAsConst()

	fmt.Printf("Which day it is? %s\n", Sunday)
	// Which day it is? Sunday
}
