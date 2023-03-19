package factory

import "fmt"

func GetNotificationFactory(notificationType string) (INotificationFactory, error) {
	switch notificationType {
	case "SMS":
		return &SmsNotification{}, nil
	case "EMAIL":
		return &EmailNotification{}, nil
	default:
		return nil, fmt.Errorf("NO NOTIFICATION TYPE")
	}
}

func SendNotification(factory INotificationFactory, msg string) {
	factory.SendNotification(msg)
}

func GetMethod(factory INotificationFactory) {
	fmt.Println("METHOD: ", factory.GetSender().GetSenderMethod())
}

func GetChannel(factory INotificationFactory) {
	fmt.Println("CHANNEL", factory.GetSender().GetSenderChannel())
}
