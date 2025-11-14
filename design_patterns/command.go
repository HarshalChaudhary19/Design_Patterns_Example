package designpatterns

import "fmt"

//Turn line of codes as action into a small objects (a command)
//That can be stored,passed around,delayed,queued,undone etc

type Light struct{} //Ye struct se he On & Off methods ko

func (Light) On() {
	fmt.Println("Light is ON")
}

func (Light) Off() {
	fmt.Println("Light is OFF")
}

type Command interface {
	Execute()
}

type OnCommand struct {
	light Light
}

type OfCommand struct {
	light Light
}

func (c OnCommand) Execute() {
	c.light.On()
}

func (c OfCommand) Execute() {
	c.light.Off()
}

//Button that invokes the command
//It doesn't know what the command actually does

type Remote struct {
	button Command
}

func (r Remote) PressButton() {
	r.button.Execute()
}

func CommandExample() {
	light := Light{}
	oncommand := OnCommand{light}
	offcomamnd := OfCommand{light}

	//Plugging command into the remote
	remote := Remote{button: oncommand}
	remote.PressButton()

	remote = Remote{button: offcomamnd}
	remote.PressButton()
}
