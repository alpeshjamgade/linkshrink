package models

import "time"

type Url struct {
	Id         string    `json:"id" db:"id"`
	Url        string    `json:"url" db:"url"`
	ShortUrl   string    `json:"short_url" db:"short_url"`
	InsertedAt time.Time `json:"inserted_at" db:"inserted_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}
