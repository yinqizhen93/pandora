package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Stock holds the schema definition for the Stock entity.
type Stock struct {
	ent.Schema
}

// Fields of the Stock.
func (Stock) Fields() []ent.Field {
	return []ent.Field{
		field.String("market").MaxLen(4),
		field.String("code"),
		field.String("name"),
		field.Time("date"),
		field.Float32("open"),
		field.Float32("close"),
		field.Float32("high"),
		field.Float32("low"),
		field.Int32("volume"),
		field.Int32("outstandingShare").StorageKey("outstanding_share"),
		field.Float32("turnover"),
	}
}

// Edges of the Stock.
func (Stock) Edges() []ent.Edge {
	return nil
}
