package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("uid").
			Immutable().
			Unique(),
		field.String("username").
			Unique(),
		field.String("mail").
			Unique(),
		field.Int("prefecture_id").
			Nillable(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}
