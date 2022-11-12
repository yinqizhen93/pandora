package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Department holds the schema definition for the Department entity.
type Department struct {
	ent.Schema
}

// Fields of the Department.
func (Department) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").Unique(),
		field.String("name"),
		field.Int("parent_id"),
		field.Int8("is_deleted").Default(0),
		field.Int("created_by"),
		field.Int("updated_by"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Department.
func (Department) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("department"),
		edge.To("children1", Department.Type).
			From("parent1").Unique(),
	}
}
