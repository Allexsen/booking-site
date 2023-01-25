package models

import (
	"time"

	database "github.com/Allexsen/booking-site/db"
)

// Used to model payments
type Payment struct {
	ID              int       `db:"id"`
	BookingID       int       `db:"booking_id"`
	Method          string    `db:"method"`
	Amount          int       `db:"amount"`
	TransactionTime time.Time `db:"transaction_time"`
	Refunded        bool      `db:"refunded"`
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
