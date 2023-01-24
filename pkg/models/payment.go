package models

import "time"

// Used to model payments
type Payment struct {
	ID        int
	BookingID int
	Method    string
	Amount    int
	Date      time.Time
}
