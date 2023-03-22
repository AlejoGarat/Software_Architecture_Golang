package usecases

import (
	iadapters "election-service/adapters/interfaces"
	idataaccess "election-service/dataaccess/interfaces"
	"election-service/helpers"
	imemorydataaccess "election-service/memorydataaccess/interfaces"
	"election-service/models/read"
	"election-service/models/write"
	"election-service/pipe"
	"election-service/rabbit"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	age "github.com/bearbin/go-age"
	cron "github.com/robfig/cron/v3"
)

type ElectionUsecase struct {
	electionRepository    idataaccess.ElectionRepository
	electionAdapter       iadapters.ElectionAdapterTarget
	electionMemoryRepo    imemorydataaccess.ElectionMemoryRepository
	startElectionPipeline *pipe.Pipeline
	endElectionPipeline   *pipe.Pipeline
	helpers               helpers.Helpers
}

func NewElectionUseCase(repo idataaccess.ElectionRepository, adapter iadapters.ElectionAdapterTarget, memoryRepo imemorydataaccess.ElectionMemoryRepository,
	startElectionPipeline *pipe.Pipeline, endElectionPipeline *pipe.Pipeline, helpers helpers.Helpers) *ElectionUsecase {

	return &ElectionUsecase{
		electionRepository:    repo,
		electionAdapter:       adapter,
		electionMemoryRepo:    memoryRepo,
		startElectionPipeline: startElectionPipeline,
		endElectionPipeline:   endElectionPipeline,
		helpers:               helpers,
	}
}

func (electionUsecase *ElectionUsecase) AddElection(electionData read.ElectionData) error {
	var log write.LoggingModel

	err, election := electionUsecase.createElection(electionData)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Add Election", Actor: "Electoral Authority", Description: err.Error()}
		electionUsecase.helpers.LogHelper.SendLog(log)
		return err
	}

	go electionUsecase.createStartCrone(election)

	return nil
}

func (electionUsecase *ElectionUsecase) GetElectionResult(electionId string) (read.ElectionResult, error) {
	var log write.LoggingModel

	electionResult, err := electionUsecase.electionMemoryRepo.GetElectionResult(electionId)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Get Election Result", Actor: "Electoral Authority", Description: err.Error()}
		electionUsecase.helpers.LogHelper.SendLog(log)
		return electionResult, err
	}

	return electionResult, nil
}

func (electionUsecase *ElectionUsecase) createElection(electionData read.ElectionData) (error, write.Election) {
	date, err := time.Parse(time.RFC822, electionData.StartDate+" "+electionData.StartTime+" UTC")

	return err, write.Election{
		Id:              electionData.Id,
		Description:     electionData.Description,
		Url:             electionData.Url,
		StartDate:       date,
		TotalVoteAmount: 0,
	}
}

func (eu *ElectionUsecase) createStartCrone(election write.Election) {
	c := cron.New()
	c.AddFunc("@every 1h", func() {
		eu.checkDate(c, election)
	})
	c.Start()
	eu.checkDate(c, election)

}

func (eu *ElectionUsecase) checkDate(c *cron.Cron, election write.Election) {
	if time.Now().UTC().After(election.StartDate.Add(-time.Hour * 1)) {
		eu.activateElection(election)
		c.Stop()
	}
}

func (eu *ElectionUsecase) activateElection(election write.Election) {
	var log write.LoggingModel

	response, err := http.Get(election.Url)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Activate Election", Actor: "Electoral Authority", Description: err.Error()}
		eu.helpers.LogHelper.SendLog(log)
	}

	electionData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Activate Election", Actor: "Electoral Authority", Description: err.Error()}
		eu.helpers.LogHelper.SendLog(log)
	}

	completeElection, err := eu.electionAdapter.ConvertElection(electionData)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Activate Election", Actor: "Electoral Authority", Description: err.Error()}
		eu.helpers.LogHelper.SendLog(log)
	}

	err = eu.startElectionPipeline.ExecuteElectionFilters(completeElection)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Activate Election", Actor: "Electoral Authority", Description: err.Error()}
		eu.helpers.LogHelper.SendLog(log)
	}

	election.EndDate = completeElection.EndDate
	election.VotationMode = completeElection.VotationMode
	election.EligibleVoters = len(completeElection.Voters)

	err = eu.electionRepository.AddElection(election)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Activate Election", Actor: "Electoral Authority", Description: err.Error()}
		eu.helpers.LogHelper.SendLog(log)
	}

	err = eu.electionMemoryRepo.AddElection(election)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Activate Election", Actor: "Electoral Authority", Description: err.Error()}
		eu.helpers.LogHelper.SendLog(log)
	}

	eu.addCompleteElectionData(completeElection)

	err = eu.openCertificate(completeElection)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Activate Election", Actor: "Electoral Authority", Description: err.Error()}
		eu.helpers.LogHelper.SendLog(log)
	}

	err = eu.signalUpdateMemoryData(completeElection)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Activate Election", Actor: "Electoral Authority", Description: err.Error()}
		eu.helpers.LogHelper.SendLog(log)
	}

	c := cron.New()
	c.AddFunc("@every 1h", func() { eu.closeElection(c, completeElection) })
	c.Start()
}

func (eu *ElectionUsecase) signalUpdateMemoryData(election write.CompleteElection) error {
	jsonElection, err := json.Marshal(election)

	if err != nil {
		return err
	}

	eu.helpers.LogHelper.SendSignal(jsonElection)

	return nil
}

func (eu *ElectionUsecase) addCompleteElectionData(completeElection write.CompleteElection) {
	var log write.LoggingModel

	err := eu.addCandidates(completeElection.Candidates, completeElection.Id)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Add Election Data", Actor: "Electoral Authority", Description: err.Error()}
		eu.helpers.LogHelper.SendLog(log)
	}

	err = eu.addCircuits(completeElection.Circuits)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Add Election Data", Actor: "Electoral Authority", Description: err.Error()}
		eu.helpers.LogHelper.SendLog(log)
	}

	votingDepartments := eu.getCircuitDepartments(completeElection.Circuits)

	err = eu.addVoters(completeElection, votingDepartments)

	votersPerDepartment, votersPerAge, votersPerGender, votersPerDepartmentGender, votersPerDepartmentAge := eu.calculateVotersPerDepartment(completeElection, votingDepartments)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Add Election Data", Actor: "Electoral Authority", Description: err.Error()}
		eu.helpers.LogHelper.SendLog(log)
	}

	err = eu.addPoliticalParties(completeElection.PoliticalParties, completeElection.Id)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Add Election Data", Actor: "Electoral Authority", Description: err.Error()}
		eu.helpers.LogHelper.SendLog(log)
	}

	err = eu.addVotersPerDepartment(completeElection.Id, votersPerDepartment)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Add Election Data", Actor: "Electoral Authority", Description: err.Error()}
		eu.helpers.LogHelper.SendLog(log)
	}

	err = eu.addExtraTables(completeElection, votersPerAge, votersPerGender, votersPerDepartmentGender, votersPerDepartmentAge)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Add Election Data", Actor: "Electoral Authority", Description: err.Error()}
		eu.helpers.LogHelper.SendLog(log)
	}
}

func (eu *ElectionUsecase) addExtraTables(completeElection write.CompleteElection, votersPerAge map[string]map[int]int, votersPerGender map[string]map[string]int,
	votersPerDepartmentGender map[string]map[string]int, votersPerDepartmentAge map[string]map[int]int) error {

	err := eu.electionRepository.AddVotersPerAge(votersPerAge, completeElection.Id)

	if err != nil {
		return err
	}

	err = eu.electionRepository.AddVotersPerGender(votersPerGender, completeElection.Id)

	if err != nil {
		return err
	}

	err = eu.electionRepository.AddVotersPerDepartmentByGenderAndAge(votersPerDepartmentGender, votersPerDepartmentAge, completeElection.Id)

	return err
}

func (eu *ElectionUsecase) addVoters(completeElection write.CompleteElection, votingDepartments map[string]string) error {
	voters := completeElection.Voters

	for _, voter := range voters {
		voter.VotingDepartment = votingDepartments[voter.CircuitId]
		err := eu.electionRepository.AddVoter(voter, completeElection.Id)

		if err != nil {
			return err
		}
	}

	return nil
}

func (eu *ElectionUsecase) getCircuitDepartments(circuits []write.Circuit) map[string]string {
	circuitDepartments := make(map[string]string)

	for _, circuit := range circuits {
		circuitDepartments[circuit.Id] = circuit.Department
	}

	return circuitDepartments
}

func (eu *ElectionUsecase) addCandidates(candidates []write.Candidate, electionId string) error {
	err := eu.electionRepository.AddCandiadtes(candidates, electionId)

	if err != nil {
		return err
	}

	allCandidates := eu.FillCandidateList(candidates)

	return eu.electionMemoryRepo.AddCandidates(allCandidates, electionId)
}

func (eu *ElectionUsecase) FillCandidateList(candidates []write.Candidate) []string {
	var allCandidates []string

	for _, candidate := range candidates {
		allCandidates = append(allCandidates, candidate.CandidateId)
	}

	return allCandidates
}

func (eu *ElectionUsecase) addVotersPerDepartment(electionId string, votersPerDepartment map[string]int) error {

	var departments []string

	for dept, voters := range votersPerDepartment {
		err := eu.electionMemoryRepo.AddVotersPerDepartment(read.VotersPerDepartment{
			ElectionId:   electionId,
			Department:   dept,
			VotersAmount: voters,
		})

		if err != nil {
			return err
		}

		departments = append(departments, dept)
	}

	return eu.electionMemoryRepo.AddDepartments(departments, electionId)
}

func (eu *ElectionUsecase) calculateVotersPerDepartment(completeElection write.CompleteElection, voterDepartments map[string]string) (map[string]int, map[string]map[int]int,
	map[string]map[string]int, map[string]map[string]int, map[string]map[int]int) {

	var log write.LoggingModel

	votersPerDepartment := make(map[string]int)
	votersPerAge := make(map[string]map[int]int)
	votersPerGender := make(map[string]map[string]int)
	votersPerDepartmentAge := make(map[string]map[int]int)
	votersPerDepartmentGender := make(map[string]map[string]int)

	for _, voter := range completeElection.Voters {

		birthDate, err := time.Parse("2006-01-02", voter.BirthDate)
		voterDepartment := voterDepartments[voter.CircuitId]

		if err != nil {
			log = write.LoggingModel{Type: "Error", Operation: "Calculate Voters Per Department", Actor: "Electoral Authority", Description: err.Error()}
			eu.helpers.LogHelper.SendLog(log)
		}

		voterAge := age.Age(birthDate)
		votersPerDepartment[voterDepartment] = votersPerDepartment[voterDepartment] + 1

		if votersPerAge[voter.CircuitId] == nil {
			votersPerAge[voter.CircuitId] = make(map[int]int)
			votersPerGender[voter.CircuitId] = make(map[string]int)
		}
		votersPerAge[voter.CircuitId][voterAge] = votersPerAge[voter.CircuitId][voterAge] + 1
		votersPerGender[voter.CircuitId][voter.Gender] = votersPerGender[voter.CircuitId][voter.Gender] + 1

		if votersPerDepartmentAge[voterDepartment] == nil {
			votersPerDepartmentAge[voterDepartment] = make(map[int]int)
			votersPerDepartmentGender[voterDepartment] = make(map[string]int)
		}
		votersPerDepartmentAge[voterDepartment][voterAge] = votersPerDepartmentAge[voterDepartment][voterAge] + 1
		votersPerDepartmentGender[voterDepartment][voter.Gender] = votersPerDepartmentGender[voterDepartment][voter.Gender] + 1
	}

	return votersPerDepartment, votersPerAge, votersPerGender, votersPerDepartmentGender, votersPerDepartmentAge
}

func (eu *ElectionUsecase) addPoliticalParties(political_parties []write.PoliticalParty, electionId string) error {
	err := eu.electionRepository.AddPoliticalParties(political_parties, electionId)

	if err != nil {
		return err
	}

	politicalParties := eu.FillPoliticalPartiesList(political_parties)

	return eu.electionMemoryRepo.AddPoliticalParties(politicalParties, electionId)
}

func (eu *ElectionUsecase) FillPoliticalPartiesList(political_parties []write.PoliticalParty) []string {
	var politicalParties []string

	for _, politicalParty := range political_parties {
		politicalParties = append(politicalParties, politicalParty.Name)
	}

	return politicalParties
}

func (eu *ElectionUsecase) addCircuits(circuits []write.Circuit) error {
	return eu.electionRepository.AddCircuits(circuits)
}

func (eu *ElectionUsecase) openCertificate(completeElection write.CompleteElection) error {
	certificate := eu.createOpenCertificate(completeElection)

	return eu.sendCertificate(certificate, "open-certificate-queue")
}

func (eu *ElectionUsecase) createOpenCertificate(completeElection write.CompleteElection) write.StartCertificate {
	return write.StartCertificate{
		StartDate:        completeElection.StartDate,
		PoliticalParties: completeElection.PoliticalParties,
		Candidates:       completeElection.Candidates,
		VotersAmmount:    len(completeElection.Voters),
		VotationMode:     completeElection.VotationMode,
	}
}

func (eu *ElectionUsecase) closeElection(c *cron.Cron, election write.CompleteElection) {
	var log write.LoggingModel

	if time.Now().UTC().After(election.EndDate.Add(time.Minute * 30)) {
		err := eu.endElectionPipeline.ExecuteElectionFilters(election)

		if err != nil {
			log = write.LoggingModel{Type: "Error", Operation: "Close Election", Actor: "Electoral Authority", Description: err.Error()}
			eu.helpers.LogHelper.SendLog(log)
		}

		closeCertificate := eu.createCloseCertificate(election)
		eu.sendCertificate(closeCertificate, "close-certificate-queue")
		c.Stop()
	}
}

func (eu *ElectionUsecase) createCloseCertificate(election write.CompleteElection) write.CloseCertificate {

	voteAmount, candidateVoters, partyVoters := eu.calculateElectionResult(election)
	return write.CloseCertificate{
		StartDate:      election.StartDate,
		EndDate:        election.EndDate,
		VoterAmount:    len(election.Voters),
		VoteAmount:     voteAmount,
		CandidateVotes: candidateVoters,
		PartyVotes:     partyVoters,
	}
}

func (eu *ElectionUsecase) sendCertificate(certificate interface{}, queue string) error {
	worker, err := rabbit.NewRabbitConnection()

	if err != nil {
		return err
	}

	jsonCertificate, err := json.Marshal(certificate)

	worker.Send(queue, []byte(string(jsonCertificate)))

	worker.Close()

	return err
}

func (eu *ElectionUsecase) calculateElectionResult(election write.CompleteElection) (int, []read.StringIntCandidate, []read.StringIntParty) {
	var log write.LoggingModel

	electionResult, err := eu.electionMemoryRepo.GetElectionResult(election.Id)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Close Election", Actor: "Electoral Authority", Description: err.Error()}
		eu.helpers.LogHelper.SendLog(log)
	}

	return electionResult.ElectionVotingData.TotalVoteAmount, electionResult.VotesPerCandidate, electionResult.VotesPerPoliticalParty
}
