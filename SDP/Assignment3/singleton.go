package main
import (
	"fmt"
)

type counter struct {
	// private 
	counterMap map[string]int
}

// Our single instance
var myCounter *counter

func (c counter) getInstance() *counter {
    if myCounter == nil {
	    fmt.Println("Counter just created.")
	    myCounter = &counter{make(map[string]int)}
     } else {
	    fmt.Println("Returning existing instance.")
    }
    return myCounter 
}

func (c *counter) addType(myType string) {
	if _, ok := c.counterMap[myType]; ok {
    		c.counterMap[myType]++
	} else {
		c.counterMap[myType] = 1
	}
}

func (c *counter) getCount(myType string) {
	if val, ok := c.counterMap[myType]; ok {
		fmt.Println(val, " is count of ", myType)
	} else {
		fmt.Println("Error, no such types.")
		
	}
}
	



func main() {
	cnt := myCounter.getInstance()
	cnt.addType("array")
	cnt.addType("array")
	cnt.addType("Person")
	cnt.addType("Chemical")
	cnt.getCount("array")
	cnt.getCount("Person")
	cnt.getCount("Chemical")
	
	
}

