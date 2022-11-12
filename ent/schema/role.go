package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.Text("descript").Comment("角色的详细说明"),
		field.Int8("status").Default(1).Comment("1-启用；0-禁用"),
		field.Int8("is_deleted").Default(0).Comment("1-已删除；0-未删除"),
		field.String("access_api"),
		field.String("access_method").
			SchemaType(map[string]string{
				dialect.MySQL: "char(6)", // Override MySQL.
			}),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("roles"),
	}
}
