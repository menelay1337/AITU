package main

import "fmt"


type Button interface {
	onClick()
	render()
}

type Dialog interface {
	createButton() Button
}

type WindowsDialog struct {
}


func (wd *WindowsDialog) createButton() Button {
	return &WindowsButton{}
}

type WebDialog struct {
	
}

func (wd *WebDialog) createButton() Button {
	return &HTMLbutton{}
}

type WindowsButton struct {
}

func (wb *WindowsButton) render(){
	fmt.Println("Rendering the button")
}

func (wb *WindowsButton) onClick(){
	fmt.Println("Rendering click on button on OS")
}

type HTMLbutton struct {
}

func (hb *HTMLbutton) render(){
	fmt.Println("Rendering the button in browser")
}

func (hb *HTMLbutton) onClick(){
	fmt.Println("Rendering click on button on Web")
}

func main() {
	win := WindowsDialog{}
	web := WebDialog{}

	btn1 := win.createButton()
	btn2 := web.createButton()
	btn1.render()
	btn1.onClick()
	btn2.render()
	btn2.onClick()
}
