package app

import "time"

type Product struct {
	Name      string    `bson:"name"`
	Price     float64   `bson:"price"`
	Updates   int       `bson:"updates"`
	FetchedAt time.Time `bson:"fetched_at"`
}
