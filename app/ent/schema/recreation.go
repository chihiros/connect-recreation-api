package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Recreation struct {
	ent.Schema
}

// Fields of the Recreation.
func (Recreation) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}).
			Immutable().
			Unique(),
		field.UUID("recreation_id", uuid.UUID{}).
			Immutable().
			Unique(),
		field.JSON("genre", []int{}),
		field.String("title"),
		field.String("content"),
		field.Int("target_number"),
		field.Int("required_time"),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}
