package main

func main() {

	tv := new(TV)
	radio := new(Radio)

	tvRemote := NewRemoteControl(tv)
	radioRemote := NewRemoteControl(radio)

	tvRemote.TurnOn()
	radioRemote.SwitchOff()
}
