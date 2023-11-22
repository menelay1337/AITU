package main

import "fmt"

type Client struct {
}

func (c *Client) PluggingLightning(com Computer){
	fmt.Println("Plugging Lightning to machine")
	com.PluggedLightning()
}

type Computer interface {
	PluggedLightning()
}

type Mac struct {
}

func (m *Mac) PluggedLightning(){
	fmt.Println("Plugged lightning to Mac machine\n")
}

type Win struct {
}

func (w *Win) PluggedUSB(){
	fmt.Println("Plugged to USB in Win.")
}

type WinAdapter struct {
	win *Win
}

func (adap *WinAdapter) PluggedLightning(){
	fmt.Println("Converted lightning to USB")
	adap.win.PluggedUSB()
}

func main() {
	client := &Client{}
	mac := &Mac{}
	windows := &Win{}

	client.PluggingLightning(mac)

	winadap := &WinAdapter{windows}
	client.PluggingLightning(winadap)

}
