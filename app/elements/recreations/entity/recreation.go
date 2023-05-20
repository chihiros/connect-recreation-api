package entity

import "time"

type Recreation struct {
	ID        int       `json:"id"`
	UUID      string    `json:"uid"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Target    Target    `json:"target"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Target struct {
	Age    int `json:"age"`
	Nunmer int `json:"number"`
}
