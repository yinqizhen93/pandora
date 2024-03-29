// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"
	"pandora/ent/department"
	"pandora/ent/predicate"
	"pandora/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DepartmentQuery is the builder for querying Department entities.
type DepartmentQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Department
	// eager-loading edges.
	withUsers     *UserQuery
	withParent1   *DepartmentQuery
	withChildren1 *DepartmentQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DepartmentQuery builder.
func (dq *DepartmentQuery) Where(ps ...predicate.Department) *DepartmentQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit adds a limit step to the query.
func (dq *DepartmentQuery) Limit(limit int) *DepartmentQuery {
	dq.limit = &limit
	return dq
}

// Offset adds an offset step to the query.
func (dq *DepartmentQuery) Offset(offset int) *DepartmentQuery {
	dq.offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DepartmentQuery) Unique(unique bool) *DepartmentQuery {
	dq.unique = &unique
	return dq
}

// Order adds an order step to the query.
func (dq *DepartmentQuery) Order(o ...OrderFunc) *DepartmentQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryUsers chains the current query on the "users" edge.
func (dq *DepartmentQuery) QueryUsers() *UserQuery {
	query := &UserQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(department.Table, department.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, department.UsersTable, department.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryParent1 chains the current query on the "parent1" edge.
func (dq *DepartmentQuery) QueryParent1() *DepartmentQuery {
	query := &DepartmentQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(department.Table, department.FieldID, selector),
			sqlgraph.To(department.Table, department.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, department.Parent1Table, department.Parent1Column),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren1 chains the current query on the "children1" edge.
func (dq *DepartmentQuery) QueryChildren1() *DepartmentQuery {
	query := &DepartmentQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(department.Table, department.FieldID, selector),
			sqlgraph.To(department.Table, department.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, department.Children1Table, department.Children1Column),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Department entity from the query.
// Returns a *NotFoundError when no Department was found.
func (dq *DepartmentQuery) First(ctx context.Context) (*Department, error) {
	nodes, err := dq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{department.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DepartmentQuery) FirstX(ctx context.Context) *Department {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Department ID from the query.
// Returns a *NotFoundError when no Department ID was found.
func (dq *DepartmentQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{department.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DepartmentQuery) FirstIDX(ctx context.Context) int {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Department entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Department entity is found.
// Returns a *NotFoundError when no Department entities are found.
func (dq *DepartmentQuery) Only(ctx context.Context) (*Department, error) {
	nodes, err := dq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{department.Label}
	default:
		return nil, &NotSingularError{department.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DepartmentQuery) OnlyX(ctx context.Context) *Department {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Department ID in the query.
// Returns a *NotSingularError when more than one Department ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DepartmentQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{department.Label}
	default:
		err = &NotSingularError{department.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DepartmentQuery) OnlyIDX(ctx context.Context) int {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Departments.
func (dq *DepartmentQuery) All(ctx context.Context) ([]*Department, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dq *DepartmentQuery) AllX(ctx context.Context) []*Department {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Department IDs.
func (dq *DepartmentQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := dq.Select(department.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DepartmentQuery) IDsX(ctx context.Context) []int {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DepartmentQuery) Count(ctx context.Context) (int, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DepartmentQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DepartmentQuery) Exist(ctx context.Context) (bool, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DepartmentQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DepartmentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DepartmentQuery) Clone() *DepartmentQuery {
	if dq == nil {
		return nil
	}
	return &DepartmentQuery{
		config:        dq.config,
		limit:         dq.limit,
		offset:        dq.offset,
		order:         append([]OrderFunc{}, dq.order...),
		predicates:    append([]predicate.Department{}, dq.predicates...),
		withUsers:     dq.withUsers.Clone(),
		withParent1:   dq.withParent1.Clone(),
		withChildren1: dq.withChildren1.Clone(),
		// clone intermediate query.
		sql:    dq.sql.Clone(),
		path:   dq.path,
		unique: dq.unique,
	}
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DepartmentQuery) WithUsers(opts ...func(*UserQuery)) *DepartmentQuery {
	query := &UserQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withUsers = query
	return dq
}

// WithParent1 tells the query-builder to eager-load the nodes that are connected to
// the "parent1" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DepartmentQuery) WithParent1(opts ...func(*DepartmentQuery)) *DepartmentQuery {
	query := &DepartmentQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withParent1 = query
	return dq
}

// WithChildren1 tells the query-builder to eager-load the nodes that are connected to
// the "children1" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DepartmentQuery) WithChildren1(opts ...func(*DepartmentQuery)) *DepartmentQuery {
	query := &DepartmentQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withChildren1 = query
	return dq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Code string `json:"code,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Department.Query().
//		GroupBy(department.FieldCode).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (dq *DepartmentQuery) GroupBy(field string, fields ...string) *DepartmentGroupBy {
	group := &DepartmentGroupBy{config: dq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Code string `json:"code,omitempty"`
//	}
//
//	client.Department.Query().
//		Select(department.FieldCode).
//		Scan(ctx, &v)
//
func (dq *DepartmentQuery) Select(fields ...string) *DepartmentSelect {
	dq.fields = append(dq.fields, fields...)
	return &DepartmentSelect{DepartmentQuery: dq}
}

func (dq *DepartmentQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dq.fields {
		if !department.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DepartmentQuery) sqlAll(ctx context.Context) ([]*Department, error) {
	var (
		nodes       = []*Department{}
		withFKs     = dq.withFKs
		_spec       = dq.querySpec()
		loadedTypes = [3]bool{
			dq.withUsers != nil,
			dq.withParent1 != nil,
			dq.withChildren1 != nil,
		}
	)
	if dq.withParent1 != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, department.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Department{config: dq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := dq.withUsers; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Department)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Users = []*User{}
		}
		query.withFKs = true
		query.Where(predicate.User(func(s *sql.Selector) {
			s.Where(sql.InValues(department.UsersColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.user_department
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "user_department" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_department" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Users = append(node.Edges.Users, n)
		}
	}

	if query := dq.withParent1; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Department)
		for i := range nodes {
			if nodes[i].department_children1 == nil {
				continue
			}
			fk := *nodes[i].department_children1
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(department.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "department_children1" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parent1 = n
			}
		}
	}

	if query := dq.withChildren1; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Department)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Children1 = []*Department{}
		}
		query.withFKs = true
		query.Where(predicate.Department(func(s *sql.Selector) {
			s.Where(sql.InValues(department.Children1Column, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.department_children1
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "department_children1" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "department_children1" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Children1 = append(node.Edges.Children1, n)
		}
	}

	return nodes, nil
}

func (dq *DepartmentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	_spec.Node.Columns = dq.fields
	if len(dq.fields) > 0 {
		_spec.Unique = dq.unique != nil && *dq.unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DepartmentQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (dq *DepartmentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   department.Table,
			Columns: department.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: department.FieldID,
			},
		},
		From:   dq.sql,
		Unique: true,
	}
	if unique := dq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, department.FieldID)
		for i := range fields {
			if fields[i] != department.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DepartmentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(department.Table)
	columns := dq.fields
	if len(columns) == 0 {
		columns = department.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.unique != nil && *dq.unique {
		selector.Distinct()
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

// DepartmentGroupBy is the group-by builder for Department entities.
type DepartmentGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DepartmentGroupBy) Aggregate(fns ...AggregateFunc) *DepartmentGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dgb *DepartmentGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dgb.path(ctx)
	if err != nil {
		return err
	}
	dgb.sql = query
	return dgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dgb *DepartmentGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := dgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (dgb *DepartmentGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DepartmentGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dgb *DepartmentGroupBy) StringsX(ctx context.Context) []string {
	v, err := dgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dgb *DepartmentGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{department.Label}
	default:
		err = fmt.Errorf("ent: DepartmentGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dgb *DepartmentGroupBy) StringX(ctx context.Context) string {
	v, err := dgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (dgb *DepartmentGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DepartmentGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dgb *DepartmentGroupBy) IntsX(ctx context.Context) []int {
	v, err := dgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dgb *DepartmentGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{department.Label}
	default:
		err = fmt.Errorf("ent: DepartmentGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dgb *DepartmentGroupBy) IntX(ctx context.Context) int {
	v, err := dgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (dgb *DepartmentGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DepartmentGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dgb *DepartmentGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := dgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dgb *DepartmentGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{department.Label}
	default:
		err = fmt.Errorf("ent: DepartmentGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dgb *DepartmentGroupBy) Float64X(ctx context.Context) float64 {
	v, err := dgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (dgb *DepartmentGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DepartmentGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dgb *DepartmentGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := dgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (dgb *DepartmentGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{department.Label}
	default:
		err = fmt.Errorf("ent: DepartmentGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dgb *DepartmentGroupBy) BoolX(ctx context.Context) bool {
	v, err := dgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dgb *DepartmentGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dgb.fields {
		if !department.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dgb *DepartmentGroupBy) sqlQuery() *sql.Selector {
	selector := dgb.sql.Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dgb.fields)+len(dgb.fns))
		for _, f := range dgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dgb.fields...)...)
}

// DepartmentSelect is the builder for selecting fields of Department entities.
type DepartmentSelect struct {
	*DepartmentQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DepartmentSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	ds.sql = ds.DepartmentQuery.sqlQuery(ctx)
	return ds.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ds *DepartmentSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ds.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ds *DepartmentSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DepartmentSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ds *DepartmentSelect) StringsX(ctx context.Context) []string {
	v, err := ds.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ds *DepartmentSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ds.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{department.Label}
	default:
		err = fmt.Errorf("ent: DepartmentSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ds *DepartmentSelect) StringX(ctx context.Context) string {
	v, err := ds.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ds *DepartmentSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DepartmentSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ds *DepartmentSelect) IntsX(ctx context.Context) []int {
	v, err := ds.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ds *DepartmentSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ds.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{department.Label}
	default:
		err = fmt.Errorf("ent: DepartmentSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ds *DepartmentSelect) IntX(ctx context.Context) int {
	v, err := ds.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ds *DepartmentSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DepartmentSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ds *DepartmentSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ds.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ds *DepartmentSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ds.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{department.Label}
	default:
		err = fmt.Errorf("ent: DepartmentSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ds *DepartmentSelect) Float64X(ctx context.Context) float64 {
	v, err := ds.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ds *DepartmentSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DepartmentSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ds *DepartmentSelect) BoolsX(ctx context.Context) []bool {
	v, err := ds.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ds *DepartmentSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ds.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{department.Label}
	default:
		err = fmt.Errorf("ent: DepartmentSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ds *DepartmentSelect) BoolX(ctx context.Context) bool {
	v, err := ds.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ds *DepartmentSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ds.sql.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
