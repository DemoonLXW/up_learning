package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

func (Role) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "role"},
	}
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.Uint8("id").
			Unique(),
		field.String("name").
			Unique().
			NotEmpty().
			Nillable(),
		field.String("description").
			Optional().
			Nillable(),
		field.Time("created_time").
			Default(time.Now).
			Nillable(),
		field.Time("deleted_time").
			Default(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)).
			Nillable(),
		field.Time("modified_time").
			Default(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)).
			Nillable(),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("permissions", Permission.Type).
			Through("role_permission", RolePermission.Type),
		edge.To("menu", Menu.Type),
		edge.From("users", User.Type).
			Ref("roles").
			Through("user_role", UserRole.Type),
	}
}
