// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/Shikugawa/potluq/ent/club"
	"github.com/Shikugawa/potluq/ent/music"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// MusicCreate is the builder for creating a Music entity.
type MusicCreate struct {
	config
	club map[int]struct{}
}

// SetClubID sets the club edge to Club by id.
func (mc *MusicCreate) SetClubID(id int) *MusicCreate {
	if mc.club == nil {
		mc.club = make(map[int]struct{})
	}
	mc.club[id] = struct{}{}
	return mc
}

// SetNillableClubID sets the club edge to Club by id if the given value is not nil.
func (mc *MusicCreate) SetNillableClubID(id *int) *MusicCreate {
	if id != nil {
		mc = mc.SetClubID(*id)
	}
	return mc
}

// SetClub sets the club edge to Club.
func (mc *MusicCreate) SetClub(c *Club) *MusicCreate {
	return mc.SetClubID(c.ID)
}

// Save creates the Music in the database.
func (mc *MusicCreate) Save(ctx context.Context) (*Music, error) {
	if len(mc.club) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"club\"")
	}
	return mc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MusicCreate) SaveX(ctx context.Context) *Music {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mc *MusicCreate) sqlSave(ctx context.Context) (*Music, error) {
	var (
		m    = &Music{config: mc.config}
		spec = &sqlgraph.CreateSpec{
			Table: music.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: music.FieldID,
			},
		}
	)
	if nodes := mc.club; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   music.ClubTable,
			Columns: []string{music.ClubColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: club.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges = append(spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, mc.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := spec.ID.Value.(int64)
	m.ID = int(id)
	return m, nil
}
