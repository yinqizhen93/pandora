package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Material holds the schema definition for the Material entity.
type Material struct {
	ent.Schema
}

// Fields of the Material.
func (Material) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("code"),
		field.Text("describe"),
		field.Float("price"),
		field.Time("buyDate").StorageKey("buy_date"),
	}
}

// Edges of the Material.
func (Material) Edges() []ent.Edge {
	return nil
}
