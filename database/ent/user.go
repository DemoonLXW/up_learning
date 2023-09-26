// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/DemoonLXW/up_learning/database/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID uint32 `json:"id,omitempty"`
	// Account holds the value of the "account" field.
	Account string `json:"account,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Email holds the value of the "email" field.
	Email *string `json:"email,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone *string `json:"phone,omitempty"`
	// Introduction holds the value of the "introduction" field.
	Introduction string `json:"introduction,omitempty"`
	// IsDisabled holds the value of the "is_disabled" field.
	IsDisabled bool `json:"is_disabled,omitempty"`
	// CreatedTime holds the value of the "created_time" field.
	CreatedTime *time.Time `json:"created_time,omitempty"`
	// DeletedTime holds the value of the "deleted_time" field.
	DeletedTime *time.Time `json:"deleted_time,omitempty"`
	// ModifiedTime holds the value of the "modified_time" field.
	ModifiedTime *time.Time `json:"modified_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Roles holds the value of the roles edge.
	Roles []*Role `json:"roles,omitempty"`
	// Students holds the value of the students edge.
	Students []*Student `json:"students,omitempty"`
	// Files holds the value of the files edge.
	Files []*File `json:"files,omitempty"`
	// UserRole holds the value of the user_role edge.
	UserRole []*UserRole `json:"user_role,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// RolesOrErr returns the Roles value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RolesOrErr() ([]*Role, error) {
	if e.loadedTypes[0] {
		return e.Roles, nil
	}
	return nil, &NotLoadedError{edge: "roles"}
}

// StudentsOrErr returns the Students value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) StudentsOrErr() ([]*Student, error) {
	if e.loadedTypes[1] {
		return e.Students, nil
	}
	return nil, &NotLoadedError{edge: "students"}
}

// FilesOrErr returns the Files value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FilesOrErr() ([]*File, error) {
	if e.loadedTypes[2] {
		return e.Files, nil
	}
	return nil, &NotLoadedError{edge: "files"}
}

// UserRoleOrErr returns the UserRole value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserRoleOrErr() ([]*UserRole, error) {
	if e.loadedTypes[3] {
		return e.UserRole, nil
	}
	return nil, &NotLoadedError{edge: "user_role"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldIsDisabled:
			values[i] = new(sql.NullBool)
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldAccount, user.FieldPassword, user.FieldUsername, user.FieldEmail, user.FieldPhone, user.FieldIntroduction:
			values[i] = new(sql.NullString)
		case user.FieldCreatedTime, user.FieldDeletedTime, user.FieldModifiedTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = uint32(value.Int64)
		case user.FieldAccount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field account", values[i])
			} else if value.Valid {
				u.Account = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = new(string)
				*u.Email = value.String
			}
		case user.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				u.Phone = new(string)
				*u.Phone = value.String
			}
		case user.FieldIntroduction:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field introduction", values[i])
			} else if value.Valid {
				u.Introduction = value.String
			}
		case user.FieldIsDisabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_disabled", values[i])
			} else if value.Valid {
				u.IsDisabled = value.Bool
			}
		case user.FieldCreatedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_time", values[i])
			} else if value.Valid {
				u.CreatedTime = new(time.Time)
				*u.CreatedTime = value.Time
			}
		case user.FieldDeletedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_time", values[i])
			} else if value.Valid {
				u.DeletedTime = new(time.Time)
				*u.DeletedTime = value.Time
			}
		case user.FieldModifiedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field modified_time", values[i])
			} else if value.Valid {
				u.ModifiedTime = new(time.Time)
				*u.ModifiedTime = value.Time
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryRoles queries the "roles" edge of the User entity.
func (u *User) QueryRoles() *RoleQuery {
	return NewUserClient(u.config).QueryRoles(u)
}

// QueryStudents queries the "students" edge of the User entity.
func (u *User) QueryStudents() *StudentQuery {
	return NewUserClient(u.config).QueryStudents(u)
}

// QueryFiles queries the "files" edge of the User entity.
func (u *User) QueryFiles() *FileQuery {
	return NewUserClient(u.config).QueryFiles(u)
}

// QueryUserRole queries the "user_role" edge of the User entity.
func (u *User) QueryUserRole() *UserRoleQuery {
	return NewUserClient(u.config).QueryUserRole(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("account=")
	builder.WriteString(u.Account)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(u.Password)
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	if v := u.Email; v != nil {
		builder.WriteString("email=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := u.Phone; v != nil {
		builder.WriteString("phone=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("introduction=")
	builder.WriteString(u.Introduction)
	builder.WriteString(", ")
	builder.WriteString("is_disabled=")
	builder.WriteString(fmt.Sprintf("%v", u.IsDisabled))
	builder.WriteString(", ")
	if v := u.CreatedTime; v != nil {
		builder.WriteString("created_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := u.DeletedTime; v != nil {
		builder.WriteString("deleted_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := u.ModifiedTime; v != nil {
		builder.WriteString("modified_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
