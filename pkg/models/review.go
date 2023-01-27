package models

import (
	database "github.com/Allexsen/booking-site/db"
)

// Used to model reviews
type Review struct {
	ID        int    `db:"id" json:"-"`
	Rating    int    `db:"rating" json:"stars"`
	Body      string `db:"body" json:"review"`
	Author    string `db:"author" json:"author"`
	BookingID int    `db:"booking_id" json:"-"`
}

func GetReviews() ([]Review, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	q := `SELECT rating, body, author
		FROM reviews`

	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}

	var reviews []Review
	for rows.Next() {
		var r Review
		err := rows.Scan(&r.Rating, &r.Body, &r.Author)
		if err != nil {
			return nil, err
		}

		reviews = append(reviews, r)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *Review) Store() error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	q := `INSERT INTO reviews(rating, body, author, booking_id)
		VALUES(?, ?, ?, ?)`

	_, err = db.Exec(q, r.Rating, r.Body, r.Author, r.BookingID)
	return err
}
