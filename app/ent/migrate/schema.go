// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ProfilesColumns holds the columns for the "profiles" table.
	ProfilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeUUID, Unique: true},
		{Name: "nickname", Type: field.TypeString, Unique: true},
		{Name: "icon_url", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// ProfilesTable holds the schema information for the "profiles" table.
	ProfilesTable = &schema.Table{
		Name:       "profiles",
		Columns:    ProfilesColumns,
		PrimaryKey: []*schema.Column{ProfilesColumns[0]},
	}
	// RecreationsColumns holds the columns for the "recreations" table.
	RecreationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeUUID, Unique: true},
		{Name: "recreation_id", Type: field.TypeUUID, Unique: true},
		{Name: "genre", Type: field.TypeJSON},
		{Name: "title", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
		{Name: "youtube_id", Type: field.TypeString},
		{Name: "target_number", Type: field.TypeInt},
		{Name: "required_time", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// RecreationsTable holds the schema information for the "recreations" table.
	RecreationsTable = &schema.Table{
		Name:       "recreations",
		Columns:    RecreationsColumns,
		PrimaryKey: []*schema.Column{RecreationsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ProfilesTable,
		RecreationsTable,
	}
)

func init() {
}
