package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Profile holds the schema definition for the Profile entity.
type Profile struct {
	ent.Schema
}

// Fields of the Profile.
func (Profile) Fields() []ent.Field {
	return []ent.Field{
		field.String("nickname").
			Unique(),
		// field.String("uuid").
		// 	Immutable().
		// 	Unique(),
		field.UUID("uuid", uuid.UUID{}).
			Immutable().
			Unique(),
		field.String("icon_url"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now),
	}
}
