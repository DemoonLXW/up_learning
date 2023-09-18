package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user"},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id").
			Unique(),
		field.String("account").
			NotEmpty().
			Unique(),

		field.String("password").
			NotEmpty(),
		field.String("username").
			Optional().
			Unique(),
		field.String("email").
			Unique().
			Optional().
			Nillable(),
		field.String("phone").
			Unique().
			Optional().
			Nillable(),
		field.String("introduction").
			Optional(),
		field.Bool("is_disabled").
			Default(false),
		field.Time("created_time").
			Default(time.Now),
		field.Time("deleted_time").
			Default(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)),
		field.Time("modified_time").
			Default(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", Role.Type).
			Through("user_role", UserRole.Type),
		edge.To("students", Student.Type),
	}
}
