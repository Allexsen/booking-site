package models

import database "github.com/Allexsen/booking-site/db"

// Used to model reviews
type Review struct {
	ID        int    `db:"id" json:"-"`
	Rating    int    `db:"rating" json:"stars"`
	Body      string `db:"body" json:"review"`
	Author    string `db:"author" json:"author"`
	BookingID int    `db:"booking_id" json:"-"`
}

func (r *Review) Store() error {
	db := database.GetDB()
	q := `INSERT INTO reviews(rating, body, author, booking_id)
		VALUES(?, ?, ?, ?)`

	_, err := db.Exec(q, r.Rating, r.Body, r.Author, r.BookingID)
	return err
}
