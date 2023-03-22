package usecases

import (
	"crypto/sha256"
	idataaccess "data-update-service/dataaccess/interfaces"
	imemorydataaccess "data-update-service/memorydataaccess/interfaces"
	"data-update-service/models/read"
	"errors"
	"fmt"
	"time"

	age "github.com/bearbin/go-age"
	"github.com/robfig/cron/v3"
)

const (
	yearPosition  = 1
	monthPosition = 2
	dayPosition   = 3
)

type VoteDataUseCase struct {
	voteDataRepository idataaccess.VoteDataRepository
	memoryRepository   imemorydataaccess.MemoryRepository
}

func NewVoteDataUseCase(voteDataRepository idataaccess.VoteDataRepository, memoryRepo imemorydataaccess.MemoryRepository) *VoteDataUseCase {
	return &VoteDataUseCase{voteDataRepository: voteDataRepository, memoryRepository: memoryRepo}
}

func (voteDataUseCase *VoteDataUseCase) UpdateVoteData(vote read.Vote) error {

	err := voteDataUseCase.voteDataRepository.UpdateSchedule(vote)

	if err != nil {
		return err
	}

	/*if voterHasVoted := voteDataUseCase.voteDataRepository.VoterHasVoted(hashData(vote.VoterDocument), vote.ElectionId); voterHasVoted {
		return nil
	}*/

	voter, err := voteDataUseCase.voteDataRepository.GetVoterByDocument(vote.VoterDocument, vote.ElectionId)

	if err != nil {
		return err
	}

	birthDate, err := time.Parse("2006-01-02", voter.DateOfBirth)

	if err != nil {
		return err
	}

	voter.Age = age.Age(birthDate)

	if err := voteDataUseCase.voteDataRepository.UpdateCircuitAgeVotes(vote, voter); err != nil {
		return err
	}

	if err := voteDataUseCase.voteDataRepository.UpdateCircuitGenderVotes(vote, voter); err != nil {
		return err
	}

	if err := voteDataUseCase.voteDataRepository.UpdateDepartmentData(vote, voter); err != nil {
		return err
	}

	err = voteDataUseCase.voteDataRepository.UpdateTotalVotes(vote.ElectionId)

	if err != nil {
		return err
	}

	err = voteDataUseCase.memoryRepository.UpdateTotalVotes(vote.ElectionId)

	return err
}

func (vuc *VoteDataUseCase) StartUpdateCrone(election read.Election) error {
	c := cron.New()
	c.AddFunc("@every 5m", func() {
		if time.Now().Before(election.EndDate) {
			vuc.updateMemoryRepositoryData(election.ElectionId)
		} else {
			c.Stop()
		}
	})
	c.Start()

	return nil
}

func (vuc *VoteDataUseCase) updateMemoryRepositoryData(electionId string) error {
	votesPerParty, votesPerCandidate, votesPerDepartment, err := vuc.CalculateVotes(electionId)

	if err != nil {
		return err
	}

	err = vuc.memoryRepository.UpdateVotesPerParty(votesPerParty, electionId)

	if err != nil {
		return err
	}

	err = vuc.memoryRepository.UpdateVotesPerCandidate(votesPerCandidate, electionId)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = vuc.memoryRepository.UpdateVotesPerDepartment(votesPerDepartment, electionId)

	return err
}

func (vuc *VoteDataUseCase) CalculateVotes(electionId string) ([]read.StringIntParty, []read.StringIntCandidate, []read.StringIntDepartment, error) {
	votesParty, votesCandidate, votesDepartment := make(map[string]int), make(map[string]int), make(map[string]int)
	var votesPerParty []read.StringIntParty
	var votesPerCandidate []read.StringIntCandidate
	var votesPerDepartment []read.StringIntDepartment

	votes, err := vuc.voteDataRepository.GetVotes(electionId)

	if err != nil {
		return votesPerParty, votesPerCandidate, votesPerDepartment, err
	}

	for _, vote := range votes {
		candidate, _ := vuc.getCandidate(vote.CandidateId, electionId)
		department, _ := vuc.getDepartment(vote.CircuitId, electionId)
		votesParty[candidate.PoliticalPartyId] = votesParty[candidate.PoliticalPartyId] + 1
		votesCandidate[candidate.CandidateId] = votesCandidate[candidate.CandidateId] + 1
		votesDepartment[department] = votesDepartment[department] + 1
	}

	for party, votes := range votesParty {
		votesPerParty = append(votesPerParty, read.StringIntParty{Id: party, Amount: votes})
	}

	for candidate, votes := range votesCandidate {
		votesPerCandidate = append(votesPerCandidate, read.StringIntCandidate{Id: candidate, Amount: votes})
	}

	for department, votes := range votesDepartment {
		votesPerDepartment = append(votesPerDepartment, read.StringIntDepartment{Id: department, Amount: votes})
	}

	return votesPerParty, votesPerCandidate, votesPerDepartment, nil
}

func (vuc *VoteDataUseCase) getCandidate(hashedCandidateId string, electionId string) (read.Candidate, error) {
	var cand read.Candidate
	candidates, _ := vuc.voteDataRepository.GetCandidates(electionId)

	for _, candidate := range candidates {
		if hashData(candidate.CandidateId) == hashedCandidateId {
			return candidate, nil
		}
	}

	return cand, errors.New("Candidate does not exist.")
}

func (vuc *VoteDataUseCase) getDepartment(circuitId string, electionId string) (string, error) {
	var dept string
	circuits, err := vuc.voteDataRepository.GetCircuits(electionId)

	if err != nil {
		return dept, err
	}

	for _, circuit := range circuits {
		if circuit.Id == circuitId {
			return circuit.Department, nil
		}
	}

	return dept, errors.New("Department does not exist.")
}

func hashData(data string) string {
	encryptedData := sha256.New()

	encryptedData.Write([]byte(data))

	return fmt.Sprintf("%x", encryptedData.Sum(nil))
}
