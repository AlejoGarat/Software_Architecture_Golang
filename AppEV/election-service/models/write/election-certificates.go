package write

import (
	"election-service/models/read"
	"time"
)

type StartCertificate struct {
	StartDate        time.Time        `json:"start_date"`
	PoliticalParties []PoliticalParty `json:"political_parties"`
	Candidates       []Candidate      `json:"candidates"`
	VotersAmmount    int              `json:"voters_amount"`
	VotationMode     string           `json:"votation_mode"`
}

type CloseCertificate struct {
	StartDate      time.Time                 `json:"start_date"`
	EndDate        time.Time                 `json:"end_date"`
	VoterAmount    int                       `json:"voters_amount"`
	VoteAmount     int                       `json:"votes_amount"`
	CandidateVotes []read.StringIntCandidate `json:"candidate_votes"`
	PartyVotes     []read.StringIntParty     `json:"party_votes"`
}
