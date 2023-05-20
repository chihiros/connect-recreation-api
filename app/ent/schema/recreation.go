package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Recreation struct {
	ent.Schema
}

// Fields of the Recreation.
func (Recreation) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").
			Immutable().
			Unique(),
		field.String("uuid").
			Immutable().
			Unique(),
		field.JSON("genre", []int{}), // Use JSON field for 'genre'.
		field.String("title"),
		field.Int("target_number"),
		field.Int("requred_time"),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}
