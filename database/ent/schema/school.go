package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// School holds the schema definition for the School entity.
type School struct {
	ent.Schema
}

func (School) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "school"},
	}
}

// Fields of the School.
func (School) Fields() []ent.Field {
	return []ent.Field{
		field.Uint16("id").
			Unique(),
		field.String("code").
			Unique(),
		field.String("name").
			Unique(),
		field.String("location"),
		field.String("competent_department"),
		field.Uint8("education_level"),
		field.String("remark"),
		field.Bool("is_disabled").
			Default(false),
		field.Time("created_time").
			Nillable().
			Default(time.Now),
		field.Time("deleted_time").
			Nillable().
			Optional(),
		field.Time("modified_time").
			Nillable().
			Optional(),
	}
}

// Edges of the School.
func (School) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("students", Student.Type),
	}
}
