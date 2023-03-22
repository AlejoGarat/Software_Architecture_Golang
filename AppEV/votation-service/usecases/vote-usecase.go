package usecases

import (
	"math/rand"
	"strconv"

	idataaccess "votation-service/dataaccess/interfaces"
	"votation-service/helpers"
	imessenger "votation-service/messenger/interfaces"
	"votation-service/models/read"
	"votation-service/models/write"

	"github.com/robfig/cron/v3"
)

const (
	randomSeed = 1000000
)

type VoteUseCase struct {
	voteRepository idataaccess.VoteRepository
	messenger      imessenger.Messenger
	helpers        helpers.Helpers
}

func NewVoteUseCase(voteRepository idataaccess.VoteRepository, messenger imessenger.Messenger, helpers helpers.Helpers) *VoteUseCase {
	return &VoteUseCase{voteRepository: voteRepository, messenger: messenger, helpers: helpers}
}

func (voteUseCase *VoteUseCase) AddVote(vote write.Vote) {
	go voteUseCase.voteBalancer(vote)
}

func (voteUseCase *VoteUseCase) SendMailConstancy(constancyReq read.ConstancyRequest) error {
	var log write.LoggingModel

	constancyData, err := voteUseCase.voteRepository.GetConstancyData(constancyReq)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Mail Constancy", Actor: "Voter", Description: err.Error()}
		voteUseCase.helpers.LogHelper.SendLog(log)
		return err
	}

	go voteUseCase.mailConstancyBalancer(constancyData)

	return nil
}

func (voteUseCase *VoteUseCase) UpdateConstancyDBData(constancyDBData write.ConstancyDBData, electionId string) error {
	return voteUseCase.voteRepository.SetConstancyData(constancyDBData, electionId)
}

func (voteUseCase *VoteUseCase) voteBalancer(vote write.Vote) {
	var log write.LoggingModel
	vote.VoteId = createVoteId()

	go func() {
		constancyCron := cron.New()
		err := voteUseCase.generateSMSConstancy(vote, constancyCron)

		if err != nil {
			log = write.LoggingModel{Type: "Error", Operation: "Add Vote", Actor: "Voter", Description: err.Error()}
			voteUseCase.helpers.LogHelper.SendLog(log)
		}
	}()

	voteCron := cron.New()
	err := voteUseCase.voteRepository.AddVote(vote)

	if err != nil {
		voteCron.AddFunc("@every 3m", func() {
			log = write.LoggingModel{Type: "Error", Operation: "Add Vote", Actor: "Voter", Description: err.Error()}
			voteUseCase.helpers.LogHelper.SendLog(log)

			err = voteUseCase.voteRepository.AddVote(vote)

			if err == nil {
				voteCron.Stop()
			}
		})

		voteCron.Start()
	}

	err = voteUseCase.helpers.DataUpdateHelper.SendVote(vote)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Add Vote", Actor: "Voter", Description: err.Error()}
		voteUseCase.helpers.LogHelper.SendLog(log)
	}
}

func (voteUseCase *VoteUseCase) mailConstancyBalancer(constancyData write.Constancy) {
	var log write.LoggingModel
	err := voteUseCase.helpers.MailHelper.SendMail(constancyData)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Mail Constancy", Actor: "Voter", Description: err.Error()}
		voteUseCase.helpers.LogHelper.SendLog(log)
	}
}

func (voteUseCase *VoteUseCase) generateSMSConstancy(vote write.Vote, constancyCron *cron.Cron) error {

	electionData, err := voteUseCase.voteRepository.GetElectionById(vote.ElectionId)

	if err != nil {
		return err
	}

	voterData, err := voteUseCase.voteRepository.GetVoterById(vote.VoterDocument, vote.ElectionId)

	if err != nil {
		return err
	}

	constancy := write.Constancy{
		Election:             electionData,
		VoteEmissionInfoDate: vote.Info,
		VoterDocument:        vote.VoterDocument,
		Name:                 voterData.Name,
		Surname:              voterData.Surname,
		VoteId:               vote.VoteId,
	}

	err = voteUseCase.messenger.SendMessage(constancy)

	if err != nil {
		return err
	}

	constancyCron.Stop()

	return nil
}

func createVoteId() string {
	val := rand.Intn(randomSeed)
	return strconv.Itoa(val)
}
