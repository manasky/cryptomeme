// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"github.com/manasky/cryptomeme/internal/ent/marketchart"
	"github.com/manasky/cryptomeme/internal/ent/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// MarketChartQuery is the builder for querying MarketChart entities.
type MarketChartQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.MarketChart
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MarketChartQuery builder.
func (mcq *MarketChartQuery) Where(ps ...predicate.MarketChart) *MarketChartQuery {
	mcq.predicates = append(mcq.predicates, ps...)
	return mcq
}

// Limit adds a limit step to the query.
func (mcq *MarketChartQuery) Limit(limit int) *MarketChartQuery {
	mcq.limit = &limit
	return mcq
}

// Offset adds an offset step to the query.
func (mcq *MarketChartQuery) Offset(offset int) *MarketChartQuery {
	mcq.offset = &offset
	return mcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mcq *MarketChartQuery) Unique(unique bool) *MarketChartQuery {
	mcq.unique = &unique
	return mcq
}

// Order adds an order step to the query.
func (mcq *MarketChartQuery) Order(o ...OrderFunc) *MarketChartQuery {
	mcq.order = append(mcq.order, o...)
	return mcq
}

// First returns the first MarketChart entity from the query.
// Returns a *NotFoundError when no MarketChart was found.
func (mcq *MarketChartQuery) First(ctx context.Context) (*MarketChart, error) {
	nodes, err := mcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{marketchart.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mcq *MarketChartQuery) FirstX(ctx context.Context) *MarketChart {
	node, err := mcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MarketChart ID from the query.
// Returns a *NotFoundError when no MarketChart ID was found.
func (mcq *MarketChartQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = mcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{marketchart.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mcq *MarketChartQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := mcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MarketChart entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one MarketChart entity is not found.
// Returns a *NotFoundError when no MarketChart entities are found.
func (mcq *MarketChartQuery) Only(ctx context.Context) (*MarketChart, error) {
	nodes, err := mcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{marketchart.Label}
	default:
		return nil, &NotSingularError{marketchart.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mcq *MarketChartQuery) OnlyX(ctx context.Context) *MarketChart {
	node, err := mcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MarketChart ID in the query.
// Returns a *NotSingularError when exactly one MarketChart ID is not found.
// Returns a *NotFoundError when no entities are found.
func (mcq *MarketChartQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = mcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{marketchart.Label}
	default:
		err = &NotSingularError{marketchart.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mcq *MarketChartQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := mcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MarketCharts.
func (mcq *MarketChartQuery) All(ctx context.Context) ([]*MarketChart, error) {
	if err := mcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mcq *MarketChartQuery) AllX(ctx context.Context) []*MarketChart {
	nodes, err := mcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MarketChart IDs.
func (mcq *MarketChartQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := mcq.Select(marketchart.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mcq *MarketChartQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := mcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mcq *MarketChartQuery) Count(ctx context.Context) (int, error) {
	if err := mcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mcq *MarketChartQuery) CountX(ctx context.Context) int {
	count, err := mcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mcq *MarketChartQuery) Exist(ctx context.Context) (bool, error) {
	if err := mcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mcq *MarketChartQuery) ExistX(ctx context.Context) bool {
	exist, err := mcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MarketChartQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mcq *MarketChartQuery) Clone() *MarketChartQuery {
	if mcq == nil {
		return nil
	}
	return &MarketChartQuery{
		config:     mcq.config,
		limit:      mcq.limit,
		offset:     mcq.offset,
		order:      append([]OrderFunc{}, mcq.order...),
		predicates: append([]predicate.MarketChart{}, mcq.predicates...),
		// clone intermediate query.
		sql:  mcq.sql.Clone(),
		path: mcq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.MarketChart.Query().
//		GroupBy(marketchart.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (mcq *MarketChartQuery) GroupBy(field string, fields ...string) *MarketChartGroupBy {
	group := &MarketChartGroupBy{config: mcq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mcq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.MarketChart.Query().
//		Select(marketchart.FieldCreateTime).
//		Scan(ctx, &v)
//
func (mcq *MarketChartQuery) Select(fields ...string) *MarketChartSelect {
	mcq.fields = append(mcq.fields, fields...)
	return &MarketChartSelect{MarketChartQuery: mcq}
}

func (mcq *MarketChartQuery) prepareQuery(ctx context.Context) error {
	for _, f := range mcq.fields {
		if !marketchart.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mcq.path != nil {
		prev, err := mcq.path(ctx)
		if err != nil {
			return err
		}
		mcq.sql = prev
	}
	return nil
}

func (mcq *MarketChartQuery) sqlAll(ctx context.Context) ([]*MarketChart, error) {
	var (
		nodes = []*MarketChart{}
		_spec = mcq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &MarketChart{config: mcq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, mcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (mcq *MarketChartQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mcq.querySpec()
	return sqlgraph.CountNodes(ctx, mcq.driver, _spec)
}

func (mcq *MarketChartQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (mcq *MarketChartQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   marketchart.Table,
			Columns: marketchart.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: marketchart.FieldID,
			},
		},
		From:   mcq.sql,
		Unique: true,
	}
	if unique := mcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, marketchart.FieldID)
		for i := range fields {
			if fields[i] != marketchart.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mcq *MarketChartQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mcq.driver.Dialect())
	t1 := builder.Table(marketchart.Table)
	columns := mcq.fields
	if len(columns) == 0 {
		columns = marketchart.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mcq.sql != nil {
		selector = mcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range mcq.predicates {
		p(selector)
	}
	for _, p := range mcq.order {
		p(selector)
	}
	if offset := mcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MarketChartGroupBy is the group-by builder for MarketChart entities.
type MarketChartGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mcgb *MarketChartGroupBy) Aggregate(fns ...AggregateFunc) *MarketChartGroupBy {
	mcgb.fns = append(mcgb.fns, fns...)
	return mcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (mcgb *MarketChartGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mcgb.path(ctx)
	if err != nil {
		return err
	}
	mcgb.sql = query
	return mcgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mcgb *MarketChartGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := mcgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (mcgb *MarketChartGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(mcgb.fields) > 1 {
		return nil, errors.New("ent: MarketChartGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := mcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mcgb *MarketChartGroupBy) StringsX(ctx context.Context) []string {
	v, err := mcgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mcgb *MarketChartGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mcgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{marketchart.Label}
	default:
		err = fmt.Errorf("ent: MarketChartGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mcgb *MarketChartGroupBy) StringX(ctx context.Context) string {
	v, err := mcgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (mcgb *MarketChartGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(mcgb.fields) > 1 {
		return nil, errors.New("ent: MarketChartGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := mcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mcgb *MarketChartGroupBy) IntsX(ctx context.Context) []int {
	v, err := mcgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mcgb *MarketChartGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mcgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{marketchart.Label}
	default:
		err = fmt.Errorf("ent: MarketChartGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mcgb *MarketChartGroupBy) IntX(ctx context.Context) int {
	v, err := mcgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (mcgb *MarketChartGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(mcgb.fields) > 1 {
		return nil, errors.New("ent: MarketChartGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := mcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mcgb *MarketChartGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := mcgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mcgb *MarketChartGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mcgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{marketchart.Label}
	default:
		err = fmt.Errorf("ent: MarketChartGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mcgb *MarketChartGroupBy) Float64X(ctx context.Context) float64 {
	v, err := mcgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (mcgb *MarketChartGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(mcgb.fields) > 1 {
		return nil, errors.New("ent: MarketChartGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := mcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mcgb *MarketChartGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := mcgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (mcgb *MarketChartGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mcgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{marketchart.Label}
	default:
		err = fmt.Errorf("ent: MarketChartGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mcgb *MarketChartGroupBy) BoolX(ctx context.Context) bool {
	v, err := mcgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mcgb *MarketChartGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range mcgb.fields {
		if !marketchart.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mcgb *MarketChartGroupBy) sqlQuery() *sql.Selector {
	selector := mcgb.sql.Select()
	aggregation := make([]string, 0, len(mcgb.fns))
	for _, fn := range mcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(mcgb.fields)+len(mcgb.fns))
		for _, f := range mcgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(mcgb.fields...)...)
}

// MarketChartSelect is the builder for selecting fields of MarketChart entities.
type MarketChartSelect struct {
	*MarketChartQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (mcs *MarketChartSelect) Scan(ctx context.Context, v interface{}) error {
	if err := mcs.prepareQuery(ctx); err != nil {
		return err
	}
	mcs.sql = mcs.MarketChartQuery.sqlQuery(ctx)
	return mcs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (mcs *MarketChartSelect) ScanX(ctx context.Context, v interface{}) {
	if err := mcs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (mcs *MarketChartSelect) Strings(ctx context.Context) ([]string, error) {
	if len(mcs.fields) > 1 {
		return nil, errors.New("ent: MarketChartSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := mcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (mcs *MarketChartSelect) StringsX(ctx context.Context) []string {
	v, err := mcs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (mcs *MarketChartSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = mcs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{marketchart.Label}
	default:
		err = fmt.Errorf("ent: MarketChartSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (mcs *MarketChartSelect) StringX(ctx context.Context) string {
	v, err := mcs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (mcs *MarketChartSelect) Ints(ctx context.Context) ([]int, error) {
	if len(mcs.fields) > 1 {
		return nil, errors.New("ent: MarketChartSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := mcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (mcs *MarketChartSelect) IntsX(ctx context.Context) []int {
	v, err := mcs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (mcs *MarketChartSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = mcs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{marketchart.Label}
	default:
		err = fmt.Errorf("ent: MarketChartSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (mcs *MarketChartSelect) IntX(ctx context.Context) int {
	v, err := mcs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (mcs *MarketChartSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(mcs.fields) > 1 {
		return nil, errors.New("ent: MarketChartSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := mcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (mcs *MarketChartSelect) Float64sX(ctx context.Context) []float64 {
	v, err := mcs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (mcs *MarketChartSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = mcs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{marketchart.Label}
	default:
		err = fmt.Errorf("ent: MarketChartSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (mcs *MarketChartSelect) Float64X(ctx context.Context) float64 {
	v, err := mcs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (mcs *MarketChartSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(mcs.fields) > 1 {
		return nil, errors.New("ent: MarketChartSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := mcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (mcs *MarketChartSelect) BoolsX(ctx context.Context) []bool {
	v, err := mcs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (mcs *MarketChartSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = mcs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{marketchart.Label}
	default:
		err = fmt.Errorf("ent: MarketChartSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (mcs *MarketChartSelect) BoolX(ctx context.Context) bool {
	v, err := mcs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mcs *MarketChartSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mcs.sql.Query()
	if err := mcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
