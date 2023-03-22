package write

type AlertConfiguration struct {
	ElectionId         string   `json:"election_id"`
	MaxVoteAmount      int      `json:"max_vote_amount"`
	MaxConstancyAmount int      `json:"max_constancy_amount"`
	MailRecipients     []string `json:"mail_recipients"`
}
