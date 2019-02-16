package main

type IWorker interface {
	getWorker() string
}

type Carpenter struct {

}

func (Carpenter* Carpenter) getWorker() string {
	return "Carpenter"
}

type SteelWelder struct {

}

func (SteelWelder* SteelWelder) getWorker() string {
	return "Steel Welder"
}