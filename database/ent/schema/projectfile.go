package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProjectFile holds the schema definition for the ProjectFile entity.
type ProjectFile struct {
	ent.Schema
}

func (ProjectFile) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "project_file"},
		// field.ID("rid", "pid"),
	}
}

// Fields of the ProjectFile.
func (ProjectFile) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("pid"),
		field.Uint32("fid"),
		field.Time("created_time").
			Nillable().
			Default(time.Now),
	}
}

// Edges of the ProjectFile.
func (ProjectFile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("project", Project.Type).
			Unique().
			Required().
			Field("pid"),
		edge.To("files", File.Type).
			Unique().
			Required().
			Field("fid"),
	}
}
