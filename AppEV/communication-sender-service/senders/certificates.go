package senders

import (
	c "communication-sender/config"
	"communication-sender/models"
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
)

func SendOpenCertificate(openCertificate models.StartCertificate) {
	c.Setup()
	mailRecipients := viper.GetStringSlice("mail_open_certificate_recipients")
	jsonCertificate, _ := json.Marshal(openCertificate)

	for _, mailRecipient := range mailRecipients {
		fmt.Println("Mail to: " + mailRecipient)
		fmt.Println("Content: Election Opening Certificate")
		fmt.Println(string(jsonCertificate))
		fmt.Println("---------------------------------")
		fmt.Println("---------------------------------")
	}
}

func SendCloseCertificate(closeCertificate models.CloseCertificate) {
	c.Setup()
	mailRecipients := viper.GetStringSlice("mail_close_certificate_recipients")
	jsonCertificate, _ := json.Marshal(closeCertificate)

	for _, mailRecipient := range mailRecipients {
		fmt.Println("Mail to: " + mailRecipient)
		fmt.Println("Content: Election Ending Certificate")
		fmt.Println(string(jsonCertificate))
		fmt.Println("---------------------------------")
		fmt.Println("---------------------------------")
	}
}
