package model

import "time"

type CachedRate struct {
	Rate      Rate      `bson:"rate"`
	CreatedAt time.Time `bson:"created_at"`
}
