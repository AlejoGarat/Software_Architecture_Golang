package models

type IdCircuit = string
type Department = string
type Location = string

type Circuit struct {
	IdCircuit  IdCircuit  `json:"id_circuit"`
	Department Department `json:"department"`
	Location   Location   `json:"location" `
}
