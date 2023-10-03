package main
import "fmt"

type notificator interface {
	notify(string, int)
}

type userMail struct {
	email string
}

func (m *userMail) notify(message string, id int) {
	fmt.Printf("'%s' sended as email to %s email with %d id\n", message, m.email, id)
}

type userPhone struct {
	phone string
}

func (p *userPhone) notify(message string, id int) {
	fmt.Printf("'%s' sended as sms to %s number with %d id\n", message, p.phone, id)
}

type userAddress struct{
	homeAddress string
}

func (a *userAddress) notify(message string, id int) {
	fmt.Printf("Ship department: '%s' sended to %s address of user with %d id\n", message, a.homeAddress, id)
}

type receiver struct {
	name string
	id int
	notifier notificator
}

func (user receiver) send(message string) {
	switch user.notifier.(type) {
		
		case *userMail:
			fmt.Println("Sending email . . .")
		case *userPhone:
			fmt.Println("Sending sms to phone number . . .")
		case *userAddress:
			fmt.Println("Sending application to the shipping department . . . ")
		default:
			err := fmt.Errorf("Message wasn't sent!")
			panic(err)
		
	}
	user.notifier.notify(message, user.id)
}
func main() {
	Ken := receiver {
		name: "Ken",
		id: 234123,
		notifier: &userMail{"ken_luganski@raven.com"},
	}
	Ken.send(fmt.Sprintf("Good evening honorable %s, we should notify that you should come to office in these days\n", Ken.name))

	Alisa := receiver {
		name: "Alisa",
		id: 123465,
		notifier: &userPhone{"+7 705 532 23 54"},	
	}

	Alisa.send(fmt.Sprintf("Good evening honorable %s, we should notify that you should come to the office in these days\n", Alisa.name))

	Kim := receiver {
		name: "Kim",
		id: 212342,
		notifier: &userAddress{"Astana city, Street Mangylyk el 45, 54th apartment"},
	}
	
	Kim.send(fmt.Sprintf("Good evening honorable %s, we should notify that you should come to the office in these days\n", Ken.name))

	
}
