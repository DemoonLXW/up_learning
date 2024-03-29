// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DemoonLXW/up_learning/database/ent/file"
	"github.com/DemoonLXW/up_learning/database/ent/predicate"
	"github.com/DemoonLXW/up_learning/database/ent/samplefile"
)

// SampleFileQuery is the builder for querying SampleFile entities.
type SampleFileQuery struct {
	config
	ctx        *QueryContext
	order      []samplefile.OrderOption
	inters     []Interceptor
	predicates []predicate.SampleFile
	withFile   *FileQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SampleFileQuery builder.
func (sfq *SampleFileQuery) Where(ps ...predicate.SampleFile) *SampleFileQuery {
	sfq.predicates = append(sfq.predicates, ps...)
	return sfq
}

// Limit the number of records to be returned by this query.
func (sfq *SampleFileQuery) Limit(limit int) *SampleFileQuery {
	sfq.ctx.Limit = &limit
	return sfq
}

// Offset to start from.
func (sfq *SampleFileQuery) Offset(offset int) *SampleFileQuery {
	sfq.ctx.Offset = &offset
	return sfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sfq *SampleFileQuery) Unique(unique bool) *SampleFileQuery {
	sfq.ctx.Unique = &unique
	return sfq
}

// Order specifies how the records should be ordered.
func (sfq *SampleFileQuery) Order(o ...samplefile.OrderOption) *SampleFileQuery {
	sfq.order = append(sfq.order, o...)
	return sfq
}

// QueryFile chains the current query on the "file" edge.
func (sfq *SampleFileQuery) QueryFile() *FileQuery {
	query := (&FileClient{config: sfq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(samplefile.Table, samplefile.FieldID, selector),
			sqlgraph.To(file.Table, file.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, samplefile.FileTable, samplefile.FileColumn),
		)
		fromU = sqlgraph.SetNeighbors(sfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SampleFile entity from the query.
// Returns a *NotFoundError when no SampleFile was found.
func (sfq *SampleFileQuery) First(ctx context.Context) (*SampleFile, error) {
	nodes, err := sfq.Limit(1).All(setContextOp(ctx, sfq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{samplefile.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sfq *SampleFileQuery) FirstX(ctx context.Context) *SampleFile {
	node, err := sfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SampleFile ID from the query.
// Returns a *NotFoundError when no SampleFile ID was found.
func (sfq *SampleFileQuery) FirstID(ctx context.Context) (id uint8, err error) {
	var ids []uint8
	if ids, err = sfq.Limit(1).IDs(setContextOp(ctx, sfq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{samplefile.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sfq *SampleFileQuery) FirstIDX(ctx context.Context) uint8 {
	id, err := sfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SampleFile entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SampleFile entity is found.
// Returns a *NotFoundError when no SampleFile entities are found.
func (sfq *SampleFileQuery) Only(ctx context.Context) (*SampleFile, error) {
	nodes, err := sfq.Limit(2).All(setContextOp(ctx, sfq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{samplefile.Label}
	default:
		return nil, &NotSingularError{samplefile.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sfq *SampleFileQuery) OnlyX(ctx context.Context) *SampleFile {
	node, err := sfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SampleFile ID in the query.
// Returns a *NotSingularError when more than one SampleFile ID is found.
// Returns a *NotFoundError when no entities are found.
func (sfq *SampleFileQuery) OnlyID(ctx context.Context) (id uint8, err error) {
	var ids []uint8
	if ids, err = sfq.Limit(2).IDs(setContextOp(ctx, sfq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{samplefile.Label}
	default:
		err = &NotSingularError{samplefile.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sfq *SampleFileQuery) OnlyIDX(ctx context.Context) uint8 {
	id, err := sfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SampleFiles.
func (sfq *SampleFileQuery) All(ctx context.Context) ([]*SampleFile, error) {
	ctx = setContextOp(ctx, sfq.ctx, "All")
	if err := sfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SampleFile, *SampleFileQuery]()
	return withInterceptors[[]*SampleFile](ctx, sfq, qr, sfq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sfq *SampleFileQuery) AllX(ctx context.Context) []*SampleFile {
	nodes, err := sfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SampleFile IDs.
func (sfq *SampleFileQuery) IDs(ctx context.Context) (ids []uint8, err error) {
	if sfq.ctx.Unique == nil && sfq.path != nil {
		sfq.Unique(true)
	}
	ctx = setContextOp(ctx, sfq.ctx, "IDs")
	if err = sfq.Select(samplefile.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sfq *SampleFileQuery) IDsX(ctx context.Context) []uint8 {
	ids, err := sfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sfq *SampleFileQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sfq.ctx, "Count")
	if err := sfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sfq, querierCount[*SampleFileQuery](), sfq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sfq *SampleFileQuery) CountX(ctx context.Context) int {
	count, err := sfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sfq *SampleFileQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sfq.ctx, "Exist")
	switch _, err := sfq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sfq *SampleFileQuery) ExistX(ctx context.Context) bool {
	exist, err := sfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SampleFileQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sfq *SampleFileQuery) Clone() *SampleFileQuery {
	if sfq == nil {
		return nil
	}
	return &SampleFileQuery{
		config:     sfq.config,
		ctx:        sfq.ctx.Clone(),
		order:      append([]samplefile.OrderOption{}, sfq.order...),
		inters:     append([]Interceptor{}, sfq.inters...),
		predicates: append([]predicate.SampleFile{}, sfq.predicates...),
		withFile:   sfq.withFile.Clone(),
		// clone intermediate query.
		sql:  sfq.sql.Clone(),
		path: sfq.path,
	}
}

// WithFile tells the query-builder to eager-load the nodes that are connected to
// the "file" edge. The optional arguments are used to configure the query builder of the edge.
func (sfq *SampleFileQuery) WithFile(opts ...func(*FileQuery)) *SampleFileQuery {
	query := (&FileClient{config: sfq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sfq.withFile = query
	return sfq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Fid uint32 `json:"fid,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SampleFile.Query().
//		GroupBy(samplefile.FieldFid).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sfq *SampleFileQuery) GroupBy(field string, fields ...string) *SampleFileGroupBy {
	sfq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SampleFileGroupBy{build: sfq}
	grbuild.flds = &sfq.ctx.Fields
	grbuild.label = samplefile.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Fid uint32 `json:"fid,omitempty"`
//	}
//
//	client.SampleFile.Query().
//		Select(samplefile.FieldFid).
//		Scan(ctx, &v)
func (sfq *SampleFileQuery) Select(fields ...string) *SampleFileSelect {
	sfq.ctx.Fields = append(sfq.ctx.Fields, fields...)
	sbuild := &SampleFileSelect{SampleFileQuery: sfq}
	sbuild.label = samplefile.Label
	sbuild.flds, sbuild.scan = &sfq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SampleFileSelect configured with the given aggregations.
func (sfq *SampleFileQuery) Aggregate(fns ...AggregateFunc) *SampleFileSelect {
	return sfq.Select().Aggregate(fns...)
}

func (sfq *SampleFileQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sfq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sfq); err != nil {
				return err
			}
		}
	}
	for _, f := range sfq.ctx.Fields {
		if !samplefile.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sfq.path != nil {
		prev, err := sfq.path(ctx)
		if err != nil {
			return err
		}
		sfq.sql = prev
	}
	return nil
}

func (sfq *SampleFileQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SampleFile, error) {
	var (
		nodes       = []*SampleFile{}
		_spec       = sfq.querySpec()
		loadedTypes = [1]bool{
			sfq.withFile != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SampleFile).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SampleFile{config: sfq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sfq.withFile; query != nil {
		if err := sfq.loadFile(ctx, query, nodes, nil,
			func(n *SampleFile, e *File) { n.Edges.File = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sfq *SampleFileQuery) loadFile(ctx context.Context, query *FileQuery, nodes []*SampleFile, init func(*SampleFile), assign func(*SampleFile, *File)) error {
	ids := make([]uint32, 0, len(nodes))
	nodeids := make(map[uint32][]*SampleFile)
	for i := range nodes {
		fk := nodes[i].Fid
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(file.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "fid" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (sfq *SampleFileQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sfq.querySpec()
	_spec.Node.Columns = sfq.ctx.Fields
	if len(sfq.ctx.Fields) > 0 {
		_spec.Unique = sfq.ctx.Unique != nil && *sfq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sfq.driver, _spec)
}

func (sfq *SampleFileQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(samplefile.Table, samplefile.Columns, sqlgraph.NewFieldSpec(samplefile.FieldID, field.TypeUint8))
	_spec.From = sfq.sql
	if unique := sfq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sfq.path != nil {
		_spec.Unique = true
	}
	if fields := sfq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, samplefile.FieldID)
		for i := range fields {
			if fields[i] != samplefile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if sfq.withFile != nil {
			_spec.Node.AddColumnOnce(samplefile.FieldFid)
		}
	}
	if ps := sfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sfq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sfq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sfq *SampleFileQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sfq.driver.Dialect())
	t1 := builder.Table(samplefile.Table)
	columns := sfq.ctx.Fields
	if len(columns) == 0 {
		columns = samplefile.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sfq.sql != nil {
		selector = sfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sfq.ctx.Unique != nil && *sfq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sfq.predicates {
		p(selector)
	}
	for _, p := range sfq.order {
		p(selector)
	}
	if offset := sfq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sfq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SampleFileGroupBy is the group-by builder for SampleFile entities.
type SampleFileGroupBy struct {
	selector
	build *SampleFileQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sfgb *SampleFileGroupBy) Aggregate(fns ...AggregateFunc) *SampleFileGroupBy {
	sfgb.fns = append(sfgb.fns, fns...)
	return sfgb
}

// Scan applies the selector query and scans the result into the given value.
func (sfgb *SampleFileGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sfgb.build.ctx, "GroupBy")
	if err := sfgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SampleFileQuery, *SampleFileGroupBy](ctx, sfgb.build, sfgb, sfgb.build.inters, v)
}

func (sfgb *SampleFileGroupBy) sqlScan(ctx context.Context, root *SampleFileQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sfgb.fns))
	for _, fn := range sfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sfgb.flds)+len(sfgb.fns))
		for _, f := range *sfgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sfgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sfgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SampleFileSelect is the builder for selecting fields of SampleFile entities.
type SampleFileSelect struct {
	*SampleFileQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sfs *SampleFileSelect) Aggregate(fns ...AggregateFunc) *SampleFileSelect {
	sfs.fns = append(sfs.fns, fns...)
	return sfs
}

// Scan applies the selector query and scans the result into the given value.
func (sfs *SampleFileSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sfs.ctx, "Select")
	if err := sfs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SampleFileQuery, *SampleFileSelect](ctx, sfs.SampleFileQuery, sfs, sfs.inters, v)
}

func (sfs *SampleFileSelect) sqlScan(ctx context.Context, root *SampleFileQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sfs.fns))
	for _, fn := range sfs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sfs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
