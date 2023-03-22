package read

type AlertConfiguration struct {
	MaxVoteAmount      int      `json:"max_vote_amount"`
	MaxConstancyAmount int      `json:"max_constancy_amount"`
	MailRecipients     []string `json:"mail_recipients"`
}
