package models

import "time"

type Customer struct {
	ID          int `gorm:"column:id"`
	User        uint
	Name        string
	Email       string
	PhoneNumber string
	Address     string
	Created_at  time.Time
	Updated_at  time.Time
}
