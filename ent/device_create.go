// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/Shikugawa/potraq/ent/club"
	"github.com/Shikugawa/potraq/ent/device"
	"github.com/Shikugawa/potraq/ent/user"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// DeviceCreate is the builder for creating a Device entity.
type DeviceCreate struct {
	config
	role *device.Role
	club map[int]struct{}
	user map[int]struct{}
}

// SetRole sets the role field.
func (dc *DeviceCreate) SetRole(d device.Role) *DeviceCreate {
	dc.role = &d
	return dc
}

// SetClubID sets the club edge to Club by id.
func (dc *DeviceCreate) SetClubID(id int) *DeviceCreate {
	if dc.club == nil {
		dc.club = make(map[int]struct{})
	}
	dc.club[id] = struct{}{}
	return dc
}

// SetNillableClubID sets the club edge to Club by id if the given value is not nil.
func (dc *DeviceCreate) SetNillableClubID(id *int) *DeviceCreate {
	if id != nil {
		dc = dc.SetClubID(*id)
	}
	return dc
}

// SetClub sets the club edge to Club.
func (dc *DeviceCreate) SetClub(c *Club) *DeviceCreate {
	return dc.SetClubID(c.ID)
}

// AddUserIDs adds the user edge to User by ids.
func (dc *DeviceCreate) AddUserIDs(ids ...int) *DeviceCreate {
	if dc.user == nil {
		dc.user = make(map[int]struct{})
	}
	for i := range ids {
		dc.user[ids[i]] = struct{}{}
	}
	return dc
}

// AddUser adds the user edges to User.
func (dc *DeviceCreate) AddUser(u ...*User) *DeviceCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return dc.AddUserIDs(ids...)
}

// Save creates the Device in the database.
func (dc *DeviceCreate) Save(ctx context.Context) (*Device, error) {
	if dc.role == nil {
		return nil, errors.New("ent: missing required field \"role\"")
	}
	if err := device.RoleValidator(*dc.role); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"role\": %v", err)
	}
	if len(dc.club) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"club\"")
	}
	return dc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DeviceCreate) SaveX(ctx context.Context) *Device {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dc *DeviceCreate) sqlSave(ctx context.Context) (*Device, error) {
	var (
		d    = &Device{config: dc.config}
		spec = &sqlgraph.CreateSpec{
			Table: device.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: device.FieldID,
			},
		}
	)
	if value := dc.role; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: device.FieldRole,
		})
		d.Role = *value
	}
	if nodes := dc.club; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   device.ClubTable,
			Columns: []string{device.ClubColumn},
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
	if nodes := dc.user; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   device.UserTable,
			Columns: device.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges = append(spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, dc.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := spec.ID.Value.(int64)
	d.ID = int(id)
	return d, nil
}