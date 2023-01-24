package models

// Used to model tokens
type Token struct {
	ID        int
	PaymentID int
	BookingID int
	ReviewID  int
	CanReview bool
	CanRefund bool
}
