package main

//func main() {
//	/**
//	 * The remote is our Invoker;
//	 * it will be passed a
//	 * command object that can
//	 * be used to make requests.
//	 */
//	RemoteControl := &SimpleRemoteControl{}
//
//	/**
//	 * Now we create a Light
//	 * object, this will be the
//	 * Receiver of the request
//	 */
//	light := &light{}
//
//	/**
//	 * Create a command and
//	 * pass the Receiver to it.
//	 */
//	lightOnCommand := NewLightOnCommand(light)
//
//	/**
//	 * Pass the command to the Invoker
//	 */
//	RemoteControl.SetCommand(lightOnCommand)
//
//	/**
//	 * We simulate the button being pressed.
//	 */
//	RemoteControl.ButtonWasPressed()
//
//	garage := &garage{}
//
//	GarageDoorOpenCommand := newGarageDoorOpenCommand(garage)
//
//	/**
//	 * Pass the new command to the invoker
//	 */
//	RemoteControl.SetCommand(GarageDoorOpenCommand)
//
//	RemoteControl.ButtonWasPressed()
//}
