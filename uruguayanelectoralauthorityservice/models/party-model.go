package models

type Id = string
type Name = string

type Party struct {
	Id   Id   `json:"id_party"`
	Name Name `json:"name"`
}
