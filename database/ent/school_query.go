// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DemoonLXW/up_learning/database/ent/predicate"
	"github.com/DemoonLXW/up_learning/database/ent/school"
	"github.com/DemoonLXW/up_learning/database/ent/student"
)

// SchoolQuery is the builder for querying School entities.
type SchoolQuery struct {
	config
	ctx          *QueryContext
	order        []school.OrderOption
	inters       []Interceptor
	predicates   []predicate.School
	withStudents *StudentQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SchoolQuery builder.
func (sq *SchoolQuery) Where(ps ...predicate.School) *SchoolQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *SchoolQuery) Limit(limit int) *SchoolQuery {
	sq.ctx.Limit = &limit
	return sq
}

// Offset to start from.
func (sq *SchoolQuery) Offset(offset int) *SchoolQuery {
	sq.ctx.Offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SchoolQuery) Unique(unique bool) *SchoolQuery {
	sq.ctx.Unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *SchoolQuery) Order(o ...school.OrderOption) *SchoolQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryStudents chains the current query on the "students" edge.
func (sq *SchoolQuery) QueryStudents() *StudentQuery {
	query := (&StudentClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(school.Table, school.FieldID, selector),
			sqlgraph.To(student.Table, student.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, school.StudentsTable, school.StudentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first School entity from the query.
// Returns a *NotFoundError when no School was found.
func (sq *SchoolQuery) First(ctx context.Context) (*School, error) {
	nodes, err := sq.Limit(1).All(setContextOp(ctx, sq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{school.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SchoolQuery) FirstX(ctx context.Context) *School {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first School ID from the query.
// Returns a *NotFoundError when no School ID was found.
func (sq *SchoolQuery) FirstID(ctx context.Context) (id uint16, err error) {
	var ids []uint16
	if ids, err = sq.Limit(1).IDs(setContextOp(ctx, sq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{school.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SchoolQuery) FirstIDX(ctx context.Context) uint16 {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single School entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one School entity is found.
// Returns a *NotFoundError when no School entities are found.
func (sq *SchoolQuery) Only(ctx context.Context) (*School, error) {
	nodes, err := sq.Limit(2).All(setContextOp(ctx, sq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{school.Label}
	default:
		return nil, &NotSingularError{school.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SchoolQuery) OnlyX(ctx context.Context) *School {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only School ID in the query.
// Returns a *NotSingularError when more than one School ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SchoolQuery) OnlyID(ctx context.Context) (id uint16, err error) {
	var ids []uint16
	if ids, err = sq.Limit(2).IDs(setContextOp(ctx, sq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{school.Label}
	default:
		err = &NotSingularError{school.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SchoolQuery) OnlyIDX(ctx context.Context) uint16 {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Schools.
func (sq *SchoolQuery) All(ctx context.Context) ([]*School, error) {
	ctx = setContextOp(ctx, sq.ctx, "All")
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*School, *SchoolQuery]()
	return withInterceptors[[]*School](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *SchoolQuery) AllX(ctx context.Context) []*School {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of School IDs.
func (sq *SchoolQuery) IDs(ctx context.Context) (ids []uint16, err error) {
	if sq.ctx.Unique == nil && sq.path != nil {
		sq.Unique(true)
	}
	ctx = setContextOp(ctx, sq.ctx, "IDs")
	if err = sq.Select(school.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SchoolQuery) IDsX(ctx context.Context) []uint16 {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SchoolQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sq.ctx, "Count")
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*SchoolQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SchoolQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SchoolQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sq.ctx, "Exist")
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SchoolQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SchoolQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SchoolQuery) Clone() *SchoolQuery {
	if sq == nil {
		return nil
	}
	return &SchoolQuery{
		config:       sq.config,
		ctx:          sq.ctx.Clone(),
		order:        append([]school.OrderOption{}, sq.order...),
		inters:       append([]Interceptor{}, sq.inters...),
		predicates:   append([]predicate.School{}, sq.predicates...),
		withStudents: sq.withStudents.Clone(),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

// WithStudents tells the query-builder to eager-load the nodes that are connected to
// the "students" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SchoolQuery) WithStudents(opts ...func(*StudentQuery)) *SchoolQuery {
	query := (&StudentClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withStudents = query
	return sq
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
//	client.School.Query().
//		GroupBy(school.FieldCode).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *SchoolQuery) GroupBy(field string, fields ...string) *SchoolGroupBy {
	sq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SchoolGroupBy{build: sq}
	grbuild.flds = &sq.ctx.Fields
	grbuild.label = school.Label
	grbuild.scan = grbuild.Scan
	return grbuild
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
//	client.School.Query().
//		Select(school.FieldCode).
//		Scan(ctx, &v)
func (sq *SchoolQuery) Select(fields ...string) *SchoolSelect {
	sq.ctx.Fields = append(sq.ctx.Fields, fields...)
	sbuild := &SchoolSelect{SchoolQuery: sq}
	sbuild.label = school.Label
	sbuild.flds, sbuild.scan = &sq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SchoolSelect configured with the given aggregations.
func (sq *SchoolQuery) Aggregate(fns ...AggregateFunc) *SchoolSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *SchoolQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	for _, f := range sq.ctx.Fields {
		if !school.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *SchoolQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*School, error) {
	var (
		nodes       = []*School{}
		_spec       = sq.querySpec()
		loadedTypes = [1]bool{
			sq.withStudents != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*School).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &School{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withStudents; query != nil {
		if err := sq.loadStudents(ctx, query, nodes,
			func(n *School) { n.Edges.Students = []*Student{} },
			func(n *School, e *Student) { n.Edges.Students = append(n.Edges.Students, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *SchoolQuery) loadStudents(ctx context.Context, query *StudentQuery, nodes []*School, init func(*School), assign func(*School, *Student)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint16]*School)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Student(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(school.StudentsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.school_students
		if fk == nil {
			return fmt.Errorf(`foreign-key "school_students" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "school_students" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (sq *SchoolQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	_spec.Node.Columns = sq.ctx.Fields
	if len(sq.ctx.Fields) > 0 {
		_spec.Unique = sq.ctx.Unique != nil && *sq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SchoolQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(school.Table, school.Columns, sqlgraph.NewFieldSpec(school.FieldID, field.TypeUint16))
	_spec.From = sq.sql
	if unique := sq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sq.path != nil {
		_spec.Unique = true
	}
	if fields := sq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, school.FieldID)
		for i := range fields {
			if fields[i] != school.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SchoolQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(school.Table)
	columns := sq.ctx.Fields
	if len(columns) == 0 {
		columns = school.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.ctx.Unique != nil && *sq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SchoolGroupBy is the group-by builder for School entities.
type SchoolGroupBy struct {
	selector
	build *SchoolQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SchoolGroupBy) Aggregate(fns ...AggregateFunc) *SchoolGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *SchoolGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, "GroupBy")
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SchoolQuery, *SchoolGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *SchoolGroupBy) sqlScan(ctx context.Context, root *SchoolQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sgb.flds)+len(sgb.fns))
		for _, f := range *sgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SchoolSelect is the builder for selecting fields of School entities.
type SchoolSelect struct {
	*SchoolQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *SchoolSelect) Aggregate(fns ...AggregateFunc) *SchoolSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SchoolSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, "Select")
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SchoolQuery, *SchoolSelect](ctx, ss.SchoolQuery, ss, ss.inters, v)
}

func (ss *SchoolSelect) sqlScan(ctx context.Context, root *SchoolQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
