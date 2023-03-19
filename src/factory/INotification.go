package factory

type INotificationFactory interface {
	SendNotification(msg string)
	GetSender() ISender
}
