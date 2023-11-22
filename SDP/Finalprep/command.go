package main

import "fmt"

type Command interface {
	execute()
}

type turnOffCommand struct {
	receiver Receiver
}

func (command *turnOffCommand) execute(){
	command.receiver.Off()
}

type turnOnCommand struct {
	receiver Receiver
}

func (command *turnOnCommand) execute(){
	command.receiver.On()
}


type Receiver interface {
	On()
	Off()
}



func main() {
	
}
