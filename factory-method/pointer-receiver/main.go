package main

import "fmt"
type Interviewer interface {
	AskQuestion()
}

type Developer struct  {
}
//Pointer receiver
func (developer *Developer) AskQuestion() {
	fmt.Println("Do you know Design Pattern?")
}
type HRManager struct {

}
//Pointer receiver
func (hrManager *HRManager) AskQuestion() {
	fmt.Println("Have you ever presented in WebSummit?")
}

func main() {
	fmt.Println("Factory Method")

	var interviewer  Interviewer
	interviewer = new(HRManager)
	interviewer.AskQuestion()

	interviewer = new(Developer)
	interviewer.AskQuestion()
}
