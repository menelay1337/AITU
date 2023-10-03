package main
import "fmt"

type discount interface {
	getDiscount() float32
}

type initialDiscount struct {
}

func (d initialDiscount) getDiscount() float32 {
	return 1.0
}

type summerDiscount struct {
	d discount 
}

func (s summerDiscount) getDiscount() float32 {
	
	disc := s.d.getDiscount()

	if (disc <= 1) && (disc > 0.8){
		disc -= 0.3
		return disc
	} else {
		disc -= 0.1
		return disc
	}
}

type saleDiscount struct {
	d discount
}

func (s *saleDiscount) getDiscount() float32{

	disc := s.d.getDiscount()
	
	if (disc <= 1) && (disc > 0.8){
		disc -= 0.2
		return disc 
	} else {
		return disc
	}
}

func main(){

	startDiscount := initialDiscount{}

	fmt.Println(startDiscount.getDiscount())

	discountSummer := summerDiscount {
		d : startDiscount,
	}

	fmt.Println("Summer: ", discountSummer.getDiscount())
	

	saleDisc := saleDiscount{
		d : startDiscount,
	}

	fmt.Println("Sale: ", saleDisc.getDiscount())
	

	
	
}
