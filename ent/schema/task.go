package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Enum("type").Values("once", "periodical"),
		field.Text("describe"),
		field.Time("startDate").StorageKey("start_date"),
		field.Time("endDate").StorageKey("end_date").Optional(),
		field.Int("costTime").StorageKey("cost_time").Optional(),
		field.Int8("status"),
		field.Time("createdAt").StorageKey("created_at").Default(time.Now).Immutable(),
		field.Int("createdBy").StorageKey("created_by"),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return nil
}
