package models

import "time"

type Customer struct {
	ID          int       `db:"id"`
	User        User      `db:"user"`
	Name        string    `db:"name"`
	Email       string    `db:"email"`
	PhoneNumber string    `db:"phone_number"`
	Address     string    `db:"address"`
	Created_at  time.Time `db:"created_at"`
	Updated_at  time.Time `db:"updated_at"`
}
