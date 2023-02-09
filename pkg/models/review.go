package models

import (
	"strconv"

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

func CreateReview(ratingstr, bidstr, body, author string) (int, error) {
	rating, err := strconv.Atoi(ratingstr)
	if err != nil {
		return -1, err
	}

	bid, err := strconv.Atoi(bidstr)
	if err != nil {
		return -1, err
	}

	r := Review{
		Rating:    rating,
		Body:      body,
		Author:    author,
		BookingID: bid,
	}

	err = r.store()
	if err != nil {
		return -1, err
	}

	return r.ID, nil
}

func (r *Review) store() error {
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
