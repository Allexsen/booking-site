package models

import (
	"database/sql"
	"math/rand"
	"strconv"
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

func CreateToken(pidstr, bidstr, email string) (string, error) {
	pid, err := strconv.Atoi(pidstr)
	if err != nil {
		return "", err
	}

	bid, err := strconv.Atoi(bidstr)
	if err != nil {
		return "", err
	}

	token := Token{
		PaymentID:  pid,
		BookingID:  bid,
		Email:      email,
		Refundable: true,
	}

	err = token.store()
	if err != nil {
		return "", err
	}

	return token.ID, err
}

func CheckToken(token string) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	q := `SELECT id FROM tokens WHERE id=?`
	_, err = db.Query(q, token)
	if err != nil {
		return err
	}

	return nil
}

func DisableToken(token string) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	q := `UPDATE tokens
	SET refundable=false,
	reviewable=false,
	WHERE id=?`

	_, err = db.Exec(q, token)
	return err
}

func (t *Token) store() error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	t.generateID()
	q := `INSERT INTO tokens(id, payment_id, booking_id, email, refundable)
	VALUES(?, ?, ?, ?, ?)`

	_, err = db.Exec(q, t.ID, t.PaymentID, t.BookingID, t.Email, t.Refundable)
	return err
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
