package controllers

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	fiber "github.com/gofiber/fiber/v2"
)

type DataController struct{}

type Circuit struct {
	Id         string `json:"id" fake:"{regex:[123456789]{3}}"`
	ElectionId string `json:"election_id"`
	Department string `json:"department" fake:"{randomstring:[Montevideo, Maldonado, Canelones, Colonia, Florida, Artigas, Paysandú, Salto, Flores]}"`
	Location   string `json:"location" fake:"{randomstring:[Ubicacion A, Ubicacion B, Ubicacion C]}"`
}

type PoliticalParty struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	ElectionId string `json:"election_id"`
}

type Candidate struct {
	Id               string `json:"id" fake:"{regex:[0123456789]{8}}"`
	Name             string `json:"name" fake:"{firstname}"`
	Surname          string `json:"surname" fake:"{firstname}"`
	Gender           string `json:"gender" fake:"{randomstring:[M,F]}"`
	BirthDate        string `json:"birth_date" fake:"{year}-0{regex:[123456789]}-0{regex:[123456789]}" format:"2006-01-02"`
	PoliticalPartyId string `json:"political_party_id" fake:"{number:1,6}"`
	ElectionId       string `json:"election_id"`
}

type Voter struct {
	Id                  string `json:"id" fake:"{regex:[0123456789]{8}}"`
	Credential          string `json:"credential" fake:"{regex:[AXBC]{3}}"`
	Name                string `json:"name" fake:"{firstname}"`
	Surname             string `json:"surname" fake:"{firstname}"`
	Gender              string `json:"gender" fake:"{randomstring:[M,F]}"`
	BirthDate           string `json:"birth_date" fake:"{year}-0{regex:[123456789]}-0{regex:[123456789]}" format:"2006-01-02"`
	ResidenceDepartment string `json:"residence_department" fake:"{randomstring:[Montevideo, Maldonado, Canelones, Colonia, Florida, Artigas, Paysandú, Salto, Flores]}"`
	CircuitId           string `json:"circuit_id"`
	Celphone            string `json:"cellphone" fake:"{regex:[123456789]{3}}"`
	Mail                string `json:"mail" fake:"{regex:[123456789]{3}}"`
}

type ElectoralAuthorityData struct {
	Id               string           `json:"id"`
	Description      string           `json:"description"`
	StartDate        time.Time        `json:"start_date"`
	EndDate          time.Time        `json:"end_date"`
	Circuits         []Circuit        `json:"circuits"`
	PoliticalParties []PoliticalParty `json:"political_parties"`
	Candidates       []Candidate      `json:"candidates"`
	Voters           []Voter          `json:"voters"`
	VotationMode     string           `json:"votation_mode"`
}

func NewDataController() *DataController {
	return &DataController{}
}

func (controller *DataController) GetData(c *fiber.Ctx) error {

	electionId := "1"

	var candidates []Candidate

	var voters []Voter

	var circuits []Circuit

	var parties []PoliticalParty

	partidoNacional := "Partido Nacional"
	id1 := "1"
	pnStruct := PoliticalParty{id1, partidoNacional, "1"}
	parties = append(parties, pnStruct)

	partidoColorado := "Partido Colorado"
	id2 := "2"
	pcStruct := PoliticalParty{id2, partidoColorado, "1"}
	parties = append(parties, pcStruct)

	cabildoAbierto := "Cabildo Abierto"
	id3 := "3"
	caStruct := PoliticalParty{id3, cabildoAbierto, "1"}
	parties = append(parties, caStruct)

	partidoIndependiente := "Partido Independiente"
	id4 := "4"
	piStruct := PoliticalParty{id4, partidoIndependiente, "1"}
	parties = append(parties, piStruct)

	frenteAmplio := "Frente Amplio"
	id5 := "5"
	faStruct := PoliticalParty{id5, frenteAmplio, "1"}
	parties = append(parties, faStruct)

	partidoRadical := "Partido Radical"
	id6 := "6"
	prStruct := PoliticalParty{id6, partidoRadical, "1"}
	parties = append(parties, prStruct)

	partidoVerde := "Partido Verde Minimalista"
	id7 := "7"
	pvStruct := PoliticalParty{id7, partidoVerde, "1"}
	parties = append(parties, pvStruct)
	var circuitIds [11]string

	for i := 1; i <= 10; i++ {
		var circuit Circuit
		gofakeit.Struct(&circuit)
		circuit.ElectionId = electionId
		circuits = append(circuits, circuit)
		circuitIds[i] = circuit.Id
	}

	for i := 1; i <= 10; i++ {
		var candidate Candidate
		gofakeit.Struct(&candidate)
		candidate.PoliticalPartyId = strconv.Itoa((i % len(parties)) + 1)
		candidate.ElectionId = electionId
		candidates = append(candidates, candidate)
	}

	for i := 1; i <= 10; i++ {
		var voter Voter
		gofakeit.Struct(&voter)
		voter.Id = "510890" + strconv.Itoa(i)
		voter.CircuitId = circuitIds[i]
		voter.Celphone = gofakeit.Phone()
		voter.Mail = gofakeit.Email()
		voters = append(voters, voter)
	}

	electoralAuthorityData := &ElectoralAuthorityData{
		Id:               "1",
		Description:      "Elección presidente 2022",
		StartDate:        time.Now(),
		EndDate:          time.Now().Add(time.Hour * 100),
		Circuits:         circuits,
		PoliticalParties: parties,
		Candidates:       candidates,
		Voters:           voters,
		VotationMode:     "unique",
	}

	jsonData, _ := json.Marshal(electoralAuthorityData)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Status Code":            200,
		"ElectoralAuthorityData": string(jsonData),
	})
}
