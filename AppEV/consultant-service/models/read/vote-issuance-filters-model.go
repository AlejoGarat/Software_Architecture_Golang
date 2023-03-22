package read

type VIFilter = string
type VIFilters = []VIFilter

type VoteIssuanceFilters struct {
	VIFilters VIFilters `json:"vote_filters"`
}
