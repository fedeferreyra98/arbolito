package model

import "time"

type CachedRates struct {
	Rates     map[string]Rate `bson:"rates"`
	CreatedAt time.Time       `bson:"created_at"`
}
