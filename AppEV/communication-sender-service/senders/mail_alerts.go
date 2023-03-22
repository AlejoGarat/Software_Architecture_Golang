package senders

import (
	c "communication-sender/config"
	"fmt"

	"github.com/spf13/viper"
)

func SendMailAlert(message string) {
	c.Setup()
	mailRecipients := viper.GetStringSlice("mail_alert_recipients")

	for _, mailRecipient := range mailRecipients {
		fmt.Println("Mail to: " + mailRecipient)
		fmt.Println("Alert: " + message)
		fmt.Println("---------------------------------")
		fmt.Println("---------------------------------")
	}
}
