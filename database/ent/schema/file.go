package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

func (File) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "file"},
	}
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id").
			Unique(),
		field.Uint32("uid"),
		field.String("name").
			NotEmpty(),
		field.String("path").
			NotEmpty(),
		field.Int64("size").
			Optional(),
		field.Bool("is_disabled").
			Default(false),
		field.Time("created_time").
			Nillable().
			Default(time.Now),
		field.Time("deleted_time").
			Nillable().
			Optional(),
		// Default(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)),
		field.Time("modified_time").
			Nillable().
			Optional(),
		// Default(time.Date(1999, time.November, 11, 0, 0, 0, 0, time.Local)),
	}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", User.Type).
			Ref("files").
			Unique().
			Required().
			Field("uid"),
		edge.From("project", Project.Type).
			Ref("attachments").
			Through("project_file", ProjectFile.Type),
		edge.To("sample", SampleFile.Type).
			Unique(),
	}
}
