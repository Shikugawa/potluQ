// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/Shikugawa/potraq/ent/club"
	"github.com/Shikugawa/potraq/ent/device"
	"github.com/Shikugawa/potraq/ent/predicate"
	"github.com/Shikugawa/potraq/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// DeviceUpdate is the builder for updating Device entities.
type DeviceUpdate struct {
	config
	role        *device.Role
	club        map[int]struct{}
	user        map[int]struct{}
	clearedClub bool
	removedUser map[int]struct{}
	predicates  []predicate.Device
}

// Where adds a new predicate for the builder.
func (du *DeviceUpdate) Where(ps ...predicate.Device) *DeviceUpdate {
	du.predicates = append(du.predicates, ps...)
	return du
}

// SetRole sets the role field.
func (du *DeviceUpdate) SetRole(d device.Role) *DeviceUpdate {
	du.role = &d
	return du
}

// SetClubID sets the club edge to Club by id.
func (du *DeviceUpdate) SetClubID(id int) *DeviceUpdate {
	if du.club == nil {
		du.club = make(map[int]struct{})
	}
	du.club[id] = struct{}{}
	return du
}

// SetNillableClubID sets the club edge to Club by id if the given value is not nil.
func (du *DeviceUpdate) SetNillableClubID(id *int) *DeviceUpdate {
	if id != nil {
		du = du.SetClubID(*id)
	}
	return du
}

// SetClub sets the club edge to Club.
func (du *DeviceUpdate) SetClub(c *Club) *DeviceUpdate {
	return du.SetClubID(c.ID)
}

// AddUserIDs adds the user edge to User by ids.
func (du *DeviceUpdate) AddUserIDs(ids ...int) *DeviceUpdate {
	if du.user == nil {
		du.user = make(map[int]struct{})
	}
	for i := range ids {
		du.user[ids[i]] = struct{}{}
	}
	return du
}

// AddUser adds the user edges to User.
func (du *DeviceUpdate) AddUser(u ...*User) *DeviceUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return du.AddUserIDs(ids...)
}

// ClearClub clears the club edge to Club.
func (du *DeviceUpdate) ClearClub() *DeviceUpdate {
	du.clearedClub = true
	return du
}

// RemoveUserIDs removes the user edge to User by ids.
func (du *DeviceUpdate) RemoveUserIDs(ids ...int) *DeviceUpdate {
	if du.removedUser == nil {
		du.removedUser = make(map[int]struct{})
	}
	for i := range ids {
		du.removedUser[ids[i]] = struct{}{}
	}
	return du
}

// RemoveUser removes user edges to User.
func (du *DeviceUpdate) RemoveUser(u ...*User) *DeviceUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return du.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (du *DeviceUpdate) Save(ctx context.Context) (int, error) {
	if du.role != nil {
		if err := device.RoleValidator(*du.role); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"role\": %v", err)
		}
	}
	if len(du.club) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"club\"")
	}
	return du.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DeviceUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DeviceUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DeviceUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DeviceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   device.Table,
			Columns: device.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: device.FieldID,
			},
		},
	}
	if ps := du.predicates; len(ps) > 0 {
		spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := du.role; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: device.FieldRole,
		})
	}
	if du.clearedClub {
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
		spec.Edges.Clear = append(spec.Edges.Clear, edge)
	}
	if nodes := du.club; len(nodes) > 0 {
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
		spec.Edges.Add = append(spec.Edges.Add, edge)
	}
	if nodes := du.removedUser; len(nodes) > 0 {
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
		spec.Edges.Clear = append(spec.Edges.Clear, edge)
	}
	if nodes := du.user; len(nodes) > 0 {
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
		spec.Edges.Add = append(spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DeviceUpdateOne is the builder for updating a single Device entity.
type DeviceUpdateOne struct {
	config
	id          int
	role        *device.Role
	club        map[int]struct{}
	user        map[int]struct{}
	clearedClub bool
	removedUser map[int]struct{}
}

// SetRole sets the role field.
func (duo *DeviceUpdateOne) SetRole(d device.Role) *DeviceUpdateOne {
	duo.role = &d
	return duo
}

// SetClubID sets the club edge to Club by id.
func (duo *DeviceUpdateOne) SetClubID(id int) *DeviceUpdateOne {
	if duo.club == nil {
		duo.club = make(map[int]struct{})
	}
	duo.club[id] = struct{}{}
	return duo
}

// SetNillableClubID sets the club edge to Club by id if the given value is not nil.
func (duo *DeviceUpdateOne) SetNillableClubID(id *int) *DeviceUpdateOne {
	if id != nil {
		duo = duo.SetClubID(*id)
	}
	return duo
}

// SetClub sets the club edge to Club.
func (duo *DeviceUpdateOne) SetClub(c *Club) *DeviceUpdateOne {
	return duo.SetClubID(c.ID)
}

// AddUserIDs adds the user edge to User by ids.
func (duo *DeviceUpdateOne) AddUserIDs(ids ...int) *DeviceUpdateOne {
	if duo.user == nil {
		duo.user = make(map[int]struct{})
	}
	for i := range ids {
		duo.user[ids[i]] = struct{}{}
	}
	return duo
}

// AddUser adds the user edges to User.
func (duo *DeviceUpdateOne) AddUser(u ...*User) *DeviceUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return duo.AddUserIDs(ids...)
}

// ClearClub clears the club edge to Club.
func (duo *DeviceUpdateOne) ClearClub() *DeviceUpdateOne {
	duo.clearedClub = true
	return duo
}

// RemoveUserIDs removes the user edge to User by ids.
func (duo *DeviceUpdateOne) RemoveUserIDs(ids ...int) *DeviceUpdateOne {
	if duo.removedUser == nil {
		duo.removedUser = make(map[int]struct{})
	}
	for i := range ids {
		duo.removedUser[ids[i]] = struct{}{}
	}
	return duo
}

// RemoveUser removes user edges to User.
func (duo *DeviceUpdateOne) RemoveUser(u ...*User) *DeviceUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return duo.RemoveUserIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (duo *DeviceUpdateOne) Save(ctx context.Context) (*Device, error) {
	if duo.role != nil {
		if err := device.RoleValidator(*duo.role); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"role\": %v", err)
		}
	}
	if len(duo.club) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"club\"")
	}
	return duo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DeviceUpdateOne) SaveX(ctx context.Context) *Device {
	d, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// Exec executes the query on the entity.
func (duo *DeviceUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DeviceUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DeviceUpdateOne) sqlSave(ctx context.Context) (d *Device, err error) {
	spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   device.Table,
			Columns: device.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  duo.id,
				Type:   field.TypeInt,
				Column: device.FieldID,
			},
		},
	}
	if value := duo.role; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: device.FieldRole,
		})
	}
	if duo.clearedClub {
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
		spec.Edges.Clear = append(spec.Edges.Clear, edge)
	}
	if nodes := duo.club; len(nodes) > 0 {
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
		spec.Edges.Add = append(spec.Edges.Add, edge)
	}
	if nodes := duo.removedUser; len(nodes) > 0 {
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
		spec.Edges.Clear = append(spec.Edges.Clear, edge)
	}
	if nodes := duo.user; len(nodes) > 0 {
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
		spec.Edges.Add = append(spec.Edges.Add, edge)
	}
	d = &Device{config: duo.config}
	spec.Assign = d.assignValues
	spec.ScanValues = d.scanValues()
	if err = sqlgraph.UpdateNode(ctx, duo.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return d, nil
}