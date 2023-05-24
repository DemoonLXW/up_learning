// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/DemoonLXW/up_learning/database/ent/permission"
	"github.com/DemoonLXW/up_learning/database/ent/predicate"
	"github.com/DemoonLXW/up_learning/database/ent/role"
	"github.com/DemoonLXW/up_learning/database/ent/rolepermission"
)

// RolePermissionQuery is the builder for querying RolePermission entities.
type RolePermissionQuery struct {
	config
	ctx            *QueryContext
	order          []rolepermission.OrderOption
	inters         []Interceptor
	predicates     []predicate.RolePermission
	withRole       *RoleQuery
	withPermission *PermissionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RolePermissionQuery builder.
func (rpq *RolePermissionQuery) Where(ps ...predicate.RolePermission) *RolePermissionQuery {
	rpq.predicates = append(rpq.predicates, ps...)
	return rpq
}

// Limit the number of records to be returned by this query.
func (rpq *RolePermissionQuery) Limit(limit int) *RolePermissionQuery {
	rpq.ctx.Limit = &limit
	return rpq
}

// Offset to start from.
func (rpq *RolePermissionQuery) Offset(offset int) *RolePermissionQuery {
	rpq.ctx.Offset = &offset
	return rpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rpq *RolePermissionQuery) Unique(unique bool) *RolePermissionQuery {
	rpq.ctx.Unique = &unique
	return rpq
}

// Order specifies how the records should be ordered.
func (rpq *RolePermissionQuery) Order(o ...rolepermission.OrderOption) *RolePermissionQuery {
	rpq.order = append(rpq.order, o...)
	return rpq
}

// QueryRole chains the current query on the "role" edge.
func (rpq *RolePermissionQuery) QueryRole() *RoleQuery {
	query := (&RoleClient{config: rpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rolepermission.Table, rolepermission.RoleColumn, selector),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, rolepermission.RoleTable, rolepermission.RoleColumn),
		)
		fromU = sqlgraph.SetNeighbors(rpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPermission chains the current query on the "permission" edge.
func (rpq *RolePermissionQuery) QueryPermission() *PermissionQuery {
	query := (&PermissionClient{config: rpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(rolepermission.Table, rolepermission.PermissionColumn, selector),
			sqlgraph.To(permission.Table, permission.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, rolepermission.PermissionTable, rolepermission.PermissionColumn),
		)
		fromU = sqlgraph.SetNeighbors(rpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RolePermission entity from the query.
// Returns a *NotFoundError when no RolePermission was found.
func (rpq *RolePermissionQuery) First(ctx context.Context) (*RolePermission, error) {
	nodes, err := rpq.Limit(1).All(setContextOp(ctx, rpq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{rolepermission.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rpq *RolePermissionQuery) FirstX(ctx context.Context) *RolePermission {
	node, err := rpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// Only returns a single RolePermission entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RolePermission entity is found.
// Returns a *NotFoundError when no RolePermission entities are found.
func (rpq *RolePermissionQuery) Only(ctx context.Context) (*RolePermission, error) {
	nodes, err := rpq.Limit(2).All(setContextOp(ctx, rpq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{rolepermission.Label}
	default:
		return nil, &NotSingularError{rolepermission.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rpq *RolePermissionQuery) OnlyX(ctx context.Context) *RolePermission {
	node, err := rpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// All executes the query and returns a list of RolePermissions.
func (rpq *RolePermissionQuery) All(ctx context.Context) ([]*RolePermission, error) {
	ctx = setContextOp(ctx, rpq.ctx, "All")
	if err := rpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RolePermission, *RolePermissionQuery]()
	return withInterceptors[[]*RolePermission](ctx, rpq, qr, rpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rpq *RolePermissionQuery) AllX(ctx context.Context) []*RolePermission {
	nodes, err := rpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// Count returns the count of the given query.
func (rpq *RolePermissionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rpq.ctx, "Count")
	if err := rpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rpq, querierCount[*RolePermissionQuery](), rpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rpq *RolePermissionQuery) CountX(ctx context.Context) int {
	count, err := rpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rpq *RolePermissionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rpq.ctx, "Exist")
	switch _, err := rpq.First(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rpq *RolePermissionQuery) ExistX(ctx context.Context) bool {
	exist, err := rpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RolePermissionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rpq *RolePermissionQuery) Clone() *RolePermissionQuery {
	if rpq == nil {
		return nil
	}
	return &RolePermissionQuery{
		config:         rpq.config,
		ctx:            rpq.ctx.Clone(),
		order:          append([]rolepermission.OrderOption{}, rpq.order...),
		inters:         append([]Interceptor{}, rpq.inters...),
		predicates:     append([]predicate.RolePermission{}, rpq.predicates...),
		withRole:       rpq.withRole.Clone(),
		withPermission: rpq.withPermission.Clone(),
		// clone intermediate query.
		sql:  rpq.sql.Clone(),
		path: rpq.path,
	}
}

// WithRole tells the query-builder to eager-load the nodes that are connected to
// the "role" edge. The optional arguments are used to configure the query builder of the edge.
func (rpq *RolePermissionQuery) WithRole(opts ...func(*RoleQuery)) *RolePermissionQuery {
	query := (&RoleClient{config: rpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rpq.withRole = query
	return rpq
}

// WithPermission tells the query-builder to eager-load the nodes that are connected to
// the "permission" edge. The optional arguments are used to configure the query builder of the edge.
func (rpq *RolePermissionQuery) WithPermission(opts ...func(*PermissionQuery)) *RolePermissionQuery {
	query := (&PermissionClient{config: rpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rpq.withPermission = query
	return rpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Rid uint8 `json:"rid,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RolePermission.Query().
//		GroupBy(rolepermission.FieldRid).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rpq *RolePermissionQuery) GroupBy(field string, fields ...string) *RolePermissionGroupBy {
	rpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RolePermissionGroupBy{build: rpq}
	grbuild.flds = &rpq.ctx.Fields
	grbuild.label = rolepermission.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Rid uint8 `json:"rid,omitempty"`
//	}
//
//	client.RolePermission.Query().
//		Select(rolepermission.FieldRid).
//		Scan(ctx, &v)
func (rpq *RolePermissionQuery) Select(fields ...string) *RolePermissionSelect {
	rpq.ctx.Fields = append(rpq.ctx.Fields, fields...)
	sbuild := &RolePermissionSelect{RolePermissionQuery: rpq}
	sbuild.label = rolepermission.Label
	sbuild.flds, sbuild.scan = &rpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RolePermissionSelect configured with the given aggregations.
func (rpq *RolePermissionQuery) Aggregate(fns ...AggregateFunc) *RolePermissionSelect {
	return rpq.Select().Aggregate(fns...)
}

func (rpq *RolePermissionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rpq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rpq); err != nil {
				return err
			}
		}
	}
	for _, f := range rpq.ctx.Fields {
		if !rolepermission.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rpq.path != nil {
		prev, err := rpq.path(ctx)
		if err != nil {
			return err
		}
		rpq.sql = prev
	}
	return nil
}

func (rpq *RolePermissionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RolePermission, error) {
	var (
		nodes       = []*RolePermission{}
		_spec       = rpq.querySpec()
		loadedTypes = [2]bool{
			rpq.withRole != nil,
			rpq.withPermission != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RolePermission).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RolePermission{config: rpq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rpq.withRole; query != nil {
		if err := rpq.loadRole(ctx, query, nodes, nil,
			func(n *RolePermission, e *Role) { n.Edges.Role = e }); err != nil {
			return nil, err
		}
	}
	if query := rpq.withPermission; query != nil {
		if err := rpq.loadPermission(ctx, query, nodes, nil,
			func(n *RolePermission, e *Permission) { n.Edges.Permission = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rpq *RolePermissionQuery) loadRole(ctx context.Context, query *RoleQuery, nodes []*RolePermission, init func(*RolePermission), assign func(*RolePermission, *Role)) error {
	ids := make([]uint8, 0, len(nodes))
	nodeids := make(map[uint8][]*RolePermission)
	for i := range nodes {
		fk := nodes[i].Rid
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(role.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "rid" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (rpq *RolePermissionQuery) loadPermission(ctx context.Context, query *PermissionQuery, nodes []*RolePermission, init func(*RolePermission), assign func(*RolePermission, *Permission)) error {
	ids := make([]uint16, 0, len(nodes))
	nodeids := make(map[uint16][]*RolePermission)
	for i := range nodes {
		fk := nodes[i].Pid
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(permission.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "pid" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (rpq *RolePermissionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rpq.querySpec()
	_spec.Unique = false
	_spec.Node.Columns = nil
	return sqlgraph.CountNodes(ctx, rpq.driver, _spec)
}

func (rpq *RolePermissionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(rolepermission.Table, rolepermission.Columns, nil)
	_spec.From = rpq.sql
	if unique := rpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rpq.path != nil {
		_spec.Unique = true
	}
	if fields := rpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		for i := range fields {
			_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
		}
		if rpq.withRole != nil {
			_spec.Node.AddColumnOnce(rolepermission.FieldRid)
		}
		if rpq.withPermission != nil {
			_spec.Node.AddColumnOnce(rolepermission.FieldPid)
		}
	}
	if ps := rpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rpq *RolePermissionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rpq.driver.Dialect())
	t1 := builder.Table(rolepermission.Table)
	columns := rpq.ctx.Fields
	if len(columns) == 0 {
		columns = rolepermission.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rpq.sql != nil {
		selector = rpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rpq.ctx.Unique != nil && *rpq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rpq.predicates {
		p(selector)
	}
	for _, p := range rpq.order {
		p(selector)
	}
	if offset := rpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RolePermissionGroupBy is the group-by builder for RolePermission entities.
type RolePermissionGroupBy struct {
	selector
	build *RolePermissionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rpgb *RolePermissionGroupBy) Aggregate(fns ...AggregateFunc) *RolePermissionGroupBy {
	rpgb.fns = append(rpgb.fns, fns...)
	return rpgb
}

// Scan applies the selector query and scans the result into the given value.
func (rpgb *RolePermissionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rpgb.build.ctx, "GroupBy")
	if err := rpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RolePermissionQuery, *RolePermissionGroupBy](ctx, rpgb.build, rpgb, rpgb.build.inters, v)
}

func (rpgb *RolePermissionGroupBy) sqlScan(ctx context.Context, root *RolePermissionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rpgb.fns))
	for _, fn := range rpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rpgb.flds)+len(rpgb.fns))
		for _, f := range *rpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RolePermissionSelect is the builder for selecting fields of RolePermission entities.
type RolePermissionSelect struct {
	*RolePermissionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rps *RolePermissionSelect) Aggregate(fns ...AggregateFunc) *RolePermissionSelect {
	rps.fns = append(rps.fns, fns...)
	return rps
}

// Scan applies the selector query and scans the result into the given value.
func (rps *RolePermissionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rps.ctx, "Select")
	if err := rps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RolePermissionQuery, *RolePermissionSelect](ctx, rps.RolePermissionQuery, rps, rps.inters, v)
}

func (rps *RolePermissionSelect) sqlScan(ctx context.Context, root *RolePermissionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rps.fns))
	for _, fn := range rps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
