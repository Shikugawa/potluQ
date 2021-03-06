// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/Shikugawa/potluq/ent/club"
	"github.com/Shikugawa/potluq/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ClubDelete is the builder for deleting a Club entity.
type ClubDelete struct {
	config
	predicates []predicate.Club
}

// Where adds a new predicate to the delete builder.
func (cd *ClubDelete) Where(ps ...predicate.Club) *ClubDelete {
	cd.predicates = append(cd.predicates, ps...)
	return cd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cd *ClubDelete) Exec(ctx context.Context) (int, error) {
	return cd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (cd *ClubDelete) ExecX(ctx context.Context) int {
	n, err := cd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cd *ClubDelete) sqlExec(ctx context.Context) (int, error) {
	spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: club.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: club.FieldID,
			},
		},
	}
	if ps := cd.predicates; len(ps) > 0 {
		spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, cd.driver, spec)
}

// ClubDeleteOne is the builder for deleting a single Club entity.
type ClubDeleteOne struct {
	cd *ClubDelete
}

// Exec executes the deletion query.
func (cdo *ClubDeleteOne) Exec(ctx context.Context) error {
	n, err := cdo.cd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &ErrNotFound{club.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *ClubDeleteOne) ExecX(ctx context.Context) {
	cdo.cd.ExecX(ctx)
}
