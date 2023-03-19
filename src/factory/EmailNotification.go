package factory

import "fmt"

type EmailNotification struct {
}

func (EmailNotification) String() string {
	return fmt.Sprint(
		"====================" +
			"EMAIL NOTIFICATION" +
			"====================")
}

func (EmailNotification) SendNotification(msg string) {
	fmt.Printf("NOTIFICATION: %s\n", msg)
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "EMAIL METHOD"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return fmt.Sprint("Outlook")
}
