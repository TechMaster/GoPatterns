package main

import "fmt"
type Interviewer interface {
	AskQuestion()
}

type Developer struct  {
}
// Method value receiver
func (developer Developer) AskQuestion() {
	fmt.Println("Do you know Design Pattern?")
}
type HRManager struct {

}
// Method value receiver
func (hrManager HRManager) AskQuestion() {
	fmt.Println("Have you ever presented in WebSummit?")
}

func main() {
	fmt.Println("Factory Method - version 1")

	var interviewer  Interviewer
	interviewer = HRManager{}
	interviewer.AskQuestion()

	interviewer = Developer{}
	interviewer.AskQuestion()
}
