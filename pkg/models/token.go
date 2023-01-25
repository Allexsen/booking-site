package models

import (
	"database/sql"
	"math/rand"
	"time"

	database "github.com/Allexsen/booking-site/db"
)

// Used to model tokens
type Token struct {
	ID         string        `db:"id"`
	PaymentID  int           `db:"payment_id"`
	BookingID  int           `db:"booking_id"`
	ReviewID   sql.NullInt32 `db:"review_id"`
	Email      string        `db:"email"`
	Reviewable bool          `db:"reviewable"`
	Refundable bool          `db:"refundable"`
}

func (t *Token) generateID() {
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Doesn't even need to check if it's duplicate
	// chance of having duplicate is
	// (number of existing tokens)/62^20
	// Basically impossible
	rand.Seed(time.Now().Unix())
	var token [20]byte
	for K := 0; K < 20; K++ {
		token[K] = charSet[rand.Intn(len(charSet))]
	}

	t.ID = string(token[:])
}

func (t *Token) Store() error {
	t.generateID()
	db := database.GetDB()
	q := `INSERT INTO tokens(id, payment_id, booking_id, email, refundable)
		VALUES(?, ?, ?, ?, ?)`

	_, err := db.Exec(q, t.ID, t.PaymentID, t.BookingID, t.Email, t.Refundable)
	return err
}

func (t *Token) DeclineRefund() error {
	db := database.GetDB()
	q := `UPDATE tokens
		SET refundable=false
		WHERE id=?`

	_, err := db.Exec(q, t.ID)
	return err
}

func (t *Token) EnableReview() error {
	db := database.GetDB()
	q := `UPDATE tokens
		SET reviewable=true
		WHERE id=?`

	_, err := db.Exec(q, t.ID)
	return err
}

func (t *Token) Deactivate() error {
	db := database.GetDB()
	q := `UPDATE tokens
		SET refundable=false, reviewable=false
		WHERE id=?`

	_, err := db.Exec(q, t.ID)
	return err
}
