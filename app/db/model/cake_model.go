package model

import "time"

type Cake struct {
	Id          int
	Title       string
	Description string
	Rating      float32
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
