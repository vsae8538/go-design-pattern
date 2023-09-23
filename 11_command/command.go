package command

import "fmt"

// Command defines the interface for commands that can be executed.
type Command interface {
	Execute()
}

// StartCommand implements the Command interface and contains a receiver.
type StartCommand struct {
	mb *MotherBoard
}

// NewStartCommand creates a new StartCommand.
func NewStartCommand(mb *MotherBoard) *StartCommand {
	return &StartCommand{mb: mb}
}

// Execute implements the Command interface.
func (c *StartCommand) Execute() {
	c.mb.Start()
}

// RebootCommand implements the Command interface and contains a receiver.
type RebootCommand struct {
	mb *MotherBoard
}

// NewRebootCommand creates a new RebootCommand.
func NewRebootCommand(mb *MotherBoard) *RebootCommand {
	return &RebootCommand{
		mb: mb,
	}
}

// Execute implements the Command interface.
func (c *RebootCommand) Execute() {
	c.mb.Reboot()
}

// MotherBoard implements the receiver that the commands act on.
type MotherBoard struct{}

// Start starts the motherboard.
func (*MotherBoard) Start() {
	fmt.Print("system starting\n")
}

// Reboot reboots the motherboard.
func (*MotherBoard) Reboot() {
	fmt.Print("system rebooting\n")
}

// Box contains two commands.
type Box struct {
	button1 Command
	button2 Command
}

// NewBox creates a new Box.
func NewBox(button1, button2 Command) *Box {
	return &Box{
		button1: button1,
		button2: button2,
	}
}

// PressButton1 executes the first command.
func (b *Box) PressButton1() {
	b.button1.Execute()
}

// PressButton2 executes the second command.
func (b *Box) PressButton2() {
	b.button2.Execute()
}
