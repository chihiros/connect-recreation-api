package schema

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
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
		field.UUID("uuid", uuid.UUID{}).
			Immutable().
			Unique(),
		field.JSON("genre", &IntSlice{}), // Use JSON field for 'genre'.
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

type IntSlice []int

// Scan implements the sql.Scanner interface.
func (is *IntSlice) Scan(src interface{}) error {
	asBytes, ok := src.([]byte)
	if !ok {
		return errors.New("Scan source was not []bytes")
	}

	str := string(asBytes)

	// Remove leading and trailing brackets
	str = str[1 : len(str)-1]

	// Split on comma
	parts := strings.Split(str, ",")

	// Convert strings to integers
	result := make([]int, 0, len(parts))
	for _, p := range parts {
		var i int
		_, err := fmt.Sscanf(p, "%d", &i)
		if err != nil {
			return err
		}
		result = append(result, i)
	}

	*is = result
	return nil
}

// Value implements the driver.Valuer interface.
func (is IntSlice) Value() (driver.Value, error) {
	strs := make([]string, len(is))
	for i, v := range is {
		strs[i] = fmt.Sprint(v)
	}
	return "{" + strings.Join(strs, ",") + "}", nil
}
