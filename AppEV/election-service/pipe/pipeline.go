package pipe

import (
	"context"
	idataaccess "election-service/dataaccess/interfaces"
	"election-service/models/write"
	"election-service/pipe/filters"
	ielectionFilter "election-service/pipe/interfaces"

	"golang.org/x/sync/errgroup"
)

type Pipeline struct {
	filters    []ielectionFilter.ElectionFilter
	repository idataaccess.ElectionRepository
}

func NewPipeline(repository idataaccess.ElectionRepository) *Pipeline {
	pipeline := &Pipeline{repository: repository}

	return pipeline
}

func (p *Pipeline) addFilter(f ...ielectionFilter.ElectionFilter) {
	p.filters = append(p.filters, f...)
}

func (p *Pipeline) GetFilters() []ielectionFilter.ElectionFilter {
	return p.filters
}

func (p *Pipeline) AddFiltersToExecute(importedFilters []string) {
	for _, f := range importedFilters {
		switch f {
		case "candidates_in_party_filter":
			p.addFilter(filters.NewCandidatesInParty())
		case "candidates_and_voters_filter":
			p.addFilter(filters.NewCandidatesAndVoters())
		case "correct_votation_mode_filter":
			p.addFilter(filters.NewCorrectVotationMode())
		case "party_has_candidate_filter":
			p.addFilter(filters.NewPartyHasCandidate())
		case "political_parties_filter":
			p.addFilter(filters.NewPoliticalParties())
		case "valid_date_filter":
			p.addFilter(filters.NewValidDate())
		case "close_election_date":
			p.addFilter(filters.NewCloseElectionDate())
		case "total_vote_amount":
			p.addFilter(filters.NewTotalVoteAmount(p.repository))
		}

	}
}

func (p *Pipeline) process(filter ielectionFilter.ElectionFilter, election write.CompleteElection, outChannel chan error, done chan struct{}) {
	err := filter.Filter(election)

	if err != nil {
		outChannel <- err
	}
}

func (p *Pipeline) analyzeChannel(election write.CompleteElection, outChannel chan error, done chan struct{}) bool {
	select {
	case <-outChannel:
		return true
	case <-done:
		return false
	}
}

func (p *Pipeline) feedFilters(done chan struct{}) chan ielectionFilter.ElectionFilter {
	out := make(chan ielectionFilter.ElectionFilter)

	filters := p.GetFilters()

	go func() {
		defer close(out)
		for _, f := range filters {
			select {
			case out <- f:
			case <-done:
				return
			}
		}
	}()

	return out
}

func (p *Pipeline) ExecuteElectionFilters(election write.CompleteElection) error {
	done := make(chan struct{})
	defer close(done)
	inChannel := p.feedFilters(done)

	g, _ := errgroup.WithContext(context.Background())

	for filter := range inChannel {
		f := filter

		g.Go(func() error {
			return f.Filter(election)
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	return nil

}
