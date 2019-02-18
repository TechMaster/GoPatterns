package main

type IFactory interface {
	GetDoor() IDoor
	GetWorker() IWorker
}
//--------- type WoodenDoorFactory
type WoodenDoorFactory struct {
	door IDoor
	worker IWorker
}

//Constructor
func NewWoodenDoorFactory() IFactory {
	var woodenDoorFactory = new(WoodenDoorFactory)
	woodenDoorFactory.door =  NewWoodenDoor("African Oak")
	woodenDoorFactory.worker = new(Carpenter)
	return woodenDoorFactory
}

func (WoodenDoorFactory WoodenDoorFactory) GetDoor() IDoor {
	return WoodenDoorFactory.door
}

func (WoodenDoorFactory WoodenDoorFactory) GetWorker() IWorker {
	return WoodenDoorFactory.worker
}

//--------- type SteelDoorFactory
type SteelDoorFactory struct {
	door IDoor
	worker IWorker
}

//Constructor
func NewSteelDoorFactory() IFactory {
	var steelDoorFactory = new(SteelDoorFactory)
	steelDoorFactory.door = NewSteelDoor("US Carnege Steel")
	steelDoorFactory.worker = new(SteelWelder)
	return steelDoorFactory
}

func (SteelDoorFactory SteelDoorFactory) GetDoor() IDoor {
	return SteelDoorFactory.door
}

func (SteelDoorFactory SteelDoorFactory) GetWorker() IWorker {
	return SteelDoorFactory.worker
}