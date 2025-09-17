package entity

import "time"

type transaction struct {
	ID         int
	FROMWALLET int
	TOWALLET   int
	AMOUNT     float64
	CREATEDAT  time.Time
}
