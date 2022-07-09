package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// Stock holds the schema definition for the Stock entity.
type Stock struct {
	ent.Schema
}

// Fields of the Stock.
func (Stock) Fields() []ent.Field {
	return []ent.Field{
		field.String("market").SchemaType(map[string]string{
			dialect.MySQL: "char(4)",
		}).StructTag(`export:"市场"`),
		field.String("code").StructTag(`export:"股票代码"`),
		field.String("name").StructTag(`export:"股票简称"`),
		field.Time("date").StructTag(`export:"日期"`),
		field.Float32("open").StructTag(`export:"开盘价"`),
		field.Float32("close").StructTag(`export:"收盘价"`),
		field.Float32("high").StructTag(`export:"最高价"`),
		field.Float32("low").StructTag(`export:"最低价"`),
		field.Int32("volume").StructTag(`export:"-"`),
		field.Int32("outstandingShare").StorageKey("outstanding_share").StructTag(`export:"流通量"`),
		field.Float32("turnover"),
	}
}

// Edges of the Stock.
func (Stock) Edges() []ent.Edge {
	return nil
}
