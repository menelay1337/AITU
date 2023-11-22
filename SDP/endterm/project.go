package main
import (
	"fmt"
	"sync"
)

type Device interface {
	Update(device string, action string)
	getStatus()
}

type thermostat struct {
	name string
	status string
	currentTemp float32
}

func (t *thermostat) getStatus(){
	fmt.Printf("%s device has %s status.\n", t.name, t.status)
}


func (t *thermostat) Update(device string, action string) {
	if (device != "thermostat") {
		return
	} else {
		switch action {
			case "off":
				if (t.status == "off") {
					fmt.Println("Thermostat already turned off!")
				} else {
					fmt.Println("Turning off thermostat...")
					t.status = "off"
				}
			case "on":
				if (t.status == "on") {
					fmt.Println("Thermostat already turned on!")
				} else if (t.status == "off") {
					fmt.Println("Turning on thermostat...")
					t.status = "on"
				} else {
					fmt.Println("Thermostat is working...")
				}
			case "fahrenheit":
				if (t.status == "on") {
					fmt.Printf("Current temperature is: %d Fahrenheit.\n", int((9.0/5.0)*float32(t.currentTemp)+ 32.0))
				} else {
					fmt.Println("Turn on thermostat")
				}
			case "celcius":
				if (t.status == "on") {
					fmt.Printf("Current temperature is: %2.2f Celcius\n", t.currentTemp)
				} else {
					fmt.Println("Turn on thermostat")
				}
			default: 
				fmt.Println("Incorrect action!")
		}
	}
}


type doorLocker struct {
	name string
	status string
}	

func (d *doorLocker) getStatus(){
	fmt.Printf("%s device has %s status.\n", d.name, d.status)
}

func (d *doorLocker) Update(device string, action string) {
	if ( device != "doorLocker" ) {
		return
	} else {
		switch action {

			case "lock":
				if (d.status == "unlock") {
					d.status = "lock"
					fmt.Println("Door is locked!")
				} else if ( d.status == "lock") {
					fmt.Println("Door is already locked!")
				}

			case "unlock":
				if (d.status == "lock") {
					d.status = "unlock"
					fmt.Println("Door is unlocked!")
				} else if ( d.status == "lock") {
					fmt.Println("Door is already unlocked!")
				}

			default: 
				fmt.Println("Incorrect action!")
		}
	}
}

type lighting struct {
	name string
	status string
}

func (l *lighting) getStatus(){
	fmt.Printf("%s device has %s status.\n", l.name, l.status)
}

func (l *lighting) Update(device string, action string) {
	if ( device != "lighting" ) {
		return 
	} else {
		switch action {

			case "on":
				if (l.status == "on") {
					fmt.Println("Light is already on!")
				} else if (l.status == "off") {
					fmt.Println("Light is on now!")
					l.status = "on"
				}
			
			case "off":
				if (l.status == "off") {
					fmt.Println("Light is already off!")
				} else if (l.status == "on") {
					fmt.Println("Light is off now!")
					l.status = "off"
				}

			default:
				fmt.Println("Incorrect action")
		}
	}
}

type Devices interface {
	addDevice(Device)
	removeDevice(Device)
	UpdateDevices(device string, action string)
}

type  ConcreteDevices struct {
	devices []Device
	mut sync.Mutex
}

// Adding device for device list

func (d *ConcreteDevices) addDevice(device Device) {
	d.mut.Lock()
	defer d.mut.Unlock()
	d.devices = append(d.devices, device)
}

// Removing device from Device list

func (d *ConcreteDevices) removeDevices(device Device) {
	d.mut.Lock()
	defer d.mut.Unlock()
	for i, o := range d.devices {
		if o == device {
			d.devices = append(d.devices[:i], d.devices[i+1:]...)
			break
		}
	}
}

// Sending action for some device
func (d *ConcreteDevices) UpdateDevices(device string, action string) {
	d.mut.Lock()
	defer d.mut.Unlock()
	for _, dev := range d.devices {
		dev.Update(device, action)
	}
}



type SmartHomeMonitor struct {
	container ConcreteDevices
}

// Singleton instance
var instance *SmartHomeMonitor
var once sync.Once

// Instance constructor 
func GetSmartHomeMonitor() *SmartHomeMonitor {
	once.Do(func() {
		instance = &SmartHomeMonitor{}
	})
	return instance
}

// Facade pattern

// SmartHome provides a simplified interface for interacting with the SmartHomeMonitor.
type SmartHome struct {
	status string
	monitor *SmartHomeMonitor
}

// getInstance method for facade 
func NewSmartHomeFacade() *SmartHome {
	return &SmartHome{
		monitor: GetSmartHomeMonitor(),
	}
}

func (f *SmartHome) enterHome(){
	f.status = "entered"
	fmt.Println("Home owner is entering...")
	f.monitor.container.UpdateDevices("doorLocker", "unlock")
	f.monitor.container.UpdateDevices("thermostat", "on")
	f.monitor.container.UpdateDevices("lighting", "on")
	f.monitor.container.UpdateDevices("doorLocker", "lock")
	fmt.Println("Home owner has entered.")
}

func (f *SmartHome) leaveHome(){
	f.status = "left"
	fmt.Println("Home owner is leaving...")
	f.monitor.container.UpdateDevices("doorLocker", "unlock")
	f.monitor.container.UpdateDevices("thermostat", "off")
	f.monitor.container.UpdateDevices("lighting", "off")
	f.monitor.container.UpdateDevices("doorLocker", "lock")
	fmt.Println("Home owner has left.")
}

func (f *SmartHome) getStatus(){
	for _, elem := range f.monitor.container.devices {
		elem.getStatus()
	}
}



func main(){
	light := &lighting{"lighting", "on"}
	door := &doorLocker{"doorLocker", "lock"}	
	thermostat1 := &thermostat{"thermostat", "on", 18.0}
		
	devices := ConcreteDevices{}
	devices.addDevice(light)
	devices.addDevice(door)
	devices.addDevice(thermostat1)

	SmartHome := NewSmartHomeFacade()
	SmartHome.monitor.container = devices
	
	fmt.Println("\nCLI is starting. . .\n")
	fmt.Println("Hey user you have yourself Smart home!")
	fmt.Println("To handle your Home through CLI you have 6 options in general")
	fmt.Println("to leave this CLI enter \"done\" keyword.")

	var input string
	for input != "done" {
		fmt.Println("enter | leave | status | lighting | door | thermostat | done")
		fmt.Print("Please enter your next command: ")
		fmt.Scan(&input)
		if ( input == "done" ) {
			break
		}
		switch input {
			case "status":
				fmt.Println()
				SmartHome.getStatus()

			case "enter":
				if (SmartHome.status == "entered") {
					fmt.Println("\nYou're already entered.")
					break
				}
				
				fmt.Println()
				SmartHome.enterHome()

			case "leave":
				if (SmartHome.status == "left") {
					fmt.Println("\nYou're already left.")
					break
				}

				fmt.Println()
				SmartHome.leaveHome()

			case "thermostat":

				for {
					fmt.Println("\nChoose \"fahrenheit\" or \"celcius\" to obtain current temperuture in house.")
					fmt.Print("Or choose \"on\" or \"off\" options to turn thermostat. [done is available] : ")
					fmt.Scan(&input)
					fmt.Println()
					if input == "done" {
						break
					} else if (input == "fahrenheit" || input == "celcius" || input == "on" || input == "off") {
						SmartHome.monitor.container.UpdateDevices("thermostat", input)
						break
					}
				}
				
				input = ""

			case "lighting":
				for {
					fmt.Print("Choose to turn \"on\" or \"off\" light at home. [done is available] :  ")
					fmt.Scan(&input)
					if input == "done" {
						break
					} else if (input == "on" || input == "off") {
						SmartHome.monitor.container.UpdateDevices("lighting", input)
						break
					}
				}
				
				input = ""

			case "door":
				for {
					fmt.Print("Choose to \"lock\" or \"unlock\" door at home. [done is available] :  ")
					fmt.Scan(&input)
					if input == "done" {
						break
					} else if (input == "lock" || input == "unlock") {
						SmartHome.monitor.container.UpdateDevices("doorLocker", input)
						break
					}
				}
				
				input = ""

			default: 
				fmt.Println("Incorrect input!")
			
		}
		fmt.Println()
	}
	fmt.Println("\nCLI is finished.\n")
}
