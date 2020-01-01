// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/Shikugawa/potluq/ent/club"
	"github.com/Shikugawa/potluq/ent/device"
	"github.com/Shikugawa/potluq/ent/predicate"
	"github.com/Shikugawa/potluq/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// DeviceQuery is the builder for querying Device entities.
type DeviceQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Device
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (dq *DeviceQuery) Where(ps ...predicate.Device) *DeviceQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit adds a limit step to the query.
func (dq *DeviceQuery) Limit(limit int) *DeviceQuery {
	dq.limit = &limit
	return dq
}

// Offset adds an offset step to the query.
func (dq *DeviceQuery) Offset(offset int) *DeviceQuery {
	dq.offset = &offset
	return dq
}

// Order adds an order step to the query.
func (dq *DeviceQuery) Order(o ...Order) *DeviceQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryClub chains the current query on the club edge.
func (dq *DeviceQuery) QueryClub() *ClubQuery {
	query := &ClubQuery{config: dq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(device.Table, device.FieldID, dq.sqlQuery()),
		sqlgraph.To(club.Table, club.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, device.ClubTable, device.ClubColumn),
	)
	query.sql = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
	return query
}

// QueryUser chains the current query on the user edge.
func (dq *DeviceQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: dq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(device.Table, device.FieldID, dq.sqlQuery()),
		sqlgraph.To(user.Table, user.FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, device.UserTable, device.UserPrimaryKey...),
	)
	query.sql = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
	return query
}

// First returns the first Device entity in the query. Returns *ErrNotFound when no device was found.
func (dq *DeviceQuery) First(ctx context.Context) (*Device, error) {
	ds, err := dq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(ds) == 0 {
		return nil, &ErrNotFound{device.Label}
	}
	return ds[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DeviceQuery) FirstX(ctx context.Context) *Device {
	d, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return d
}

// FirstID returns the first Device id in the query. Returns *ErrNotFound when no id was found.
func (dq *DeviceQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &ErrNotFound{device.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (dq *DeviceQuery) FirstXID(ctx context.Context) int {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Device entity in the query, returns an error if not exactly one entity was returned.
func (dq *DeviceQuery) Only(ctx context.Context) (*Device, error) {
	ds, err := dq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(ds) {
	case 1:
		return ds[0], nil
	case 0:
		return nil, &ErrNotFound{device.Label}
	default:
		return nil, &ErrNotSingular{device.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DeviceQuery) OnlyX(ctx context.Context) *Device {
	d, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// OnlyID returns the only Device id in the query, returns an error if not exactly one id was returned.
func (dq *DeviceQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &ErrNotFound{device.Label}
	default:
		err = &ErrNotSingular{device.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (dq *DeviceQuery) OnlyXID(ctx context.Context) int {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Devices.
func (dq *DeviceQuery) All(ctx context.Context) ([]*Device, error) {
	return dq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dq *DeviceQuery) AllX(ctx context.Context) []*Device {
	ds, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return ds
}

// IDs executes the query and returns a list of Device ids.
func (dq *DeviceQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := dq.Select(device.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DeviceQuery) IDsX(ctx context.Context) []int {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DeviceQuery) Count(ctx context.Context) (int, error) {
	return dq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DeviceQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DeviceQuery) Exist(ctx context.Context) (bool, error) {
	return dq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DeviceQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DeviceQuery) Clone() *DeviceQuery {
	return &DeviceQuery{
		config:     dq.config,
		limit:      dq.limit,
		offset:     dq.offset,
		order:      append([]Order{}, dq.order...),
		unique:     append([]string{}, dq.unique...),
		predicates: append([]predicate.Device{}, dq.predicates...),
		// clone intermediate query.
		sql: dq.sql.Clone(),
	}
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Role device.Role `json:"role,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Device.Query().
//		GroupBy(device.FieldRole).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (dq *DeviceQuery) GroupBy(field string, fields ...string) *DeviceGroupBy {
	group := &DeviceGroupBy{config: dq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = dq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Role device.Role `json:"role,omitempty"`
//	}
//
//	client.Device.Query().
//		Select(device.FieldRole).
//		Scan(ctx, &v)
//
func (dq *DeviceQuery) Select(field string, fields ...string) *DeviceSelect {
	selector := &DeviceSelect{config: dq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = dq.sqlQuery()
	return selector
}

func (dq *DeviceQuery) sqlAll(ctx context.Context) ([]*Device, error) {
	var (
		nodes []*Device
		spec  = dq.querySpec()
	)
	spec.ScanValues = func() []interface{} {
		node := &Device{config: dq.config}
		nodes = append(nodes, node)
		return node.scanValues()
	}
	spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, spec); err != nil {
		return nil, err
	}
	return nodes, nil
}

func (dq *DeviceQuery) sqlCount(ctx context.Context) (int, error) {
	spec := dq.querySpec()
	return sqlgraph.CountNodes(ctx, dq.driver, spec)
}

func (dq *DeviceQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (dq *DeviceQuery) querySpec() *sqlgraph.QuerySpec {
	spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   device.Table,
			Columns: device.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: device.FieldID,
			},
		},
		From:   dq.sql,
		Unique: true,
	}
	if ps := dq.predicates; len(ps) > 0 {
		spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.limit; limit != nil {
		spec.Limit = *limit
	}
	if offset := dq.offset; offset != nil {
		spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return spec
}

func (dq *DeviceQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(device.Table)
	selector := builder.Select(t1.Columns(device.Columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(device.Columns...)...)
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DeviceGroupBy is the builder for group-by Device entities.
type DeviceGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DeviceGroupBy) Aggregate(fns ...Aggregate) *DeviceGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the group-by query and scan the result into the given value.
func (dgb *DeviceGroupBy) Scan(ctx context.Context, v interface{}) error {
	return dgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dgb *DeviceGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := dgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (dgb *DeviceGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DeviceGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dgb *DeviceGroupBy) StringsX(ctx context.Context) []string {
	v, err := dgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (dgb *DeviceGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DeviceGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dgb *DeviceGroupBy) IntsX(ctx context.Context) []int {
	v, err := dgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (dgb *DeviceGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DeviceGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dgb *DeviceGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := dgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (dgb *DeviceGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DeviceGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dgb *DeviceGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := dgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dgb *DeviceGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dgb.sqlQuery().Query()
	if err := dgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dgb *DeviceGroupBy) sqlQuery() *sql.Selector {
	selector := dgb.sql
	columns := make([]string, 0, len(dgb.fields)+len(dgb.fns))
	columns = append(columns, dgb.fields...)
	for _, fn := range dgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(dgb.fields...)
}

// DeviceSelect is the builder for select fields of Device entities.
type DeviceSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (ds *DeviceSelect) Scan(ctx context.Context, v interface{}) error {
	return ds.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ds *DeviceSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ds.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ds *DeviceSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DeviceSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ds *DeviceSelect) StringsX(ctx context.Context) []string {
	v, err := ds.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ds *DeviceSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DeviceSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ds *DeviceSelect) IntsX(ctx context.Context) []int {
	v, err := ds.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ds *DeviceSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DeviceSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ds *DeviceSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ds.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ds *DeviceSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DeviceSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ds *DeviceSelect) BoolsX(ctx context.Context) []bool {
	v, err := ds.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ds *DeviceSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ds.sqlQuery().Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ds *DeviceSelect) sqlQuery() sql.Querier {
	selector := ds.sql
	selector.Select(selector.Columns(ds.fields...)...)
	return selector
}
