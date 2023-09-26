package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SampleFile holds the schema definition for the SampleFile entity.
type SampleFile struct {
	ent.Schema
}

func (SampleFile) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sample_file"},
	}
}

// Fields of the SampleFile.
func (SampleFile) Fields() []ent.Field {
	return []ent.Field{
		field.Uint8("id").
			Unique(),
		field.Uint32("fid").
			Unique(),
		field.String("type").
			Unique(),
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

// Edges of the SampleFile.
func (SampleFile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("file", File.Type).
			Ref("sample").
			Unique().
			Required().
			Field("fid"),
	}
}
