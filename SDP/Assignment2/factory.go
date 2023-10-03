package main

import "fmt"

type NotificationService interface {
	SendNotification(message string) 
}

type EmailService struct {}

func (s EmailService) SendNotification(message string) {
	fmt.Printf("message '%s' sended by email.\n", message)
}

type SmsService struct {}

func (s SmsService) SendNotification(message string) {
	fmt.Printf("message '%s' sended by sms.\n", message)
}

type NotificationServiceFactory interface {
	CreateNotificationService() NotificationService
}

type EmailServiceFactory struct {}

func (f EmailServiceFactory) CreateNotificationService() NotificationService {
	return EmailService{}
}

type SmsServiceFactory struct {}

func (f SmsServiceFactory) CreateNotificationService() NotificationService {
	return SmsService{}
}

func main() {

	smsf := SmsServiceFactory{}
	smsS := smsf.CreateNotificationService()
	smsS.SendNotification("Hello buddy!")

	emailF := EmailServiceFactory{}
	emailS := emailF.CreateNotificationService()
	emailS.SendNotification("How are you!")
}
