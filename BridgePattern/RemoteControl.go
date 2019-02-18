package main

type RemoteControl struct {
	device IDevice
}

func NewRemoteControl (device IDevice) *RemoteControl {
	rc := new(RemoteControl)
	rc.device = device
	return rc
}
func (remoteControl *RemoteControl) TurnOn() {
	remoteControl.device.TurnOn()
}

func (remoteControl *RemoteControl) SwitchOff() {
	remoteControl.device.SwitchOff()
}