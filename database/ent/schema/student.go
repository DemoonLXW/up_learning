package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Student holds the schema definition for the Student entity.
type Student struct {
	ent.Schema
}

func (Student) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "student"},
	}
}

// Fields of the Student.
func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id").
			Unique(),
		field.Uint32("uid").
			Unique(),
		field.Uint32("cid"),
		field.String("student_id").
			Unique().
			NotEmpty(),
		field.String("name").
			NotEmpty(),
		field.Uint8("gender"),
		field.Time("birthday"),
		field.Time("admission_date"),
		field.Uint8("state"),
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

// Edges of the Student.
func (Student) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("class", Class.Type).
			Ref("students").
			Unique().
			Required().
			Field("cid"),
		edge.From("user", User.Type).
			Ref("students").
			Unique().
			Required().
			Field("uid"),
	}
}
