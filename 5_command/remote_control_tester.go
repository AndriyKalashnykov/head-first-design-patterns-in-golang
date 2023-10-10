package main

import (
	"fmt"

	"head-first-design-patterns-in-golang/5_command/api"
)

func main() {
	remoteControl := api.NewRemoteControl()

	// Instantiate all the devices
	livingRoomLight := &api.Light{"Living Room"}
	kitchenLight := &api.Light{"Kitchen"}
	ceilingFan := &api.CeilingFan{"Living Room"}
	garage := &api.Garage{}
	stereo := &api.Stereo{"Living Room"}

	// Instantiate all the Command Objects
	livingRoomLightOn := api.NewLightOnCommand(livingRoomLight)
	livingRoomLightOff := api.NewLightOffCommand(livingRoomLight)
	kitchenLightOn := api.NewLightOnCommand(kitchenLight)
	kitchenLightOff := api.NewLightOffCommand(kitchenLight)

	ceilingFanOn := api.NewCeilingFanOnCommand(ceilingFan)
	ceilingFanOff := api.NewCeilingFanOffCommand(ceilingFan)

	stereoOnWithCD := api.NewStereoOnCommand(stereo)
	stereoOff := api.NewStereoOffCommand(stereo)

	garageDoorOpen := api.NewGarageDoorOpenCommand(garage)
	garageDoorClose := api.NewGarageDoorCloseCommand(garage)

	livingRoomOnCommands := []api.ICommand{livingRoomLightOn, ceilingFanOn, stereoOnWithCD}
	livingRoomOffCommands := []api.ICommand{livingRoomLightOff, ceilingFanOff, stereoOff}
	livingRoomMacroOnCommand := api.NewMacroCommand(livingRoomOnCommands)
	livingRoomMacroOffCommand := api.NewMacroCommand(livingRoomOffCommands)

	// Load commands into the remote slots.
	remoteControl.SetCommand(0, livingRoomLightOn, livingRoomLightOff)
	remoteControl.SetCommand(1, kitchenLightOn, kitchenLightOff)
	remoteControl.SetCommand(2, ceilingFanOn, ceilingFanOff)
	remoteControl.SetCommand(3, stereoOnWithCD, stereoOff)
	remoteControl.SetCommand(4, garageDoorOpen, garageDoorClose)
	remoteControl.SetCommand(5, livingRoomMacroOnCommand, livingRoomMacroOffCommand)

	fmt.Println(remoteControl)

	//  We step through each slot and push its On and Off button.
	remoteControl.OnButtonWasPushed(0)
	remoteControl.OffButtonWasPushed(0)
	remoteControl.OnButtonWasPushed(1)
	remoteControl.OffButtonWasPushed(1)
	remoteControl.OnButtonWasPushed(2)
	remoteControl.OffButtonWasPushed(2)
	remoteControl.OnButtonWasPushed(3)
	remoteControl.OffButtonWasPushed(3)
	remoteControl.OnButtonWasPushed(4)
	remoteControl.OffButtonWasPushed(4)

	fmt.Println("---Pushing Macro On---")
	remoteControl.OnButtonWasPushed(5)

	fmt.Println("---Pushing Macro Off---")
	remoteControl.OffButtonWasPushed(5)

}
