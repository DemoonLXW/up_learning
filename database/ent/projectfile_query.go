// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/DemoonLXW/up_learning/database/ent/file"
	"github.com/DemoonLXW/up_learning/database/ent/predicate"
	"github.com/DemoonLXW/up_learning/database/ent/project"
	"github.com/DemoonLXW/up_learning/database/ent/projectfile"
)

// ProjectFileQuery is the builder for querying ProjectFile entities.
type ProjectFileQuery struct {
	config
	ctx         *QueryContext
	order       []projectfile.OrderOption
	inters      []Interceptor
	predicates  []predicate.ProjectFile
	withProject *ProjectQuery
	withFiles   *FileQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProjectFileQuery builder.
func (pfq *ProjectFileQuery) Where(ps ...predicate.ProjectFile) *ProjectFileQuery {
	pfq.predicates = append(pfq.predicates, ps...)
	return pfq
}

// Limit the number of records to be returned by this query.
func (pfq *ProjectFileQuery) Limit(limit int) *ProjectFileQuery {
	pfq.ctx.Limit = &limit
	return pfq
}

// Offset to start from.
func (pfq *ProjectFileQuery) Offset(offset int) *ProjectFileQuery {
	pfq.ctx.Offset = &offset
	return pfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pfq *ProjectFileQuery) Unique(unique bool) *ProjectFileQuery {
	pfq.ctx.Unique = &unique
	return pfq
}

// Order specifies how the records should be ordered.
func (pfq *ProjectFileQuery) Order(o ...projectfile.OrderOption) *ProjectFileQuery {
	pfq.order = append(pfq.order, o...)
	return pfq
}

// QueryProject chains the current query on the "project" edge.
func (pfq *ProjectFileQuery) QueryProject() *ProjectQuery {
	query := (&ProjectClient{config: pfq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(projectfile.Table, projectfile.ProjectColumn, selector),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, projectfile.ProjectTable, projectfile.ProjectColumn),
		)
		fromU = sqlgraph.SetNeighbors(pfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFiles chains the current query on the "files" edge.
func (pfq *ProjectFileQuery) QueryFiles() *FileQuery {
	query := (&FileClient{config: pfq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(projectfile.Table, projectfile.FilesColumn, selector),
			sqlgraph.To(file.Table, file.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, projectfile.FilesTable, projectfile.FilesColumn),
		)
		fromU = sqlgraph.SetNeighbors(pfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProjectFile entity from the query.
// Returns a *NotFoundError when no ProjectFile was found.
func (pfq *ProjectFileQuery) First(ctx context.Context) (*ProjectFile, error) {
	nodes, err := pfq.Limit(1).All(setContextOp(ctx, pfq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{projectfile.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pfq *ProjectFileQuery) FirstX(ctx context.Context) *ProjectFile {
	node, err := pfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// Only returns a single ProjectFile entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ProjectFile entity is found.
// Returns a *NotFoundError when no ProjectFile entities are found.
func (pfq *ProjectFileQuery) Only(ctx context.Context) (*ProjectFile, error) {
	nodes, err := pfq.Limit(2).All(setContextOp(ctx, pfq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{projectfile.Label}
	default:
		return nil, &NotSingularError{projectfile.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pfq *ProjectFileQuery) OnlyX(ctx context.Context) *ProjectFile {
	node, err := pfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// All executes the query and returns a list of ProjectFiles.
func (pfq *ProjectFileQuery) All(ctx context.Context) ([]*ProjectFile, error) {
	ctx = setContextOp(ctx, pfq.ctx, "All")
	if err := pfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ProjectFile, *ProjectFileQuery]()
	return withInterceptors[[]*ProjectFile](ctx, pfq, qr, pfq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pfq *ProjectFileQuery) AllX(ctx context.Context) []*ProjectFile {
	nodes, err := pfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// Count returns the count of the given query.
func (pfq *ProjectFileQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pfq.ctx, "Count")
	if err := pfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pfq, querierCount[*ProjectFileQuery](), pfq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pfq *ProjectFileQuery) CountX(ctx context.Context) int {
	count, err := pfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pfq *ProjectFileQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pfq.ctx, "Exist")
	switch _, err := pfq.First(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pfq *ProjectFileQuery) ExistX(ctx context.Context) bool {
	exist, err := pfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProjectFileQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pfq *ProjectFileQuery) Clone() *ProjectFileQuery {
	if pfq == nil {
		return nil
	}
	return &ProjectFileQuery{
		config:      pfq.config,
		ctx:         pfq.ctx.Clone(),
		order:       append([]projectfile.OrderOption{}, pfq.order...),
		inters:      append([]Interceptor{}, pfq.inters...),
		predicates:  append([]predicate.ProjectFile{}, pfq.predicates...),
		withProject: pfq.withProject.Clone(),
		withFiles:   pfq.withFiles.Clone(),
		// clone intermediate query.
		sql:  pfq.sql.Clone(),
		path: pfq.path,
	}
}

// WithProject tells the query-builder to eager-load the nodes that are connected to
// the "project" edge. The optional arguments are used to configure the query builder of the edge.
func (pfq *ProjectFileQuery) WithProject(opts ...func(*ProjectQuery)) *ProjectFileQuery {
	query := (&ProjectClient{config: pfq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pfq.withProject = query
	return pfq
}

// WithFiles tells the query-builder to eager-load the nodes that are connected to
// the "files" edge. The optional arguments are used to configure the query builder of the edge.
func (pfq *ProjectFileQuery) WithFiles(opts ...func(*FileQuery)) *ProjectFileQuery {
	query := (&FileClient{config: pfq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pfq.withFiles = query
	return pfq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Pid uint32 `json:"pid,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ProjectFile.Query().
//		GroupBy(projectfile.FieldPid).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pfq *ProjectFileQuery) GroupBy(field string, fields ...string) *ProjectFileGroupBy {
	pfq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProjectFileGroupBy{build: pfq}
	grbuild.flds = &pfq.ctx.Fields
	grbuild.label = projectfile.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Pid uint32 `json:"pid,omitempty"`
//	}
//
//	client.ProjectFile.Query().
//		Select(projectfile.FieldPid).
//		Scan(ctx, &v)
func (pfq *ProjectFileQuery) Select(fields ...string) *ProjectFileSelect {
	pfq.ctx.Fields = append(pfq.ctx.Fields, fields...)
	sbuild := &ProjectFileSelect{ProjectFileQuery: pfq}
	sbuild.label = projectfile.Label
	sbuild.flds, sbuild.scan = &pfq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProjectFileSelect configured with the given aggregations.
func (pfq *ProjectFileQuery) Aggregate(fns ...AggregateFunc) *ProjectFileSelect {
	return pfq.Select().Aggregate(fns...)
}

func (pfq *ProjectFileQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pfq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pfq); err != nil {
				return err
			}
		}
	}
	for _, f := range pfq.ctx.Fields {
		if !projectfile.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pfq.path != nil {
		prev, err := pfq.path(ctx)
		if err != nil {
			return err
		}
		pfq.sql = prev
	}
	return nil
}

func (pfq *ProjectFileQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ProjectFile, error) {
	var (
		nodes       = []*ProjectFile{}
		_spec       = pfq.querySpec()
		loadedTypes = [2]bool{
			pfq.withProject != nil,
			pfq.withFiles != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ProjectFile).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ProjectFile{config: pfq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pfq.withProject; query != nil {
		if err := pfq.loadProject(ctx, query, nodes, nil,
			func(n *ProjectFile, e *Project) { n.Edges.Project = e }); err != nil {
			return nil, err
		}
	}
	if query := pfq.withFiles; query != nil {
		if err := pfq.loadFiles(ctx, query, nodes, nil,
			func(n *ProjectFile, e *File) { n.Edges.Files = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pfq *ProjectFileQuery) loadProject(ctx context.Context, query *ProjectQuery, nodes []*ProjectFile, init func(*ProjectFile), assign func(*ProjectFile, *Project)) error {
	ids := make([]uint32, 0, len(nodes))
	nodeids := make(map[uint32][]*ProjectFile)
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
	query.Where(project.IDIn(ids...))
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
func (pfq *ProjectFileQuery) loadFiles(ctx context.Context, query *FileQuery, nodes []*ProjectFile, init func(*ProjectFile), assign func(*ProjectFile, *File)) error {
	ids := make([]uint32, 0, len(nodes))
	nodeids := make(map[uint32][]*ProjectFile)
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

func (pfq *ProjectFileQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pfq.querySpec()
	_spec.Unique = false
	_spec.Node.Columns = nil
	return sqlgraph.CountNodes(ctx, pfq.driver, _spec)
}

func (pfq *ProjectFileQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(projectfile.Table, projectfile.Columns, nil)
	_spec.From = pfq.sql
	if unique := pfq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pfq.path != nil {
		_spec.Unique = true
	}
	if fields := pfq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		for i := range fields {
			_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
		}
		if pfq.withProject != nil {
			_spec.Node.AddColumnOnce(projectfile.FieldPid)
		}
		if pfq.withFiles != nil {
			_spec.Node.AddColumnOnce(projectfile.FieldFid)
		}
	}
	if ps := pfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pfq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pfq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pfq *ProjectFileQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pfq.driver.Dialect())
	t1 := builder.Table(projectfile.Table)
	columns := pfq.ctx.Fields
	if len(columns) == 0 {
		columns = projectfile.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pfq.sql != nil {
		selector = pfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pfq.ctx.Unique != nil && *pfq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pfq.predicates {
		p(selector)
	}
	for _, p := range pfq.order {
		p(selector)
	}
	if offset := pfq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pfq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProjectFileGroupBy is the group-by builder for ProjectFile entities.
type ProjectFileGroupBy struct {
	selector
	build *ProjectFileQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pfgb *ProjectFileGroupBy) Aggregate(fns ...AggregateFunc) *ProjectFileGroupBy {
	pfgb.fns = append(pfgb.fns, fns...)
	return pfgb
}

// Scan applies the selector query and scans the result into the given value.
func (pfgb *ProjectFileGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pfgb.build.ctx, "GroupBy")
	if err := pfgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProjectFileQuery, *ProjectFileGroupBy](ctx, pfgb.build, pfgb, pfgb.build.inters, v)
}

func (pfgb *ProjectFileGroupBy) sqlScan(ctx context.Context, root *ProjectFileQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pfgb.fns))
	for _, fn := range pfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pfgb.flds)+len(pfgb.fns))
		for _, f := range *pfgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pfgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pfgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProjectFileSelect is the builder for selecting fields of ProjectFile entities.
type ProjectFileSelect struct {
	*ProjectFileQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pfs *ProjectFileSelect) Aggregate(fns ...AggregateFunc) *ProjectFileSelect {
	pfs.fns = append(pfs.fns, fns...)
	return pfs
}

// Scan applies the selector query and scans the result into the given value.
func (pfs *ProjectFileSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pfs.ctx, "Select")
	if err := pfs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProjectFileQuery, *ProjectFileSelect](ctx, pfs.ProjectFileQuery, pfs, pfs.inters, v)
}

func (pfs *ProjectFileSelect) sqlScan(ctx context.Context, root *ProjectFileQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pfs.fns))
	for _, fn := range pfs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pfs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
