package usecase

import "github.com/google/uuid"

type Request struct {
	UUID     uuid.UUID `json:"uuid"`
	Nickname string    `json:"nickname"`
	IconURL  string    `json:"icon_url"`
}
