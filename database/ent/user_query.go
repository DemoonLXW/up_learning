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
	"github.com/DemoonLXW/up_learning/database/ent/file"
	"github.com/DemoonLXW/up_learning/database/ent/predicate"
	"github.com/DemoonLXW/up_learning/database/ent/project"
	"github.com/DemoonLXW/up_learning/database/ent/role"
	"github.com/DemoonLXW/up_learning/database/ent/student"
	"github.com/DemoonLXW/up_learning/database/ent/teacher"
	"github.com/DemoonLXW/up_learning/database/ent/user"
	"github.com/DemoonLXW/up_learning/database/ent/userrole"
)

// UserQuery is the builder for querying User entities.
type UserQuery struct {
	config
	ctx          *QueryContext
	order        []user.OrderOption
	inters       []Interceptor
	predicates   []predicate.User
	withRoles    *RoleQuery
	withStudent  *StudentQuery
	withTeacher  *TeacherQuery
	withFiles    *FileQuery
	withProjects *ProjectQuery
	withUserRole *UserRoleQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserQuery builder.
func (uq *UserQuery) Where(ps ...predicate.User) *UserQuery {
	uq.predicates = append(uq.predicates, ps...)
	return uq
}

// Limit the number of records to be returned by this query.
func (uq *UserQuery) Limit(limit int) *UserQuery {
	uq.ctx.Limit = &limit
	return uq
}

// Offset to start from.
func (uq *UserQuery) Offset(offset int) *UserQuery {
	uq.ctx.Offset = &offset
	return uq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uq *UserQuery) Unique(unique bool) *UserQuery {
	uq.ctx.Unique = &unique
	return uq
}

// Order specifies how the records should be ordered.
func (uq *UserQuery) Order(o ...user.OrderOption) *UserQuery {
	uq.order = append(uq.order, o...)
	return uq
}

// QueryRoles chains the current query on the "roles" edge.
func (uq *UserQuery) QueryRoles() *RoleQuery {
	query := (&RoleClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.RolesTable, user.RolesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryStudent chains the current query on the "student" edge.
func (uq *UserQuery) QueryStudent() *StudentQuery {
	query := (&StudentClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(student.Table, student.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, user.StudentTable, user.StudentColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTeacher chains the current query on the "teacher" edge.
func (uq *UserQuery) QueryTeacher() *TeacherQuery {
	query := (&TeacherClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(teacher.Table, teacher.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, user.TeacherTable, user.TeacherColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFiles chains the current query on the "files" edge.
func (uq *UserQuery) QueryFiles() *FileQuery {
	query := (&FileClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(file.Table, file.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.FilesTable, user.FilesColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProjects chains the current query on the "projects" edge.
func (uq *UserQuery) QueryProjects() *ProjectQuery {
	query := (&ProjectClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.ProjectsTable, user.ProjectsColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUserRole chains the current query on the "user_role" edge.
func (uq *UserQuery) QueryUserRole() *UserRoleQuery {
	query := (&UserRoleClient{config: uq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, selector),
			sqlgraph.To(userrole.Table, userrole.UserColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, user.UserRoleTable, user.UserRoleColumn),
		)
		fromU = sqlgraph.SetNeighbors(uq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first User entity from the query.
// Returns a *NotFoundError when no User was found.
func (uq *UserQuery) First(ctx context.Context) (*User, error) {
	nodes, err := uq.Limit(1).All(setContextOp(ctx, uq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{user.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uq *UserQuery) FirstX(ctx context.Context) *User {
	node, err := uq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first User ID from the query.
// Returns a *NotFoundError when no User ID was found.
func (uq *UserQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = uq.Limit(1).IDs(setContextOp(ctx, uq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{user.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uq *UserQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := uq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single User entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one User entity is found.
// Returns a *NotFoundError when no User entities are found.
func (uq *UserQuery) Only(ctx context.Context) (*User, error) {
	nodes, err := uq.Limit(2).All(setContextOp(ctx, uq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{user.Label}
	default:
		return nil, &NotSingularError{user.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uq *UserQuery) OnlyX(ctx context.Context) *User {
	node, err := uq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only User ID in the query.
// Returns a *NotSingularError when more than one User ID is found.
// Returns a *NotFoundError when no entities are found.
func (uq *UserQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = uq.Limit(2).IDs(setContextOp(ctx, uq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{user.Label}
	default:
		err = &NotSingularError{user.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uq *UserQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := uq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Users.
func (uq *UserQuery) All(ctx context.Context) ([]*User, error) {
	ctx = setContextOp(ctx, uq.ctx, "All")
	if err := uq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*User, *UserQuery]()
	return withInterceptors[[]*User](ctx, uq, qr, uq.inters)
}

// AllX is like All, but panics if an error occurs.
func (uq *UserQuery) AllX(ctx context.Context) []*User {
	nodes, err := uq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of User IDs.
func (uq *UserQuery) IDs(ctx context.Context) (ids []uint32, err error) {
	if uq.ctx.Unique == nil && uq.path != nil {
		uq.Unique(true)
	}
	ctx = setContextOp(ctx, uq.ctx, "IDs")
	if err = uq.Select(user.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uq *UserQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := uq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uq *UserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, uq.ctx, "Count")
	if err := uq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, uq, querierCount[*UserQuery](), uq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (uq *UserQuery) CountX(ctx context.Context) int {
	count, err := uq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uq *UserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, uq.ctx, "Exist")
	switch _, err := uq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (uq *UserQuery) ExistX(ctx context.Context) bool {
	exist, err := uq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uq *UserQuery) Clone() *UserQuery {
	if uq == nil {
		return nil
	}
	return &UserQuery{
		config:       uq.config,
		ctx:          uq.ctx.Clone(),
		order:        append([]user.OrderOption{}, uq.order...),
		inters:       append([]Interceptor{}, uq.inters...),
		predicates:   append([]predicate.User{}, uq.predicates...),
		withRoles:    uq.withRoles.Clone(),
		withStudent:  uq.withStudent.Clone(),
		withTeacher:  uq.withTeacher.Clone(),
		withFiles:    uq.withFiles.Clone(),
		withProjects: uq.withProjects.Clone(),
		withUserRole: uq.withUserRole.Clone(),
		// clone intermediate query.
		sql:  uq.sql.Clone(),
		path: uq.path,
	}
}

// WithRoles tells the query-builder to eager-load the nodes that are connected to
// the "roles" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithRoles(opts ...func(*RoleQuery)) *UserQuery {
	query := (&RoleClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withRoles = query
	return uq
}

// WithStudent tells the query-builder to eager-load the nodes that are connected to
// the "student" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithStudent(opts ...func(*StudentQuery)) *UserQuery {
	query := (&StudentClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withStudent = query
	return uq
}

// WithTeacher tells the query-builder to eager-load the nodes that are connected to
// the "teacher" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithTeacher(opts ...func(*TeacherQuery)) *UserQuery {
	query := (&TeacherClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withTeacher = query
	return uq
}

// WithFiles tells the query-builder to eager-load the nodes that are connected to
// the "files" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithFiles(opts ...func(*FileQuery)) *UserQuery {
	query := (&FileClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withFiles = query
	return uq
}

// WithProjects tells the query-builder to eager-load the nodes that are connected to
// the "projects" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithProjects(opts ...func(*ProjectQuery)) *UserQuery {
	query := (&ProjectClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withProjects = query
	return uq
}

// WithUserRole tells the query-builder to eager-load the nodes that are connected to
// the "user_role" edge. The optional arguments are used to configure the query builder of the edge.
func (uq *UserQuery) WithUserRole(opts ...func(*UserRoleQuery)) *UserQuery {
	query := (&UserRoleClient{config: uq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uq.withUserRole = query
	return uq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Account string `json:"account,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.User.Query().
//		GroupBy(user.FieldAccount).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uq *UserQuery) GroupBy(field string, fields ...string) *UserGroupBy {
	uq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserGroupBy{build: uq}
	grbuild.flds = &uq.ctx.Fields
	grbuild.label = user.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Account string `json:"account,omitempty"`
//	}
//
//	client.User.Query().
//		Select(user.FieldAccount).
//		Scan(ctx, &v)
func (uq *UserQuery) Select(fields ...string) *UserSelect {
	uq.ctx.Fields = append(uq.ctx.Fields, fields...)
	sbuild := &UserSelect{UserQuery: uq}
	sbuild.label = user.Label
	sbuild.flds, sbuild.scan = &uq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserSelect configured with the given aggregations.
func (uq *UserQuery) Aggregate(fns ...AggregateFunc) *UserSelect {
	return uq.Select().Aggregate(fns...)
}

func (uq *UserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range uq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, uq); err != nil {
				return err
			}
		}
	}
	for _, f := range uq.ctx.Fields {
		if !user.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uq.path != nil {
		prev, err := uq.path(ctx)
		if err != nil {
			return err
		}
		uq.sql = prev
	}
	return nil
}

func (uq *UserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*User, error) {
	var (
		nodes       = []*User{}
		_spec       = uq.querySpec()
		loadedTypes = [6]bool{
			uq.withRoles != nil,
			uq.withStudent != nil,
			uq.withTeacher != nil,
			uq.withFiles != nil,
			uq.withProjects != nil,
			uq.withUserRole != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*User).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &User{config: uq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := uq.withRoles; query != nil {
		if err := uq.loadRoles(ctx, query, nodes,
			func(n *User) { n.Edges.Roles = []*Role{} },
			func(n *User, e *Role) { n.Edges.Roles = append(n.Edges.Roles, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withStudent; query != nil {
		if err := uq.loadStudent(ctx, query, nodes, nil,
			func(n *User, e *Student) { n.Edges.Student = e }); err != nil {
			return nil, err
		}
	}
	if query := uq.withTeacher; query != nil {
		if err := uq.loadTeacher(ctx, query, nodes, nil,
			func(n *User, e *Teacher) { n.Edges.Teacher = e }); err != nil {
			return nil, err
		}
	}
	if query := uq.withFiles; query != nil {
		if err := uq.loadFiles(ctx, query, nodes,
			func(n *User) { n.Edges.Files = []*File{} },
			func(n *User, e *File) { n.Edges.Files = append(n.Edges.Files, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withProjects; query != nil {
		if err := uq.loadProjects(ctx, query, nodes,
			func(n *User) { n.Edges.Projects = []*Project{} },
			func(n *User, e *Project) { n.Edges.Projects = append(n.Edges.Projects, e) }); err != nil {
			return nil, err
		}
	}
	if query := uq.withUserRole; query != nil {
		if err := uq.loadUserRole(ctx, query, nodes,
			func(n *User) { n.Edges.UserRole = []*UserRole{} },
			func(n *User, e *UserRole) { n.Edges.UserRole = append(n.Edges.UserRole, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (uq *UserQuery) loadRoles(ctx context.Context, query *RoleQuery, nodes []*User, init func(*User), assign func(*User, *Role)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[uint32]*User)
	nids := make(map[uint8]map[*User]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(user.RolesTable)
		s.Join(joinT).On(s.C(role.FieldID), joinT.C(user.RolesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(user.RolesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(user.RolesPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := uint32(values[0].(*sql.NullInt64).Int64)
				inValue := uint8(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*User]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Role](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "roles" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (uq *UserQuery) loadStudent(ctx context.Context, query *StudentQuery, nodes []*User, init func(*User), assign func(*User, *Student)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint32]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(student.FieldUID)
	}
	query.Where(predicate.Student(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.StudentColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.UID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "uid" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadTeacher(ctx context.Context, query *TeacherQuery, nodes []*User, init func(*User), assign func(*User, *Teacher)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint32]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(teacher.FieldUID)
	}
	query.Where(predicate.Teacher(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.TeacherColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.UID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "uid" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadFiles(ctx context.Context, query *FileQuery, nodes []*User, init func(*User), assign func(*User, *File)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint32]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(file.FieldUID)
	}
	query.Where(predicate.File(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.FilesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.UID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "uid" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadProjects(ctx context.Context, query *ProjectQuery, nodes []*User, init func(*User), assign func(*User, *Project)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint32]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(project.FieldUID)
	}
	query.Where(predicate.Project(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.ProjectsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.UID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "uid" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uq *UserQuery) loadUserRole(ctx context.Context, query *UserRoleQuery, nodes []*User, init func(*User), assign func(*User, *UserRole)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uint32]*User)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(userrole.FieldUID)
	}
	query.Where(predicate.UserRole(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(user.UserRoleColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.UID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "uid" returned %v for node %v`, fk, n)
		}
		assign(node, n)
	}
	return nil
}

func (uq *UserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uq.querySpec()
	_spec.Node.Columns = uq.ctx.Fields
	if len(uq.ctx.Fields) > 0 {
		_spec.Unique = uq.ctx.Unique != nil && *uq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, uq.driver, _spec)
}

func (uq *UserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint32))
	_spec.From = uq.sql
	if unique := uq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if uq.path != nil {
		_spec.Unique = true
	}
	if fields := uq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for i := range fields {
			if fields[i] != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uq *UserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uq.driver.Dialect())
	t1 := builder.Table(user.Table)
	columns := uq.ctx.Fields
	if len(columns) == 0 {
		columns = user.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uq.sql != nil {
		selector = uq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uq.ctx.Unique != nil && *uq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range uq.predicates {
		p(selector)
	}
	for _, p := range uq.order {
		p(selector)
	}
	if offset := uq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserGroupBy is the group-by builder for User entities.
type UserGroupBy struct {
	selector
	build *UserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ugb *UserGroupBy) Aggregate(fns ...AggregateFunc) *UserGroupBy {
	ugb.fns = append(ugb.fns, fns...)
	return ugb
}

// Scan applies the selector query and scans the result into the given value.
func (ugb *UserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ugb.build.ctx, "GroupBy")
	if err := ugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserQuery, *UserGroupBy](ctx, ugb.build, ugb, ugb.build.inters, v)
}

func (ugb *UserGroupBy) sqlScan(ctx context.Context, root *UserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ugb.fns))
	for _, fn := range ugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ugb.flds)+len(ugb.fns))
		for _, f := range *ugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserSelect is the builder for selecting fields of User entities.
type UserSelect struct {
	*UserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (us *UserSelect) Aggregate(fns ...AggregateFunc) *UserSelect {
	us.fns = append(us.fns, fns...)
	return us
}

// Scan applies the selector query and scans the result into the given value.
func (us *UserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, us.ctx, "Select")
	if err := us.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserQuery, *UserSelect](ctx, us.UserQuery, us, us.inters, v)
}

func (us *UserSelect) sqlScan(ctx context.Context, root *UserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(us.fns))
	for _, fn := range us.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*us.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := us.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
