package api

import (
	"fmt"
	"reflect"
)

/**
 * the remote is going to
 * handle seven On and Off commands, which
 * we’ll hold in corresponding arrays.
 */
type RemoteControl struct {
	OnCommands, OffCommands []ICommand
}

func NewRemoteControl() *RemoteControl {
	return &RemoteControl{
		OnCommands:  make([]ICommand, 7),
		OffCommands: make([]ICommand, 7),
	}
}

/**
 * The SetCommand() method takes a slot position
 * and an On and Off command to be stored in
 * that slot. It puts these commands in the on and
 * off arrays for later use
 */
func (r *RemoteControl) SetCommand(slot int, onCommand, offCommand ICommand) {
	r.OnCommands[slot] = onCommand
	r.OffCommands[slot] = offCommand
}

/**
 * When an On or Off button is
 * pressed, the hardware takes
 * care of calling the corresponding
 * methods OnButtonWasPushed() or
 * OffButtonWasPushed().
 */
func (r *RemoteControl) OnButtonWasPushed(slot int) {
	fmt.Println("*****")
	r.OnCommands[slot].Execute()
	fmt.Println("*****")
}

func (r *RemoteControl) OffButtonWasPushed(slot int) {
	fmt.Println("*****")
	r.OffCommands[slot].Execute()
	fmt.Println("*****")
}

/**
 * Implementing String() to print out each slot and its
 * corresponding command.
 */
func (r *RemoteControl) String() string {
	s := fmt.Sprintf("\n------ Remote Control -------\n")

	for i := range r.OnCommands {
		if r.OnCommands[i] == nil {
			continue
		}

		onClass := r.GetClassName(r.OnCommands[i])
		offClass := r.GetClassName(r.OffCommands[i])
		s += fmt.Sprintf("[slot %d] %s   %s\n", i, onClass, offClass)
	}
	s += fmt.Sprintf("-----------------------------\n")

	return s
}

func (r *RemoteControl) GetClassName(myVar interface{}) string {
	if t := reflect.TypeOf(myVar); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}
