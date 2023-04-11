package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
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
		field.String("uuid").
			Unique(),
		field.String("icon_url"),
	}
}
