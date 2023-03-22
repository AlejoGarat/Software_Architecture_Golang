package senders

import (
	"communication-sender/models"
	"encoding/json"
	"fmt"
)

func SendMail(constancy models.Constancy) {
	fmt.Println("Constancia de voto para : " + constancy.VoterDocument)
	jsonConstancy, _ := json.Marshal(constancy)
	fmt.Println(string(jsonConstancy))
	fmt.Println("----------------------------------------------------")
	fmt.Println("----------------------------------------------------")
}
