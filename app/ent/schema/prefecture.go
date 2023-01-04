package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Prefecture struct {
	ent.Schema
}

// Fields of the Prefecture.
func (Prefecture) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}
