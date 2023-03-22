package read

type VoterId = string
type VoteId = string
type ElectionId = string

type ConstancyRequest struct {
	VoterId    VoterId    `json:"voter_id"`
	VoteId     VoteId     `json:"vote_id"`
	ElectionId ElectionId `json:"election_id"`
}
