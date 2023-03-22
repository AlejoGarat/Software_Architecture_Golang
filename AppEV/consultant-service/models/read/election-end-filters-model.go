package read

type EEFilter = string
type EEFilters = []EEFilter

type ElectionEndFilters struct {
	EEFilters EEFilters `json:"filters"`
}
