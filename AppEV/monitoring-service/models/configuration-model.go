package models

type MaxVoteAmount = int
type MaxConstancyAmount = int
type MailRecipients = []string

type ConfigurationModel struct {
	MaxVoteAmount      MaxVoteAmount  `json:"max_vote_amount" bson:"max_vote_amount"`
	MaxConstancyAmount MaxVoteAmount  `json:"max_constancy_amount" bson:"max_constancy_amount"`
	MailRecipients     MailRecipients `json:"mail_recipients" bson:"mail_recipients"`
}
