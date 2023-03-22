package pipe

import (
	"context"
	"sync"
	idataaccess "votation-service/dataaccess/interfaces"
	"votation-service/models/write"
	"votation-service/pipe/filters"
	ivoteFilter "votation-service/pipe/interfaces"

	"golang.org/x/sync/errgroup"
)

type Pipeline struct {
	filters           []ivoteFilter.VoteFilter
	voteRepository    idataaccess.VoteRepository
	filtersRepository idataaccess.FiltersRepository
}

func NewPipeline(voteRepository idataaccess.VoteRepository, filtersRepository idataaccess.FiltersRepository) *Pipeline {
	return &Pipeline{voteRepository: voteRepository, filtersRepository: filtersRepository}
}

func (p *Pipeline) AddFilter(f ...ivoteFilter.VoteFilter) {
	p.filters = append(p.filters, f...)
}

func (p *Pipeline) GetFilters() []ivoteFilter.VoteFilter {
	return p.filters
}

func (p *Pipeline) AddFiltersToExecute() {
	filters, _ := p.filtersRepository.GetFilters()
	p.filtersToApply(filters.Filters)
}

func (p *Pipeline) filtersToApply(importedFilters []string) {
	for _, f := range importedFilters {
		switch f {
		case "valid_candidate_filter":
			p.AddFilter(filters.NewCandidateVoter(p.voteRepository))
		case "valid_circuit_filter":
			p.AddFilter(filters.NewCircuitVoter(p.voteRepository))
		case "valid_election_filter":
			p.AddFilter(filters.NewEnabledVoter(p.voteRepository))
		case "unique_candidate_filter":
			p.AddFilter(filters.NewUniqueCandidate(p.voteRepository))
		case "unique_votation_mode_filter":
			p.AddFilter(filters.NewUniqueVotationMode(p.voteRepository))
		case "election_has_ended_filter":
			p.AddFilter(filters.NewEndedElection(p.voteRepository))
		}
	}
}

func (p *Pipeline) FeedFilters(done chan struct{}) chan ivoteFilter.VoteFilter {
	out := make(chan ivoteFilter.VoteFilter)

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

func (p *Pipeline) ExecuteFilters(vote write.Vote) error {
	p.AddFiltersToExecute()
	done := make(chan struct{})
	defer close(done)
	inChannel := p.FeedFilters(done)

	g, _ := errgroup.WithContext(context.Background())

	for filter := range inChannel {
		f := filter

		g.Go(func() error {
			return f.Filter(vote)
		})
	}

	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

func (p *Pipeline) Process(filter ivoteFilter.VoteFilter, vote write.Vote, outChannel chan error, done chan struct{}, wg sync.WaitGroup) {
	err := filter.Filter(vote)

	if err != nil {
		outChannel <- err
	}
}
