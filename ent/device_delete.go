// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/Shikugawa/potluq/ent/device"
	"github.com/Shikugawa/potluq/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// DeviceDelete is the builder for deleting a Device entity.
type DeviceDelete struct {
	config
	predicates []predicate.Device
}

// Where adds a new predicate to the delete builder.
func (dd *DeviceDelete) Where(ps ...predicate.Device) *DeviceDelete {
	dd.predicates = append(dd.predicates, ps...)
	return dd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dd *DeviceDelete) Exec(ctx context.Context) (int, error) {
	return dd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (dd *DeviceDelete) ExecX(ctx context.Context) int {
	n, err := dd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dd *DeviceDelete) sqlExec(ctx context.Context) (int, error) {
	spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: device.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: device.FieldID,
			},
		},
	}
	if ps := dd.predicates; len(ps) > 0 {
		spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, dd.driver, spec)
}

// DeviceDeleteOne is the builder for deleting a single Device entity.
type DeviceDeleteOne struct {
	dd *DeviceDelete
}

// Exec executes the deletion query.
func (ddo *DeviceDeleteOne) Exec(ctx context.Context) error {
	n, err := ddo.dd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &ErrNotFound{device.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ddo *DeviceDeleteOne) ExecX(ctx context.Context) {
	ddo.dd.ExecX(ctx)
}
