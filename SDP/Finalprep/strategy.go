package main

import "fmt"
import "math/rand"
import "time"

type Error interface {
}

type Payment interface {
	Transfer(int) (error)
}

type Wallet struct {
	cash int	
}

func (w *Wallet) Transfer(sum int) (err error) {
	if (w.cash < sum) {
		return fmt.Errorf("Oops i don't have enough cash.")
	}
	w.cash -= sum
	fmt.Printf("Transfered by hand\nI calculated that it's left  %d funds in my wallet.\n", w.cash)
	return nil
}

type Bankcard struct {
	total int
}

func (b *Bankcard ) Transfer(sum int) (err error) {
	if (b.total < sum) {
		return fmt.Errorf("Not enough funds.")
	}
	b.total -= sum
	fmt.Printf("Transfered succesfully!\nIt's left  %d funds.\n", b.total)
	return nil
}

type Creditcard struct {
	total int
	debt int
}

func (c *Creditcard ) Transfer(sum int) (err error) {
	if (c.total < sum) {
		return fmt.Errorf("Your payment is not approved.")
	}
	c.total -= sum/3
	c.debt += sum
	fmt.Printf("Transfered succesfully!\nPayment approved\nYour debt: %d.\n", c.debt)
	return nil
}

type Client struct {
	payment Payment
} 

func (c *Client) Pay(sum int) (err error){
	err = c.payment.Transfer(sum)
	if (err != nil) {
		fmt.Printf("Payment with %T has failed\n", c.payment)
		return err
	}
	
	fmt.Printf("Payment past successfully with %T.\n", c.payment)
	return nil

}

func (client *Client) paymentChange(paymethod Payment){
	client.payment = paymethod
}

func randInt(min int, max int) int {
	    return min + rand.Intn(max-min)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano()) 
	MyWallet := &Wallet{1000}
	bankcard := &Bankcard{2500}
	creditcard := &Creditcard{5000, 0}
	client := Client{MyWallet}
	
	Sum := randInt(1, 5000)

	myerr := client.Pay(Sum)
	if myerr != nil {
		fmt.Println(myerr)
		client.paymentChange(bankcard)
		myerr = nil
	}

	myerr = client.Pay(Sum)
	if myerr != nil {
		fmt.Println(myerr)
		client.paymentChange(creditcard)
		myerr = nil
	}
	myerr = client.Pay(Sum)
	if myerr != nil {
		fmt.Println(myerr)
		panic("no option to pay")
	}

	MyWallet = &Wallet{1000}
	bankcard = &Bankcard{2500}
	creditcard = &Creditcard{5000, 0}

	// Another option to use
	if (Sum <  MyWallet.cash){
		client.paymentChange(MyWallet)
	} else if (Sum < bankcard.total){
		client.paymentChange(bankcard)
	} else if (Sum < creditcard.total){
		client.paymentChange(creditcard)
	} else {
		panic("No payment options")
	}
	client.Pay(Sum)

}
