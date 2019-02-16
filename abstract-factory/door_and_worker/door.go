package main

type IDoor interface {
	getType() string
	getMaterial() string
}
type GenericDoor struct {
	material string
}
//-------Wooden Door---------
type WoodenDoor struct {
	GenericDoor
}
func NewWoodenDoor(material string) *WoodenDoor {
	var woodenDoor = WoodenDoor{}
	woodenDoor.material = material
	return &woodenDoor
}

func (WoodenDoor * WoodenDoor) getType() string {
	return "Wooden Door"
}

func (WoodenDoor * WoodenDoor) getMaterial() string {
	return WoodenDoor.material
}

//-------Steel Door---------
type SteelDoor struct {
	GenericDoor
}

func NewSteelDoor(material string) *SteelDoor {
	var steelDoor = SteelDoor{}
	steelDoor .material = material
	return &steelDoor
}

func (SteelDoor *SteelDoor) getType() string {
	return "Steel Door"
}

func (SteelDoor *SteelDoor) getMaterial() string {
	return SteelDoor.material
}