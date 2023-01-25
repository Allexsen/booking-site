package models

import (
	"database/sql"

	database "github.com/Allexsen/booking-site/db"
)

// Used to model booking
type Booking struct {
	ID        int          `db:"id" json:"id"`
	StartDate sql.NullTime `db:"start_date" json:"start_date"`
	EndDate   sql.NullTime `db:"end_date" json:"end_date"`
	StatusID  int          `db:"status_id"` // Booked(1), Ongoing(2), Cancelled(3), Archived(4)
	Status    string       `json:"status"`
}

func (b *Booking) Store() error {
	db := database.GetDB()
	q := `ISNERT INTO bookings(start_date, end_date, status_id)
		VALUES(?, ?, ?)`

	_, err := db.Exec(q, b.StartDate, b.EndDate, b.StatusID)
	return err
}

func (b *Booking) Activate() error {
	db := database.GetDB()
	q := `UPDATE bookings
		SET status_id=2
		WHERE id=?`
	_, err := db.Exec(q, b.ID)
	return err
}

func (b *Booking) Cancel() error {
	db := database.GetDB()
	q := `UPDATE bookings
		SET status_id="3"
		WHERE id=?`

	_, err := db.Exec(q, b.ID)
	return err
}

func (b *Booking) Archive() error {
	db := database.GetDB()
	q := `UPDATE bookings
		SET status_id=4
		WHERE id=?`

	_, err := db.Exec(q, b.ID)
	return err
}
