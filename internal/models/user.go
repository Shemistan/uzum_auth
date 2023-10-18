package models

import "time"

type User struct {
	Login      string
	Password   string
	Role       string
	DateOfBirt time.Time
}
