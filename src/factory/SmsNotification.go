package factory

import "fmt"

type SmsNotification struct {
}

func (SmsNotification) String() string {
	return fmt.Sprint(
		"====================" +
			"SMS NOTIFICATION" +
			"====================")
}

func (SmsNotification) SendNotification(msg string) {
	fmt.Printf("NOTIFICATION: %s\n", msg)
}

func (SmsNotification) GetSender() ISender {
	return SmsNotificationSender{}
}

type SmsNotificationSender struct {
}

func (SmsNotificationSender) GetSenderMethod() string {
	return "SMS METHOD"
}

func (SmsNotificationSender) GetSenderChannel() string {
	return fmt.Sprint("Twilio")
}
