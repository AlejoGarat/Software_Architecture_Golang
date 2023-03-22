package senders

import (
	"encoding/json"
	"fmt"
	"message-sender/models"
)

func SendSms(constancy models.Constancy) {
	fmt.Println("Constancia de voto para : " + constancy.VoterDocument)
	fmt.Println("----------------------------------------------------")
	jsonConstancy, _ := json.Marshal(constancy)
	fmt.Println(string(jsonConstancy))
	fmt.Println("----------------------------------------------------")
	fmt.Println("----------------------------------------------------")
}
