package adapters

import (
	"election-service/models/write"
	"encoding/json"
)

type ElectionJsonAdapter struct {
}

type Response struct {
	ElectoralAuthorityData string `json:"ElectoralAuthorityData"`
	StatusCode             int    `json:"Status Code"`
}

func (json_adapter ElectionJsonAdapter) ConvertElection(body []byte) (write.CompleteElection, error) {
	var response Response

	var election write.CompleteElection

	err := json.Unmarshal(body, &response)

	if err != nil {
		return election, err
	}

	err = json.Unmarshal([]byte(response.ElectoralAuthorityData), &election)

	return election, err
}
