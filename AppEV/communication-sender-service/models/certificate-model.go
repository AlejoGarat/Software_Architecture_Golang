package models

import "time"

type StartCertificate struct {
	StartDate        time.Time `json:"start_date"`
	PoliticalParties any       `json:"political_parties"`
	Candidates       any       `json:"candidates"`
	VotersAmmount    int       `json:"voters_amount"`
	VotationMode     string    `json:"votation_mode"`
}

type CloseCertificate struct {
	StartDate      time.Time      `json:"start_date"`
	EndDate        time.Time      `json:"end_date"`
	VoterAmount    int            `json:"voters_amount"`
	VoteAmount     int            `json:"votes_amount"`
	CandidateVotes map[string]int `json:"candidate_votes"`
	PartyVotes     map[string]int `json:"party_votes"`
}
