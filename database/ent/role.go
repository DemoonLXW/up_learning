// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/DemoonLXW/up_learning/database/ent/menu"
	"github.com/DemoonLXW/up_learning/database/ent/role"
)

// Role is the model entity for the Role schema.
type Role struct {
	config `json:"-"`
	// ID of the ent.
	ID uint8 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name *string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
	// CreatedTime holds the value of the "created_time" field.
	CreatedTime *time.Time `json:"created_time,omitempty"`
	// DeletedTime holds the value of the "deleted_time" field.
	DeletedTime *time.Time `json:"deleted_time,omitempty"`
	// ModifiedTime holds the value of the "modified_time" field.
	ModifiedTime *time.Time `json:"modified_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoleQuery when eager-loading is set.
	Edges        RoleEdges `json:"edges"`
	selectValues sql.SelectValues
}

// RoleEdges holds the relations/edges for other nodes in the graph.
type RoleEdges struct {
	// Permissions holds the value of the permissions edge.
	Permissions []*Permission `json:"permissions,omitempty"`
	// Menu holds the value of the menu edge.
	Menu *Menu `json:"menu,omitempty"`
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// RolePermission holds the value of the role_permission edge.
	RolePermission []*RolePermission `json:"role_permission,omitempty"`
	// UserRole holds the value of the user_role edge.
	UserRole []*UserRole `json:"user_role,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// PermissionsOrErr returns the Permissions value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) PermissionsOrErr() ([]*Permission, error) {
	if e.loadedTypes[0] {
		return e.Permissions, nil
	}
	return nil, &NotLoadedError{edge: "permissions"}
}

// MenuOrErr returns the Menu value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoleEdges) MenuOrErr() (*Menu, error) {
	if e.loadedTypes[1] {
		if e.Menu == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: menu.Label}
		}
		return e.Menu, nil
	}
	return nil, &NotLoadedError{edge: "menu"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[2] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// RolePermissionOrErr returns the RolePermission value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) RolePermissionOrErr() ([]*RolePermission, error) {
	if e.loadedTypes[3] {
		return e.RolePermission, nil
	}
	return nil, &NotLoadedError{edge: "role_permission"}
}

// UserRoleOrErr returns the UserRole value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) UserRoleOrErr() ([]*UserRole, error) {
	if e.loadedTypes[4] {
		return e.UserRole, nil
	}
	return nil, &NotLoadedError{edge: "user_role"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Role) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case role.FieldID:
			values[i] = new(sql.NullInt64)
		case role.FieldName, role.FieldDescription:
			values[i] = new(sql.NullString)
		case role.FieldCreatedTime, role.FieldDeletedTime, role.FieldModifiedTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Role fields.
func (r *Role) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case role.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = uint8(value.Int64)
		case role.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = new(string)
				*r.Name = value.String
			}
		case role.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				r.Description = new(string)
				*r.Description = value.String
			}
		case role.FieldCreatedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_time", values[i])
			} else if value.Valid {
				r.CreatedTime = new(time.Time)
				*r.CreatedTime = value.Time
			}
		case role.FieldDeletedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_time", values[i])
			} else if value.Valid {
				r.DeletedTime = new(time.Time)
				*r.DeletedTime = value.Time
			}
		case role.FieldModifiedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field modified_time", values[i])
			} else if value.Valid {
				r.ModifiedTime = new(time.Time)
				*r.ModifiedTime = value.Time
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Role.
// This includes values selected through modifiers, order, etc.
func (r *Role) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryPermissions queries the "permissions" edge of the Role entity.
func (r *Role) QueryPermissions() *PermissionQuery {
	return NewRoleClient(r.config).QueryPermissions(r)
}

// QueryMenu queries the "menu" edge of the Role entity.
func (r *Role) QueryMenu() *MenuQuery {
	return NewRoleClient(r.config).QueryMenu(r)
}

// QueryUsers queries the "users" edge of the Role entity.
func (r *Role) QueryUsers() *UserQuery {
	return NewRoleClient(r.config).QueryUsers(r)
}

// QueryRolePermission queries the "role_permission" edge of the Role entity.
func (r *Role) QueryRolePermission() *RolePermissionQuery {
	return NewRoleClient(r.config).QueryRolePermission(r)
}

// QueryUserRole queries the "user_role" edge of the Role entity.
func (r *Role) QueryUserRole() *UserRoleQuery {
	return NewRoleClient(r.config).QueryUserRole(r)
}

// Update returns a builder for updating this Role.
// Note that you need to call Role.Unwrap() before calling this method if this Role
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Role) Update() *RoleUpdateOne {
	return NewRoleClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Role entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Role) Unwrap() *Role {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Role is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Role) String() string {
	var builder strings.Builder
	builder.WriteString("Role(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	if v := r.Name; v != nil {
		builder.WriteString("name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := r.Description; v != nil {
		builder.WriteString("description=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := r.CreatedTime; v != nil {
		builder.WriteString("created_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := r.DeletedTime; v != nil {
		builder.WriteString("deleted_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := r.ModifiedTime; v != nil {
		builder.WriteString("modified_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Roles is a parsable slice of Role.
type Roles []*Role
