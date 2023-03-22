package models

type Info struct {
	Date string `bson:"date"`
	Hour string `bson:"hour"`
}

type Vote struct {
	CompleteInfo []Info `bson:"info"`
}
