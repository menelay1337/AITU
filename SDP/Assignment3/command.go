package main
import (
	"fmt"
	
)

// Command interface

type Command interface {
	execute()
}

type LightOffCommand struct {
	light Light
}

func (OffCommand LightOffCommand) execute() {
	OffCommand.light.turnOff()
}

// Concrete commands

type LightOnCommand struct {
	light Light
}

func (OnCommand LightOnCommand) execute() {
	OnCommand.light.turnOn()
}



// Receiver 
type Light struct{
}

func (l *Light) turnOn(){
	fmt.Println("Light is on")
}

func (l *Light) turnOff(){
	fmt.Println("Light is off")
}

// Invoker 

type RemoteControl struct {
	command Command
}

func (rc *RemoteControl) setCommand(c Command){
	rc.command = c
}

func (rc RemoteControl) pressButton() {
	rc.command.execute()
}

func main(){
	light := Light{}
	lightOn := LightOnCommand{light}
	lightOff := LightOffCommand{light}
	
	remote := RemoteControl{nil}
	remote.setCommand(lightOn)
	remote.pressButton()
	remote.setCommand(lightOff)
	remote.pressButton()
}
