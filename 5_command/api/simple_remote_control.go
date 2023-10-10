package api

type SimpleRemoteControl struct {
	slot ICommand
}

/**
 * We have a method for setting
 * the command the slot is going
 * to control. This could be called
 * multiple times if the client of
 * this code wanted to change the
 * behavior of the remote button.
 */
func (r *SimpleRemoteControl) SetCommand(command ICommand) {
	r.slot = command
}

/**
 * This method is called when the
 * button is pressed. All we do is take
 * the current command bound to the
 * slot and call its Execute() method.
 */
func (r *SimpleRemoteControl) ButtonWasPressed() {
	r.slot.Execute()
}
