package read

type EBFilter = string
type EBFilters = []EBFilter

type ElectionBeginningFilters struct {
	EBFilters EBFilters `json:"filters"`
}
