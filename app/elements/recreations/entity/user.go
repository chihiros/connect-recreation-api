package entity

import "time"

type User struct {
	ID           int       `json:"id"`
	UID          string    `json:"uid"`
	Username     string    `json:"username"`
	Mail         string    `json:"mail"`
	PrefectureID int       `json:"prefecture_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
