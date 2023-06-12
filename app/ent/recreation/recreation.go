// Code generated by ent, DO NOT EDIT.

package recreation

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the recreation type in the database.
	Label = "recreation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldRecreationID holds the string denoting the recreation_id field in the database.
	FieldRecreationID = "recreation_id"
	// FieldGenre holds the string denoting the genre field in the database.
	FieldGenre = "genre"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldYoutubeID holds the string denoting the youtube_id field in the database.
	FieldYoutubeID = "youtube_id"
	// FieldTargetNumber holds the string denoting the target_number field in the database.
	FieldTargetNumber = "target_number"
	// FieldRequiredTime holds the string denoting the required_time field in the database.
	FieldRequiredTime = "required_time"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeProfile holds the string denoting the profile edge name in mutations.
	EdgeProfile = "profile"
	// Table holds the table name of the recreation in the database.
	Table = "recreations"
	// ProfileTable is the table that holds the profile relation/edge.
	ProfileTable = "recreations"
	// ProfileInverseTable is the table name for the Profile entity.
	// It exists in this package in order to avoid circular dependency with the "profile" package.
	ProfileInverseTable = "profiles"
	// ProfileColumn is the table column denoting the profile relation/edge.
	ProfileColumn = "profile_recreations"
)

// Columns holds all SQL columns for recreation fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldRecreationID,
	FieldGenre,
	FieldTitle,
	FieldContent,
	FieldYoutubeID,
	FieldTargetNumber,
	FieldRequiredTime,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "recreations"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"profile_recreations",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the Recreation queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByRecreationID orders the results by the recreation_id field.
func ByRecreationID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRecreationID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}

// ByYoutubeID orders the results by the youtube_id field.
func ByYoutubeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldYoutubeID, opts...).ToFunc()
}

// ByTargetNumber orders the results by the target_number field.
func ByTargetNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTargetNumber, opts...).ToFunc()
}

// ByRequiredTime orders the results by the required_time field.
func ByRequiredTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRequiredTime, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByProfileField orders the results by profile field.
func ByProfileField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProfileStep(), sql.OrderByField(field, opts...))
	}
}
func newProfileStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProfileInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProfileTable, ProfileColumn),
	)
}
