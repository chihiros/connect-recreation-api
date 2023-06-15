// Code generated by ent, DO NOT EDIT.

package ent

import (
	"app/ent/predicate"
	"app/ent/profile"
	"app/ent/recreation"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProfileUpdate is the builder for updating Profile entities.
type ProfileUpdate struct {
	config
	hooks     []Hook
	mutation  *ProfileMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ProfileUpdate builder.
func (pu *ProfileUpdate) Where(ps ...predicate.Profile) *ProfileUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetNickname sets the "nickname" field.
func (pu *ProfileUpdate) SetNickname(s string) *ProfileUpdate {
	pu.mutation.SetNickname(s)
	return pu
}

// SetIconURL sets the "icon_url" field.
func (pu *ProfileUpdate) SetIconURL(s string) *ProfileUpdate {
	pu.mutation.SetIconURL(s)
	return pu
}

// SetNillableIconURL sets the "icon_url" field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableIconURL(s *string) *ProfileUpdate {
	if s != nil {
		pu.SetIconURL(*s)
	}
	return pu
}

// ClearIconURL clears the value of the "icon_url" field.
func (pu *ProfileUpdate) ClearIconURL() *ProfileUpdate {
	pu.mutation.ClearIconURL()
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *ProfileUpdate) SetCreatedAt(t time.Time) *ProfileUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableCreatedAt(t *time.Time) *ProfileUpdate {
	if t != nil {
		pu.SetCreatedAt(*t)
	}
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *ProfileUpdate) SetUpdatedAt(t time.Time) *ProfileUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableUpdatedAt(t *time.Time) *ProfileUpdate {
	if t != nil {
		pu.SetUpdatedAt(*t)
	}
	return pu
}

// AddRecreationIDs adds the "recreations" edge to the Recreation entity by IDs.
func (pu *ProfileUpdate) AddRecreationIDs(ids ...int) *ProfileUpdate {
	pu.mutation.AddRecreationIDs(ids...)
	return pu
}

// AddRecreations adds the "recreations" edges to the Recreation entity.
func (pu *ProfileUpdate) AddRecreations(r ...*Recreation) *ProfileUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.AddRecreationIDs(ids...)
}

// Mutation returns the ProfileMutation object of the builder.
func (pu *ProfileUpdate) Mutation() *ProfileMutation {
	return pu.mutation
}

// ClearRecreations clears all "recreations" edges to the Recreation entity.
func (pu *ProfileUpdate) ClearRecreations() *ProfileUpdate {
	pu.mutation.ClearRecreations()
	return pu
}

// RemoveRecreationIDs removes the "recreations" edge to Recreation entities by IDs.
func (pu *ProfileUpdate) RemoveRecreationIDs(ids ...int) *ProfileUpdate {
	pu.mutation.RemoveRecreationIDs(ids...)
	return pu
}

// RemoveRecreations removes "recreations" edges to Recreation entities.
func (pu *ProfileUpdate) RemoveRecreations(r ...*Recreation) *ProfileUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.RemoveRecreationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProfileUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProfileUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProfileUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProfileUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pu *ProfileUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ProfileUpdate {
	pu.modifiers = append(pu.modifiers, modifiers...)
	return pu
}

func (pu *ProfileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(profile.Table, profile.Columns, sqlgraph.NewFieldSpec(profile.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Nickname(); ok {
		_spec.SetField(profile.FieldNickname, field.TypeString, value)
	}
	if value, ok := pu.mutation.IconURL(); ok {
		_spec.SetField(profile.FieldIconURL, field.TypeString, value)
	}
	if pu.mutation.IconURLCleared() {
		_spec.ClearField(profile.FieldIconURL, field.TypeString)
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.SetField(profile.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(profile.FieldUpdatedAt, field.TypeTime, value)
	}
	if pu.mutation.RecreationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.RecreationsTable,
			Columns: []string{profile.RecreationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(recreation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedRecreationsIDs(); len(nodes) > 0 && !pu.mutation.RecreationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.RecreationsTable,
			Columns: []string{profile.RecreationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(recreation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RecreationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.RecreationsTable,
			Columns: []string{profile.RecreationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(recreation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(pu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{profile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProfileUpdateOne is the builder for updating a single Profile entity.
type ProfileUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ProfileMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetNickname sets the "nickname" field.
func (puo *ProfileUpdateOne) SetNickname(s string) *ProfileUpdateOne {
	puo.mutation.SetNickname(s)
	return puo
}

// SetIconURL sets the "icon_url" field.
func (puo *ProfileUpdateOne) SetIconURL(s string) *ProfileUpdateOne {
	puo.mutation.SetIconURL(s)
	return puo
}

// SetNillableIconURL sets the "icon_url" field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableIconURL(s *string) *ProfileUpdateOne {
	if s != nil {
		puo.SetIconURL(*s)
	}
	return puo
}

// ClearIconURL clears the value of the "icon_url" field.
func (puo *ProfileUpdateOne) ClearIconURL() *ProfileUpdateOne {
	puo.mutation.ClearIconURL()
	return puo
}

// SetCreatedAt sets the "created_at" field.
func (puo *ProfileUpdateOne) SetCreatedAt(t time.Time) *ProfileUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableCreatedAt(t *time.Time) *ProfileUpdateOne {
	if t != nil {
		puo.SetCreatedAt(*t)
	}
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *ProfileUpdateOne) SetUpdatedAt(t time.Time) *ProfileUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableUpdatedAt(t *time.Time) *ProfileUpdateOne {
	if t != nil {
		puo.SetUpdatedAt(*t)
	}
	return puo
}

// AddRecreationIDs adds the "recreations" edge to the Recreation entity by IDs.
func (puo *ProfileUpdateOne) AddRecreationIDs(ids ...int) *ProfileUpdateOne {
	puo.mutation.AddRecreationIDs(ids...)
	return puo
}

// AddRecreations adds the "recreations" edges to the Recreation entity.
func (puo *ProfileUpdateOne) AddRecreations(r ...*Recreation) *ProfileUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.AddRecreationIDs(ids...)
}

// Mutation returns the ProfileMutation object of the builder.
func (puo *ProfileUpdateOne) Mutation() *ProfileMutation {
	return puo.mutation
}

// ClearRecreations clears all "recreations" edges to the Recreation entity.
func (puo *ProfileUpdateOne) ClearRecreations() *ProfileUpdateOne {
	puo.mutation.ClearRecreations()
	return puo
}

// RemoveRecreationIDs removes the "recreations" edge to Recreation entities by IDs.
func (puo *ProfileUpdateOne) RemoveRecreationIDs(ids ...int) *ProfileUpdateOne {
	puo.mutation.RemoveRecreationIDs(ids...)
	return puo
}

// RemoveRecreations removes "recreations" edges to Recreation entities.
func (puo *ProfileUpdateOne) RemoveRecreations(r ...*Recreation) *ProfileUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.RemoveRecreationIDs(ids...)
}

// Where appends a list predicates to the ProfileUpdate builder.
func (puo *ProfileUpdateOne) Where(ps ...predicate.Profile) *ProfileUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProfileUpdateOne) Select(field string, fields ...string) *ProfileUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Profile entity.
func (puo *ProfileUpdateOne) Save(ctx context.Context) (*Profile, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProfileUpdateOne) SaveX(ctx context.Context) *Profile {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProfileUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProfileUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (puo *ProfileUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ProfileUpdateOne {
	puo.modifiers = append(puo.modifiers, modifiers...)
	return puo
}

func (puo *ProfileUpdateOne) sqlSave(ctx context.Context) (_node *Profile, err error) {
	_spec := sqlgraph.NewUpdateSpec(profile.Table, profile.Columns, sqlgraph.NewFieldSpec(profile.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Profile.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, profile.FieldID)
		for _, f := range fields {
			if !profile.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != profile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Nickname(); ok {
		_spec.SetField(profile.FieldNickname, field.TypeString, value)
	}
	if value, ok := puo.mutation.IconURL(); ok {
		_spec.SetField(profile.FieldIconURL, field.TypeString, value)
	}
	if puo.mutation.IconURLCleared() {
		_spec.ClearField(profile.FieldIconURL, field.TypeString)
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.SetField(profile.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(profile.FieldUpdatedAt, field.TypeTime, value)
	}
	if puo.mutation.RecreationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.RecreationsTable,
			Columns: []string{profile.RecreationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(recreation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedRecreationsIDs(); len(nodes) > 0 && !puo.mutation.RecreationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.RecreationsTable,
			Columns: []string{profile.RecreationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(recreation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RecreationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.RecreationsTable,
			Columns: []string{profile.RecreationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(recreation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(puo.modifiers...)
	_node = &Profile{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{profile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
