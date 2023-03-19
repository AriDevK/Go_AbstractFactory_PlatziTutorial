package main

import (
	"Go_AbstractFactory_PlatziTutorial/factory"
	"fmt"
)

func main() {
	fmt.Println("Starting...")
	defer fmt.Println("...End")

	smsFactory, _ := factory.GetNotificationFactory("SMS")
	emailFactory, _ := factory.GetNotificationFactory("EMAIL")

	fmt.Println(smsFactory)
	factory.SendNotification(smsFactory, "HELLO")
	factory.GetMethod(smsFactory)
	factory.GetChannel(smsFactory)

	fmt.Println(emailFactory)
	factory.SendNotification(emailFactory, "WORLD")
	factory.GetMethod(emailFactory)
	factory.GetChannel(emailFactory)
}
