package models

import (
	"database/sql"
	"time"

	database "github.com/Allexsen/booking-site/db"
)

// Used to model booking
type Booking struct {
	ID        int          `db:"id" json:"-"`
	StartDate sql.NullTime `db:"start_date" json:"start_date"`
	EndDate   sql.NullTime `db:"end_date" json:"end_date"`
	StatusID  int          `db:"status_id" json:"-"` // Booked(1), Ongoing(2), Cancelled(3), Archived(4)
	Status    string       `json:"status"`
}

func GetAllBookings() ([]Booking, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	q := `SELECT start_date, end_date
		FROM bookings
		WHERE status_id<3
		AND end_date>=?`

	rows, err := db.Query(q, time.Now())
	if err != nil {
		return nil, err
	}

	var bookings []Booking
	for rows.Next() {
		var nb Booking
		err := rows.Scan(&nb.StartDate, &nb.EndDate)
		if err != nil {
			return nil, err
		}

		bookings = append(bookings, nb)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return bookings, nil
}

func (b *Booking) Store() error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	q := `ISNERT INTO bookings(start_date, end_date, status_id)
		VALUES(?, ?, ?)`

	_, err = db.Exec(q, b.StartDate, b.EndDate, b.StatusID)
	return err
}

func (b *Booking) Activate() error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	q := `UPDATE bookings
		SET status_id=2
		WHERE id=?`
	_, err = db.Exec(q, b.ID)
	return err
}

func (b *Booking) Cancel() error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	q := `UPDATE bookings
		SET status_id="3"
		WHERE id=?`

	_, err = db.Exec(q, b.ID)
	return err
}

func (b *Booking) Archive() error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	q := `UPDATE bookings
		SET status_id=4
		WHERE id=?`

	_, err = db.Exec(q, b.ID)
	return err
}
