package entity

import "time"

type User struct {
	CreatedAt time.Time `json:"created_at"`
}
