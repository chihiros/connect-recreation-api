// Code generated by ent, DO NOT EDIT.

package ent

import (
	"app/ent/recreation"
	"app/ent/user"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Recreation is the model entity for the Recreation schema.
type Recreation struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// RecreationID holds the value of the "recreation_id" field.
	RecreationID uuid.UUID `json:"recreation_id,omitempty"`
	// Genre holds the value of the "genre" field.
	Genre []int `json:"genre,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// YoutubeID holds the value of the "youtube_id" field.
	YoutubeID string `json:"youtube_id,omitempty"`
	// TargetNumber holds the value of the "target_number" field.
	TargetNumber int `json:"target_number,omitempty"`
	// RequiredTime holds the value of the "required_time" field.
	RequiredTime int `json:"required_time,omitempty"`
	// Publish holds the value of the "publish" field.
	Publish bool `json:"publish,required"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// PublishedAt holds the value of the "published_at" field.
	PublishedAt time.Time `json:"published_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RecreationQuery when eager-loading is set.
	Edges            RecreationEdges `json:"edges"`
	user_recreations *int
	selectValues     sql.SelectValues
}

// RecreationEdges holds the relations/edges for other nodes in the graph.
type RecreationEdges struct {
	// Users holds the value of the users edge.
	Users *User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RecreationEdges) UsersOrErr() (*User, error) {
	if e.Users != nil {
		return e.Users, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Recreation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case recreation.FieldGenre:
			values[i] = new([]byte)
		case recreation.FieldPublish:
			values[i] = new(sql.NullBool)
		case recreation.FieldID, recreation.FieldTargetNumber, recreation.FieldRequiredTime:
			values[i] = new(sql.NullInt64)
		case recreation.FieldTitle, recreation.FieldContent, recreation.FieldYoutubeID:
			values[i] = new(sql.NullString)
		case recreation.FieldCreatedAt, recreation.FieldUpdatedAt, recreation.FieldPublishedAt:
			values[i] = new(sql.NullTime)
		case recreation.FieldUserID, recreation.FieldRecreationID:
			values[i] = new(uuid.UUID)
		case recreation.ForeignKeys[0]: // user_recreations
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Recreation fields.
func (r *Recreation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case recreation.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case recreation.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				r.UserID = *value
			}
		case recreation.FieldRecreationID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field recreation_id", values[i])
			} else if value != nil {
				r.RecreationID = *value
			}
		case recreation.FieldGenre:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field genre", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.Genre); err != nil {
					return fmt.Errorf("unmarshal field genre: %w", err)
				}
			}
		case recreation.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				r.Title = value.String
			}
		case recreation.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				r.Content = value.String
			}
		case recreation.FieldYoutubeID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field youtube_id", values[i])
			} else if value.Valid {
				r.YoutubeID = value.String
			}
		case recreation.FieldTargetNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field target_number", values[i])
			} else if value.Valid {
				r.TargetNumber = int(value.Int64)
			}
		case recreation.FieldRequiredTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field required_time", values[i])
			} else if value.Valid {
				r.RequiredTime = int(value.Int64)
			}
		case recreation.FieldPublish:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field publish", values[i])
			} else if value.Valid {
				r.Publish = value.Bool
			}
		case recreation.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case recreation.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		case recreation.FieldPublishedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field published_at", values[i])
			} else if value.Valid {
				r.PublishedAt = value.Time
			}
		case recreation.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_recreations", value)
			} else if value.Valid {
				r.user_recreations = new(int)
				*r.user_recreations = int(value.Int64)
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Recreation.
// This includes values selected through modifiers, order, etc.
func (r *Recreation) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryUsers queries the "users" edge of the Recreation entity.
func (r *Recreation) QueryUsers() *UserQuery {
	return NewRecreationClient(r.config).QueryUsers(r)
}

// Update returns a builder for updating this Recreation.
// Note that you need to call Recreation.Unwrap() before calling this method if this Recreation
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Recreation) Update() *RecreationUpdateOne {
	return NewRecreationClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Recreation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Recreation) Unwrap() *Recreation {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Recreation is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Recreation) String() string {
	var builder strings.Builder
	builder.WriteString("Recreation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", r.UserID))
	builder.WriteString(", ")
	builder.WriteString("recreation_id=")
	builder.WriteString(fmt.Sprintf("%v", r.RecreationID))
	builder.WriteString(", ")
	builder.WriteString("genre=")
	builder.WriteString(fmt.Sprintf("%v", r.Genre))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(r.Title)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(r.Content)
	builder.WriteString(", ")
	builder.WriteString("youtube_id=")
	builder.WriteString(r.YoutubeID)
	builder.WriteString(", ")
	builder.WriteString("target_number=")
	builder.WriteString(fmt.Sprintf("%v", r.TargetNumber))
	builder.WriteString(", ")
	builder.WriteString("required_time=")
	builder.WriteString(fmt.Sprintf("%v", r.RequiredTime))
	builder.WriteString(", ")
	builder.WriteString("publish=")
	builder.WriteString(fmt.Sprintf("%v", r.Publish))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("published_at=")
	builder.WriteString(r.PublishedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Recreations is a parsable slice of Recreation.
type Recreations []*Recreation
