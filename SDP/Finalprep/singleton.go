package main

import "fmt"

type Indicator interface {
	getStatus()
}

type GPUindicator struct {
	temp int
	load int
	frequency int
}

func (i *GPUindicator) getStatus(){
	fmt.Printf("GPU: TEMP-%d- LOAD-%d- FREQ-%dmHZ-.\n",i.temp, i.load, i.frequency)
}

type CPUindicator struct {
	temp int
	load int
	frequency int
}

func (i *CPUindicator) getStatus(){
	fmt.Printf("CPU: TEMP-%d- LOAD-%d- FREQ-%dmHZ-.\n",i.temp, i.load, i.frequency)
}
var cpuind CPUindicator = CPUindicator{75, 100, 4400}
var gpuind GPUindicator = GPUindicator{100, 100, 6000}

var instance *SystemMonitor
var standardInds []Indicator = []Indicator{&cpuind, &gpuind}


func newSystemMonitor() *SystemMonitor {
	if (instance == nil) {
		instance = &SystemMonitor{standardInds}
	}
	return instance
}

type Monitor interface {
	getStatus()
}

type SystemMonitor struct {
	indicators []Indicator		
}

func (mon *SystemMonitor) getStatus(){
	for _, elem := range mon.indicators {
		elem.getStatus()
	}
}

func main() {
	myMonitor := newSystemMonitor()
	myMonitor.getStatus()

	secondMonitor := newSystemMonitor()

	if (myMonitor == secondMonitor){
		fmt.Println("Constructor works properly!")
	} else {
		fmt.Println("Constructor doesn't work properly")
	}
}
