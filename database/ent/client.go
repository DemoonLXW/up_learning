// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/DemoonLXW/up_learning/database/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/DemoonLXW/up_learning/database/ent/permission"
	"github.com/DemoonLXW/up_learning/database/ent/role"
	"github.com/DemoonLXW/up_learning/database/ent/rolepermission"
	"github.com/DemoonLXW/up_learning/database/ent/user"
	"github.com/DemoonLXW/up_learning/database/ent/userrole"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Permission is the client for interacting with the Permission builders.
	Permission *PermissionClient
	// Role is the client for interacting with the Role builders.
	Role *RoleClient
	// RolePermission is the client for interacting with the RolePermission builders.
	RolePermission *RolePermissionClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// UserRole is the client for interacting with the UserRole builders.
	UserRole *UserRoleClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Permission = NewPermissionClient(c.config)
	c.Role = NewRoleClient(c.config)
	c.RolePermission = NewRolePermissionClient(c.config)
	c.User = NewUserClient(c.config)
	c.UserRole = NewUserRoleClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:            ctx,
		config:         cfg,
		Permission:     NewPermissionClient(cfg),
		Role:           NewRoleClient(cfg),
		RolePermission: NewRolePermissionClient(cfg),
		User:           NewUserClient(cfg),
		UserRole:       NewUserRoleClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:            ctx,
		config:         cfg,
		Permission:     NewPermissionClient(cfg),
		Role:           NewRoleClient(cfg),
		RolePermission: NewRolePermissionClient(cfg),
		User:           NewUserClient(cfg),
		UserRole:       NewUserRoleClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Permission.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Permission.Use(hooks...)
	c.Role.Use(hooks...)
	c.RolePermission.Use(hooks...)
	c.User.Use(hooks...)
	c.UserRole.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Permission.Intercept(interceptors...)
	c.Role.Intercept(interceptors...)
	c.RolePermission.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
	c.UserRole.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *PermissionMutation:
		return c.Permission.mutate(ctx, m)
	case *RoleMutation:
		return c.Role.mutate(ctx, m)
	case *RolePermissionMutation:
		return c.RolePermission.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	case *UserRoleMutation:
		return c.UserRole.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// PermissionClient is a client for the Permission schema.
type PermissionClient struct {
	config
}

// NewPermissionClient returns a client for the Permission from the given config.
func NewPermissionClient(c config) *PermissionClient {
	return &PermissionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `permission.Hooks(f(g(h())))`.
func (c *PermissionClient) Use(hooks ...Hook) {
	c.hooks.Permission = append(c.hooks.Permission, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `permission.Intercept(f(g(h())))`.
func (c *PermissionClient) Intercept(interceptors ...Interceptor) {
	c.inters.Permission = append(c.inters.Permission, interceptors...)
}

// Create returns a builder for creating a Permission entity.
func (c *PermissionClient) Create() *PermissionCreate {
	mutation := newPermissionMutation(c.config, OpCreate)
	return &PermissionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Permission entities.
func (c *PermissionClient) CreateBulk(builders ...*PermissionCreate) *PermissionCreateBulk {
	return &PermissionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Permission.
func (c *PermissionClient) Update() *PermissionUpdate {
	mutation := newPermissionMutation(c.config, OpUpdate)
	return &PermissionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PermissionClient) UpdateOne(pe *Permission) *PermissionUpdateOne {
	mutation := newPermissionMutation(c.config, OpUpdateOne, withPermission(pe))
	return &PermissionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PermissionClient) UpdateOneID(id uint16) *PermissionUpdateOne {
	mutation := newPermissionMutation(c.config, OpUpdateOne, withPermissionID(id))
	return &PermissionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Permission.
func (c *PermissionClient) Delete() *PermissionDelete {
	mutation := newPermissionMutation(c.config, OpDelete)
	return &PermissionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PermissionClient) DeleteOne(pe *Permission) *PermissionDeleteOne {
	return c.DeleteOneID(pe.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PermissionClient) DeleteOneID(id uint16) *PermissionDeleteOne {
	builder := c.Delete().Where(permission.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PermissionDeleteOne{builder}
}

// Query returns a query builder for Permission.
func (c *PermissionClient) Query() *PermissionQuery {
	return &PermissionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePermission},
		inters: c.Interceptors(),
	}
}

// Get returns a Permission entity by its id.
func (c *PermissionClient) Get(ctx context.Context, id uint16) (*Permission, error) {
	return c.Query().Where(permission.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PermissionClient) GetX(ctx context.Context, id uint16) *Permission {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRoles queries the roles edge of a Permission.
func (c *PermissionClient) QueryRoles(pe *Permission) *RoleQuery {
	query := (&RoleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(permission.Table, permission.FieldID, id),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, permission.RolesTable, permission.RolesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRolePermission queries the role_permission edge of a Permission.
func (c *PermissionClient) QueryRolePermission(pe *Permission) *RolePermissionQuery {
	query := (&RolePermissionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pe.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(permission.Table, permission.FieldID, id),
			sqlgraph.To(rolepermission.Table, rolepermission.PermissionColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, permission.RolePermissionTable, permission.RolePermissionColumn),
		)
		fromV = sqlgraph.Neighbors(pe.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PermissionClient) Hooks() []Hook {
	return c.hooks.Permission
}

// Interceptors returns the client interceptors.
func (c *PermissionClient) Interceptors() []Interceptor {
	return c.inters.Permission
}

func (c *PermissionClient) mutate(ctx context.Context, m *PermissionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PermissionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PermissionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PermissionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PermissionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Permission mutation op: %q", m.Op())
	}
}

// RoleClient is a client for the Role schema.
type RoleClient struct {
	config
}

// NewRoleClient returns a client for the Role from the given config.
func NewRoleClient(c config) *RoleClient {
	return &RoleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `role.Hooks(f(g(h())))`.
func (c *RoleClient) Use(hooks ...Hook) {
	c.hooks.Role = append(c.hooks.Role, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `role.Intercept(f(g(h())))`.
func (c *RoleClient) Intercept(interceptors ...Interceptor) {
	c.inters.Role = append(c.inters.Role, interceptors...)
}

// Create returns a builder for creating a Role entity.
func (c *RoleClient) Create() *RoleCreate {
	mutation := newRoleMutation(c.config, OpCreate)
	return &RoleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Role entities.
func (c *RoleClient) CreateBulk(builders ...*RoleCreate) *RoleCreateBulk {
	return &RoleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Role.
func (c *RoleClient) Update() *RoleUpdate {
	mutation := newRoleMutation(c.config, OpUpdate)
	return &RoleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RoleClient) UpdateOne(r *Role) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne, withRole(r))
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RoleClient) UpdateOneID(id uint8) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne, withRoleID(id))
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Role.
func (c *RoleClient) Delete() *RoleDelete {
	mutation := newRoleMutation(c.config, OpDelete)
	return &RoleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RoleClient) DeleteOne(r *Role) *RoleDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RoleClient) DeleteOneID(id uint8) *RoleDeleteOne {
	builder := c.Delete().Where(role.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RoleDeleteOne{builder}
}

// Query returns a query builder for Role.
func (c *RoleClient) Query() *RoleQuery {
	return &RoleQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRole},
		inters: c.Interceptors(),
	}
}

// Get returns a Role entity by its id.
func (c *RoleClient) Get(ctx context.Context, id uint8) (*Role, error) {
	return c.Query().Where(role.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RoleClient) GetX(ctx context.Context, id uint8) *Role {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPermissions queries the permissions edge of a Role.
func (c *RoleClient) QueryPermissions(r *Role) *PermissionQuery {
	query := (&PermissionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, id),
			sqlgraph.To(permission.Table, permission.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, role.PermissionsTable, role.PermissionsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUsers queries the users edge of a Role.
func (c *RoleClient) QueryUsers(r *Role) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, role.UsersTable, role.UsersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRolePermission queries the role_permission edge of a Role.
func (c *RoleClient) QueryRolePermission(r *Role) *RolePermissionQuery {
	query := (&RolePermissionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, id),
			sqlgraph.To(rolepermission.Table, rolepermission.RoleColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, role.RolePermissionTable, role.RolePermissionColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUserRole queries the user_role edge of a Role.
func (c *RoleClient) QueryUserRole(r *Role) *UserRoleQuery {
	query := (&UserRoleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, id),
			sqlgraph.To(userrole.Table, userrole.RoleColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, role.UserRoleTable, role.UserRoleColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RoleClient) Hooks() []Hook {
	return c.hooks.Role
}

// Interceptors returns the client interceptors.
func (c *RoleClient) Interceptors() []Interceptor {
	return c.inters.Role
}

func (c *RoleClient) mutate(ctx context.Context, m *RoleMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RoleCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RoleUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RoleDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Role mutation op: %q", m.Op())
	}
}

// RolePermissionClient is a client for the RolePermission schema.
type RolePermissionClient struct {
	config
}

// NewRolePermissionClient returns a client for the RolePermission from the given config.
func NewRolePermissionClient(c config) *RolePermissionClient {
	return &RolePermissionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `rolepermission.Hooks(f(g(h())))`.
func (c *RolePermissionClient) Use(hooks ...Hook) {
	c.hooks.RolePermission = append(c.hooks.RolePermission, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `rolepermission.Intercept(f(g(h())))`.
func (c *RolePermissionClient) Intercept(interceptors ...Interceptor) {
	c.inters.RolePermission = append(c.inters.RolePermission, interceptors...)
}

// Create returns a builder for creating a RolePermission entity.
func (c *RolePermissionClient) Create() *RolePermissionCreate {
	mutation := newRolePermissionMutation(c.config, OpCreate)
	return &RolePermissionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of RolePermission entities.
func (c *RolePermissionClient) CreateBulk(builders ...*RolePermissionCreate) *RolePermissionCreateBulk {
	return &RolePermissionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for RolePermission.
func (c *RolePermissionClient) Update() *RolePermissionUpdate {
	mutation := newRolePermissionMutation(c.config, OpUpdate)
	return &RolePermissionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RolePermissionClient) UpdateOne(rp *RolePermission) *RolePermissionUpdateOne {
	mutation := newRolePermissionMutation(c.config, OpUpdateOne)
	mutation.role = &rp.Rid
	mutation.permission = &rp.Pid
	return &RolePermissionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for RolePermission.
func (c *RolePermissionClient) Delete() *RolePermissionDelete {
	mutation := newRolePermissionMutation(c.config, OpDelete)
	return &RolePermissionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Query returns a query builder for RolePermission.
func (c *RolePermissionClient) Query() *RolePermissionQuery {
	return &RolePermissionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRolePermission},
		inters: c.Interceptors(),
	}
}

// QueryRole queries the role edge of a RolePermission.
func (c *RolePermissionClient) QueryRole(rp *RolePermission) *RoleQuery {
	return c.Query().
		Where(rolepermission.Rid(rp.Rid), rolepermission.Pid(rp.Pid)).
		QueryRole()
}

// QueryPermission queries the permission edge of a RolePermission.
func (c *RolePermissionClient) QueryPermission(rp *RolePermission) *PermissionQuery {
	return c.Query().
		Where(rolepermission.Rid(rp.Rid), rolepermission.Pid(rp.Pid)).
		QueryPermission()
}

// Hooks returns the client hooks.
func (c *RolePermissionClient) Hooks() []Hook {
	return c.hooks.RolePermission
}

// Interceptors returns the client interceptors.
func (c *RolePermissionClient) Interceptors() []Interceptor {
	return c.inters.RolePermission
}

func (c *RolePermissionClient) mutate(ctx context.Context, m *RolePermissionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RolePermissionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RolePermissionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RolePermissionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RolePermissionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown RolePermission mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id uint32) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id uint32) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uint32) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uint32) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRoles queries the roles edge of a User.
func (c *UserClient) QueryRoles(u *User) *RoleQuery {
	query := (&RoleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.RolesTable, user.RolesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUserRole queries the user_role edge of a User.
func (c *UserClient) QueryUserRole(u *User) *UserRoleQuery {
	query := (&UserRoleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(userrole.Table, userrole.UserColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, user.UserRoleTable, user.UserRoleColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// UserRoleClient is a client for the UserRole schema.
type UserRoleClient struct {
	config
}

// NewUserRoleClient returns a client for the UserRole from the given config.
func NewUserRoleClient(c config) *UserRoleClient {
	return &UserRoleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `userrole.Hooks(f(g(h())))`.
func (c *UserRoleClient) Use(hooks ...Hook) {
	c.hooks.UserRole = append(c.hooks.UserRole, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `userrole.Intercept(f(g(h())))`.
func (c *UserRoleClient) Intercept(interceptors ...Interceptor) {
	c.inters.UserRole = append(c.inters.UserRole, interceptors...)
}

// Create returns a builder for creating a UserRole entity.
func (c *UserRoleClient) Create() *UserRoleCreate {
	mutation := newUserRoleMutation(c.config, OpCreate)
	return &UserRoleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UserRole entities.
func (c *UserRoleClient) CreateBulk(builders ...*UserRoleCreate) *UserRoleCreateBulk {
	return &UserRoleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UserRole.
func (c *UserRoleClient) Update() *UserRoleUpdate {
	mutation := newUserRoleMutation(c.config, OpUpdate)
	return &UserRoleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserRoleClient) UpdateOne(ur *UserRole) *UserRoleUpdateOne {
	mutation := newUserRoleMutation(c.config, OpUpdateOne)
	mutation.user = &ur.UID
	mutation.role = &ur.Rid
	return &UserRoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserRole.
func (c *UserRoleClient) Delete() *UserRoleDelete {
	mutation := newUserRoleMutation(c.config, OpDelete)
	return &UserRoleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Query returns a query builder for UserRole.
func (c *UserRoleClient) Query() *UserRoleQuery {
	return &UserRoleQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUserRole},
		inters: c.Interceptors(),
	}
}

// QueryUser queries the user edge of a UserRole.
func (c *UserRoleClient) QueryUser(ur *UserRole) *UserQuery {
	return c.Query().
		Where(userrole.UID(ur.UID), userrole.Rid(ur.Rid)).
		QueryUser()
}

// QueryRole queries the role edge of a UserRole.
func (c *UserRoleClient) QueryRole(ur *UserRole) *RoleQuery {
	return c.Query().
		Where(userrole.UID(ur.UID), userrole.Rid(ur.Rid)).
		QueryRole()
}

// Hooks returns the client hooks.
func (c *UserRoleClient) Hooks() []Hook {
	return c.hooks.UserRole
}

// Interceptors returns the client interceptors.
func (c *UserRoleClient) Interceptors() []Interceptor {
	return c.inters.UserRole
}

func (c *UserRoleClient) mutate(ctx context.Context, m *UserRoleMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserRoleCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserRoleUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserRoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserRoleDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown UserRole mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Permission, Role, RolePermission, User, UserRole []ent.Hook
	}
	inters struct {
		Permission, Role, RolePermission, User, UserRole []ent.Interceptor
	}
)
