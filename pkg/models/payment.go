package models

import (
	"time"

	database "github.com/Allexsen/booking-site/db"
)

// Used to model payments
type Payment struct {
	ID              int
	BookingID       int
	Method          string
	Amount          int
	TransactionTime time.Time
	Refunded        bool
}

func (p *Payment) Store() error {
	db := database.GetDB()
	q := `INSERT INTO payments(booking_id, method, amount, transaction_time)
		VALUES(?, ?, ?, ?)`

	_, err := db.Exec(q, p.BookingID, p.Method, p.Amount, p.TransactionTime)
	return err
}

func (p *Payment) Refund() error {
	db := database.GetDB()
	q := `UPDATE payments
		SET refunded=true
		WHERE id=?`

	_, err := db.Exec(q, p.ID)
	return err
}
