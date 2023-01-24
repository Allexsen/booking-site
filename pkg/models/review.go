package models

// Used to model reviews
type Review struct {
	ID          int
	Rating      int
	Description string
	Author      string
	BookingID   int
}
