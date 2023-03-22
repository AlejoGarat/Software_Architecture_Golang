package read

type Filters = []string

type VoteFilters struct {
	Filters Filters `json:"vote_filters"`
}
