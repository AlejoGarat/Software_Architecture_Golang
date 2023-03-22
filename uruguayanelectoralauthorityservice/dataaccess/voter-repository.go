package dataaccess

import (
	"fmt"
	models "uruguayanelectoralauthorityservice/models"
)

type VoterRepository struct {
	voterList []models.Voter
}

func (voterRepository *VoterRepository) GetVoterById(voterId string) (models.Voter, error) {
	for _, voter := range voterRepository.voterList {
		if voter.IdDocumentVoter == voterId {
			return voter, nil
		}
	}

	return models.Voter{}, fmt.Errorf("No se encontró el votante con el documento especificado: %s", voterId)
}

func (voterRepository *VoterRepository) AddVoter(voter models.Voter) error {
	voterRepository.voterList = append(voterRepository.voterList, voter)
	return nil
}
