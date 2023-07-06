package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/DemoonLXW/up_learning/entity"
)

// Menu holds the schema definition for the Menu entity.
type Menu struct {
	ent.Schema
}

func (Menu) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "menu"},
	}
}

// Fields of the Menu.
func (Menu) Fields() []ent.Field {
	return []ent.Field{
		field.Uint8("id").
			Unique(),
		field.String("name").
			Unique().
			NotEmpty().
			Nillable(),
		field.JSON("json_menu", []*entity.Menu{}).
			Optional(),
		field.Uint8("rid"),
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

// Edges of the Menu.
func (Menu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("role", Role.Type).
			Ref("menu").
			Unique().
			Required().
			Field("rid"),
	}
}
