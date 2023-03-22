package models

import (
	"time"
)

type IdElection = string
type Description = string

type Election struct {
	IdElection  IdElection  `json:"id_election"`
	Description Description `json:"description"`
	StartTime   time.Time   `json:"start_time"`
	EndTime     time.Time   `json:"end_time"`
}
