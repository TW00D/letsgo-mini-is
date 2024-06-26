// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"letsgo-mini-is/internal/adapter/storage/mysql/ent/picture"
	"letsgo-mini-is/internal/adapter/storage/mysql/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PictureUpdate is the builder for updating Picture entities.
type PictureUpdate struct {
	config
	hooks    []Hook
	mutation *PictureMutation
}

// Where appends a list predicates to the PictureUpdate builder.
func (pu *PictureUpdate) Where(ps ...predicate.Picture) *PictureUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetContent sets the "content" field.
func (pu *PictureUpdate) SetContent(b []byte) *PictureUpdate {
	pu.mutation.SetContent(b)
	return pu
}

// SetExtension sets the "extension" field.
func (pu *PictureUpdate) SetExtension(s string) *PictureUpdate {
	pu.mutation.SetExtension(s)
	return pu
}

// SetNillableExtension sets the "extension" field if the given value is not nil.
func (pu *PictureUpdate) SetNillableExtension(s *string) *PictureUpdate {
	if s != nil {
		pu.SetExtension(*s)
	}
	return pu
}

// Mutation returns the PictureMutation object of the builder.
func (pu *PictureUpdate) Mutation() *PictureMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PictureUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PictureUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PictureUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PictureUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PictureUpdate) check() error {
	if v, ok := pu.mutation.Content(); ok {
		if err := picture.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Picture.content": %w`, err)}
		}
	}
	return nil
}

func (pu *PictureUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(picture.Table, picture.Columns, sqlgraph.NewFieldSpec(picture.FieldID, field.TypeUUID))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Content(); ok {
		_spec.SetField(picture.FieldContent, field.TypeBytes, value)
	}
	if value, ok := pu.mutation.Extension(); ok {
		_spec.SetField(picture.FieldExtension, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{picture.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PictureUpdateOne is the builder for updating a single Picture entity.
type PictureUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PictureMutation
}

// SetContent sets the "content" field.
func (puo *PictureUpdateOne) SetContent(b []byte) *PictureUpdateOne {
	puo.mutation.SetContent(b)
	return puo
}

// SetExtension sets the "extension" field.
func (puo *PictureUpdateOne) SetExtension(s string) *PictureUpdateOne {
	puo.mutation.SetExtension(s)
	return puo
}

// SetNillableExtension sets the "extension" field if the given value is not nil.
func (puo *PictureUpdateOne) SetNillableExtension(s *string) *PictureUpdateOne {
	if s != nil {
		puo.SetExtension(*s)
	}
	return puo
}

// Mutation returns the PictureMutation object of the builder.
func (puo *PictureUpdateOne) Mutation() *PictureMutation {
	return puo.mutation
}

// Where appends a list predicates to the PictureUpdate builder.
func (puo *PictureUpdateOne) Where(ps ...predicate.Picture) *PictureUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PictureUpdateOne) Select(field string, fields ...string) *PictureUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Picture entity.
func (puo *PictureUpdateOne) Save(ctx context.Context) (*Picture, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PictureUpdateOne) SaveX(ctx context.Context) *Picture {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PictureUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PictureUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PictureUpdateOne) check() error {
	if v, ok := puo.mutation.Content(); ok {
		if err := picture.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Picture.content": %w`, err)}
		}
	}
	return nil
}

func (puo *PictureUpdateOne) sqlSave(ctx context.Context) (_node *Picture, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(picture.Table, picture.Columns, sqlgraph.NewFieldSpec(picture.FieldID, field.TypeUUID))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Picture.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, picture.FieldID)
		for _, f := range fields {
			if !picture.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != picture.FieldID {
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
	if value, ok := puo.mutation.Content(); ok {
		_spec.SetField(picture.FieldContent, field.TypeBytes, value)
	}
	if value, ok := puo.mutation.Extension(); ok {
		_spec.SetField(picture.FieldExtension, field.TypeString, value)
	}
	_node = &Picture{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{picture.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
