package models

// Used to model reviews
type Review struct {
	ID        int
	Rating    int
	Body      string
	Author    string
	BookingID int
}
