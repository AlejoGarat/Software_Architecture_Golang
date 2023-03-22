package models

import "time"

type Constancy struct {
	Timestamps []time.Time `bson:"timestamps"`
}
