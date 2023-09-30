package main

import "fmt"

type Payer interface {
	Pay(int) error
} 

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("нехватает денег в кошельке")
	}
	w.Cash = w.Cash - amount
	return nil	
}

type Card struct {
	Balance int
	Owner string
	ValidUntil string
	CVV string
	Number string
}

func (c *Card) Pay(amount int) error {
	if c.Balance < amount {
		return fmt.Errorf("нехватает денег на карте")
	}
	c.Balance = c.Balance - amount
	return nil	
}

// Program to an interface not an implementation
// OCP - open for extenstion close for modifications
func Buy(p Payer) {
	switch p.(type) {
	case *Wallet:
		fmt.Println("Оплата наличными.")
	case *Card:
		plasticCard := p.(*Card)
		fmt.Printf("Вставляйте карту, %v \n", plasticCard.Owner)
	default:
		fmt.Println("Что то новое.")
	}
	err := p.Pay(10)
	if err!=nil {
		panic(err)
	}

	fmt.Printf("Спасибо за покупку через %T \n", p)
}


func main() {
	myWallet := Wallet{100}
	myCard := Card{Balance: 100, Owner: "Bob Smith"}

	Buy(&myWallet)
	Buy(&myCard)
}
