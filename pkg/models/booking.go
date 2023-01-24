package models

import "time"

// Used to model booking
type Booking struct {
	ID        int
	StartDate time.Time
	EndDate   time.Time
}
