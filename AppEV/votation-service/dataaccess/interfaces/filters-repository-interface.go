package interfaces

import "votation-service/models/read"

type FiltersRepository interface {
	GetFilters() (read.VoteFilters, error)
}
