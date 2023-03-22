package adaptees

import (
	"election-service/models/write"
	"encoding/json"
)

type ElectionJsonConverter struct {
}

func (json_converter *ElectionJsonConverter) ConvertFromJson(jsonElection []byte) (write.CompleteElection, error) {
	var completeElection write.CompleteElection
	err := json.Unmarshal(jsonElection, &completeElection)

	return completeElection, err
}
