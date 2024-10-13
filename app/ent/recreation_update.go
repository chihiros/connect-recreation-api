// Code generated by ent, DO NOT EDIT.

package ent

import (
	"app/ent/predicate"
	"app/ent/recreation"
	"app/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// RecreationUpdate is the builder for updating Recreation entities.
type RecreationUpdate struct {
	config
	hooks     []Hook
	mutation  *RecreationMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the RecreationUpdate builder.
func (ru *RecreationUpdate) Where(ps ...predicate.Recreation) *RecreationUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetGenre sets the "genre" field.
func (ru *RecreationUpdate) SetGenre(i []int) *RecreationUpdate {
	ru.mutation.SetGenre(i)
	return ru
}

// AppendGenre appends i to the "genre" field.
func (ru *RecreationUpdate) AppendGenre(i []int) *RecreationUpdate {
	ru.mutation.AppendGenre(i)
	return ru
}

// SetTitle sets the "title" field.
func (ru *RecreationUpdate) SetTitle(s string) *RecreationUpdate {
	ru.mutation.SetTitle(s)
	return ru
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (ru *RecreationUpdate) SetNillableTitle(s *string) *RecreationUpdate {
	if s != nil {
		ru.SetTitle(*s)
	}
	return ru
}

// SetContent sets the "content" field.
func (ru *RecreationUpdate) SetContent(s string) *RecreationUpdate {
	ru.mutation.SetContent(s)
	return ru
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (ru *RecreationUpdate) SetNillableContent(s *string) *RecreationUpdate {
	if s != nil {
		ru.SetContent(*s)
	}
	return ru
}

// SetYoutubeID sets the "youtube_id" field.
func (ru *RecreationUpdate) SetYoutubeID(s string) *RecreationUpdate {
	ru.mutation.SetYoutubeID(s)
	return ru
}

// SetNillableYoutubeID sets the "youtube_id" field if the given value is not nil.
func (ru *RecreationUpdate) SetNillableYoutubeID(s *string) *RecreationUpdate {
	if s != nil {
		ru.SetYoutubeID(*s)
	}
	return ru
}

// ClearYoutubeID clears the value of the "youtube_id" field.
func (ru *RecreationUpdate) ClearYoutubeID() *RecreationUpdate {
	ru.mutation.ClearYoutubeID()
	return ru
}

// SetTargetNumber sets the "target_number" field.
func (ru *RecreationUpdate) SetTargetNumber(i int) *RecreationUpdate {
	ru.mutation.ResetTargetNumber()
	ru.mutation.SetTargetNumber(i)
	return ru
}

// SetNillableTargetNumber sets the "target_number" field if the given value is not nil.
func (ru *RecreationUpdate) SetNillableTargetNumber(i *int) *RecreationUpdate {
	if i != nil {
		ru.SetTargetNumber(*i)
	}
	return ru
}

// AddTargetNumber adds i to the "target_number" field.
func (ru *RecreationUpdate) AddTargetNumber(i int) *RecreationUpdate {
	ru.mutation.AddTargetNumber(i)
	return ru
}

// SetRequiredTime sets the "required_time" field.
func (ru *RecreationUpdate) SetRequiredTime(i int) *RecreationUpdate {
	ru.mutation.ResetRequiredTime()
	ru.mutation.SetRequiredTime(i)
	return ru
}

// SetNillableRequiredTime sets the "required_time" field if the given value is not nil.
func (ru *RecreationUpdate) SetNillableRequiredTime(i *int) *RecreationUpdate {
	if i != nil {
		ru.SetRequiredTime(*i)
	}
	return ru
}

// AddRequiredTime adds i to the "required_time" field.
func (ru *RecreationUpdate) AddRequiredTime(i int) *RecreationUpdate {
	ru.mutation.AddRequiredTime(i)
	return ru
}

// SetPublish sets the "publish" field.
func (ru *RecreationUpdate) SetPublish(b bool) *RecreationUpdate {
	ru.mutation.SetPublish(b)
	return ru
}

// SetNillablePublish sets the "publish" field if the given value is not nil.
func (ru *RecreationUpdate) SetNillablePublish(b *bool) *RecreationUpdate {
	if b != nil {
		ru.SetPublish(*b)
	}
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *RecreationUpdate) SetUpdatedAt(t time.Time) *RecreationUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetPublishedAt sets the "published_at" field.
func (ru *RecreationUpdate) SetPublishedAt(t time.Time) *RecreationUpdate {
	ru.mutation.SetPublishedAt(t)
	return ru
}

// SetNillablePublishedAt sets the "published_at" field if the given value is not nil.
func (ru *RecreationUpdate) SetNillablePublishedAt(t *time.Time) *RecreationUpdate {
	if t != nil {
		ru.SetPublishedAt(*t)
	}
	return ru
}

// ClearPublishedAt clears the value of the "published_at" field.
func (ru *RecreationUpdate) ClearPublishedAt() *RecreationUpdate {
	ru.mutation.ClearPublishedAt()
	return ru
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (ru *RecreationUpdate) SetUsersID(id int) *RecreationUpdate {
	ru.mutation.SetUsersID(id)
	return ru
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (ru *RecreationUpdate) SetNillableUsersID(id *int) *RecreationUpdate {
	if id != nil {
		ru = ru.SetUsersID(*id)
	}
	return ru
}

// SetUsers sets the "users" edge to the User entity.
func (ru *RecreationUpdate) SetUsers(u *User) *RecreationUpdate {
	return ru.SetUsersID(u.ID)
}

// Mutation returns the RecreationMutation object of the builder.
func (ru *RecreationUpdate) Mutation() *RecreationMutation {
	return ru.mutation
}

// ClearUsers clears the "users" edge to the User entity.
func (ru *RecreationUpdate) ClearUsers() *RecreationUpdate {
	ru.mutation.ClearUsers()
	return ru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RecreationUpdate) Save(ctx context.Context) (int, error) {
	ru.defaults()
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RecreationUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RecreationUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RecreationUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *RecreationUpdate) defaults() {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		v := recreation.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ru *RecreationUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *RecreationUpdate {
	ru.modifiers = append(ru.modifiers, modifiers...)
	return ru
}

func (ru *RecreationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(recreation.Table, recreation.Columns, sqlgraph.NewFieldSpec(recreation.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Genre(); ok {
		_spec.SetField(recreation.FieldGenre, field.TypeJSON, value)
	}
	if value, ok := ru.mutation.AppendedGenre(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, recreation.FieldGenre, value)
		})
	}
	if value, ok := ru.mutation.Title(); ok {
		_spec.SetField(recreation.FieldTitle, field.TypeString, value)
	}
	if value, ok := ru.mutation.Content(); ok {
		_spec.SetField(recreation.FieldContent, field.TypeString, value)
	}
	if value, ok := ru.mutation.YoutubeID(); ok {
		_spec.SetField(recreation.FieldYoutubeID, field.TypeString, value)
	}
	if ru.mutation.YoutubeIDCleared() {
		_spec.ClearField(recreation.FieldYoutubeID, field.TypeString)
	}
	if value, ok := ru.mutation.TargetNumber(); ok {
		_spec.SetField(recreation.FieldTargetNumber, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedTargetNumber(); ok {
		_spec.AddField(recreation.FieldTargetNumber, field.TypeInt, value)
	}
	if value, ok := ru.mutation.RequiredTime(); ok {
		_spec.SetField(recreation.FieldRequiredTime, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedRequiredTime(); ok {
		_spec.AddField(recreation.FieldRequiredTime, field.TypeInt, value)
	}
	if value, ok := ru.mutation.Publish(); ok {
		_spec.SetField(recreation.FieldPublish, field.TypeBool, value)
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.SetField(recreation.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ru.mutation.PublishedAt(); ok {
		_spec.SetField(recreation.FieldPublishedAt, field.TypeTime, value)
	}
	if ru.mutation.PublishedAtCleared() {
		_spec.ClearField(recreation.FieldPublishedAt, field.TypeTime)
	}
	if ru.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recreation.UsersTable,
			Columns: []string{recreation.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recreation.UsersTable,
			Columns: []string{recreation.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(ru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{recreation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RecreationUpdateOne is the builder for updating a single Recreation entity.
type RecreationUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *RecreationMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetGenre sets the "genre" field.
func (ruo *RecreationUpdateOne) SetGenre(i []int) *RecreationUpdateOne {
	ruo.mutation.SetGenre(i)
	return ruo
}

// AppendGenre appends i to the "genre" field.
func (ruo *RecreationUpdateOne) AppendGenre(i []int) *RecreationUpdateOne {
	ruo.mutation.AppendGenre(i)
	return ruo
}

// SetTitle sets the "title" field.
func (ruo *RecreationUpdateOne) SetTitle(s string) *RecreationUpdateOne {
	ruo.mutation.SetTitle(s)
	return ruo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (ruo *RecreationUpdateOne) SetNillableTitle(s *string) *RecreationUpdateOne {
	if s != nil {
		ruo.SetTitle(*s)
	}
	return ruo
}

// SetContent sets the "content" field.
func (ruo *RecreationUpdateOne) SetContent(s string) *RecreationUpdateOne {
	ruo.mutation.SetContent(s)
	return ruo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (ruo *RecreationUpdateOne) SetNillableContent(s *string) *RecreationUpdateOne {
	if s != nil {
		ruo.SetContent(*s)
	}
	return ruo
}

// SetYoutubeID sets the "youtube_id" field.
func (ruo *RecreationUpdateOne) SetYoutubeID(s string) *RecreationUpdateOne {
	ruo.mutation.SetYoutubeID(s)
	return ruo
}

// SetNillableYoutubeID sets the "youtube_id" field if the given value is not nil.
func (ruo *RecreationUpdateOne) SetNillableYoutubeID(s *string) *RecreationUpdateOne {
	if s != nil {
		ruo.SetYoutubeID(*s)
	}
	return ruo
}

// ClearYoutubeID clears the value of the "youtube_id" field.
func (ruo *RecreationUpdateOne) ClearYoutubeID() *RecreationUpdateOne {
	ruo.mutation.ClearYoutubeID()
	return ruo
}

// SetTargetNumber sets the "target_number" field.
func (ruo *RecreationUpdateOne) SetTargetNumber(i int) *RecreationUpdateOne {
	ruo.mutation.ResetTargetNumber()
	ruo.mutation.SetTargetNumber(i)
	return ruo
}

// SetNillableTargetNumber sets the "target_number" field if the given value is not nil.
func (ruo *RecreationUpdateOne) SetNillableTargetNumber(i *int) *RecreationUpdateOne {
	if i != nil {
		ruo.SetTargetNumber(*i)
	}
	return ruo
}

// AddTargetNumber adds i to the "target_number" field.
func (ruo *RecreationUpdateOne) AddTargetNumber(i int) *RecreationUpdateOne {
	ruo.mutation.AddTargetNumber(i)
	return ruo
}

// SetRequiredTime sets the "required_time" field.
func (ruo *RecreationUpdateOne) SetRequiredTime(i int) *RecreationUpdateOne {
	ruo.mutation.ResetRequiredTime()
	ruo.mutation.SetRequiredTime(i)
	return ruo
}

// SetNillableRequiredTime sets the "required_time" field if the given value is not nil.
func (ruo *RecreationUpdateOne) SetNillableRequiredTime(i *int) *RecreationUpdateOne {
	if i != nil {
		ruo.SetRequiredTime(*i)
	}
	return ruo
}

// AddRequiredTime adds i to the "required_time" field.
func (ruo *RecreationUpdateOne) AddRequiredTime(i int) *RecreationUpdateOne {
	ruo.mutation.AddRequiredTime(i)
	return ruo
}

// SetPublish sets the "publish" field.
func (ruo *RecreationUpdateOne) SetPublish(b bool) *RecreationUpdateOne {
	ruo.mutation.SetPublish(b)
	return ruo
}

// SetNillablePublish sets the "publish" field if the given value is not nil.
func (ruo *RecreationUpdateOne) SetNillablePublish(b *bool) *RecreationUpdateOne {
	if b != nil {
		ruo.SetPublish(*b)
	}
	return ruo
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *RecreationUpdateOne) SetUpdatedAt(t time.Time) *RecreationUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetPublishedAt sets the "published_at" field.
func (ruo *RecreationUpdateOne) SetPublishedAt(t time.Time) *RecreationUpdateOne {
	ruo.mutation.SetPublishedAt(t)
	return ruo
}

// SetNillablePublishedAt sets the "published_at" field if the given value is not nil.
func (ruo *RecreationUpdateOne) SetNillablePublishedAt(t *time.Time) *RecreationUpdateOne {
	if t != nil {
		ruo.SetPublishedAt(*t)
	}
	return ruo
}

// ClearPublishedAt clears the value of the "published_at" field.
func (ruo *RecreationUpdateOne) ClearPublishedAt() *RecreationUpdateOne {
	ruo.mutation.ClearPublishedAt()
	return ruo
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (ruo *RecreationUpdateOne) SetUsersID(id int) *RecreationUpdateOne {
	ruo.mutation.SetUsersID(id)
	return ruo
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (ruo *RecreationUpdateOne) SetNillableUsersID(id *int) *RecreationUpdateOne {
	if id != nil {
		ruo = ruo.SetUsersID(*id)
	}
	return ruo
}

// SetUsers sets the "users" edge to the User entity.
func (ruo *RecreationUpdateOne) SetUsers(u *User) *RecreationUpdateOne {
	return ruo.SetUsersID(u.ID)
}

// Mutation returns the RecreationMutation object of the builder.
func (ruo *RecreationUpdateOne) Mutation() *RecreationMutation {
	return ruo.mutation
}

// ClearUsers clears the "users" edge to the User entity.
func (ruo *RecreationUpdateOne) ClearUsers() *RecreationUpdateOne {
	ruo.mutation.ClearUsers()
	return ruo
}

// Where appends a list predicates to the RecreationUpdate builder.
func (ruo *RecreationUpdateOne) Where(ps ...predicate.Recreation) *RecreationUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RecreationUpdateOne) Select(field string, fields ...string) *RecreationUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Recreation entity.
func (ruo *RecreationUpdateOne) Save(ctx context.Context) (*Recreation, error) {
	ruo.defaults()
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RecreationUpdateOne) SaveX(ctx context.Context) *Recreation {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RecreationUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RecreationUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *RecreationUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		v := recreation.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ruo *RecreationUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *RecreationUpdateOne {
	ruo.modifiers = append(ruo.modifiers, modifiers...)
	return ruo
}

func (ruo *RecreationUpdateOne) sqlSave(ctx context.Context) (_node *Recreation, err error) {
	_spec := sqlgraph.NewUpdateSpec(recreation.Table, recreation.Columns, sqlgraph.NewFieldSpec(recreation.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Recreation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, recreation.FieldID)
		for _, f := range fields {
			if !recreation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != recreation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Genre(); ok {
		_spec.SetField(recreation.FieldGenre, field.TypeJSON, value)
	}
	if value, ok := ruo.mutation.AppendedGenre(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, recreation.FieldGenre, value)
		})
	}
	if value, ok := ruo.mutation.Title(); ok {
		_spec.SetField(recreation.FieldTitle, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Content(); ok {
		_spec.SetField(recreation.FieldContent, field.TypeString, value)
	}
	if value, ok := ruo.mutation.YoutubeID(); ok {
		_spec.SetField(recreation.FieldYoutubeID, field.TypeString, value)
	}
	if ruo.mutation.YoutubeIDCleared() {
		_spec.ClearField(recreation.FieldYoutubeID, field.TypeString)
	}
	if value, ok := ruo.mutation.TargetNumber(); ok {
		_spec.SetField(recreation.FieldTargetNumber, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedTargetNumber(); ok {
		_spec.AddField(recreation.FieldTargetNumber, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.RequiredTime(); ok {
		_spec.SetField(recreation.FieldRequiredTime, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedRequiredTime(); ok {
		_spec.AddField(recreation.FieldRequiredTime, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.Publish(); ok {
		_spec.SetField(recreation.FieldPublish, field.TypeBool, value)
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.SetField(recreation.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ruo.mutation.PublishedAt(); ok {
		_spec.SetField(recreation.FieldPublishedAt, field.TypeTime, value)
	}
	if ruo.mutation.PublishedAtCleared() {
		_spec.ClearField(recreation.FieldPublishedAt, field.TypeTime)
	}
	if ruo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recreation.UsersTable,
			Columns: []string{recreation.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recreation.UsersTable,
			Columns: []string{recreation.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(ruo.modifiers...)
	_node = &Recreation{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{recreation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
