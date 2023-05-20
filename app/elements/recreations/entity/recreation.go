package entity

import (
	"time"

	"github.com/google/uuid"
)

type Recreation struct {
	ID           int       `json:"id"`
	UserID       uuid.UUID `json:"user_id"` // レクリエーションを登録したユーザーのID
	UUID         uuid.UUID `json:"uid"`     // レクリエーションのユニークID
	Title        string    `json:"title"`   // レクリエーションの名前
	Content      string    `json:"content"` // レクリエーションの説明
	TargetNumber int       `json:"target"`  // レクリエーションの対象人数
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
