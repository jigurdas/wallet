package entity

import "time"

type Transaction struct {
	ID         int64
	FROMWALLET int
	TOWALLET   int
	AMOUNT     int64
	CREATEDAT  time.Time
}
