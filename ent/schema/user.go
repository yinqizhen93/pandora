package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique(),
		field.String("password").Sensitive(), // Sensitive Fields不会被打印，并且在编码时将被忽略
		field.String("email").Unique(),
		// 字段名建议写成下划线隔开，代码生成时会自动变为PhoneNumber, 如果和数据库字段名不一致，可用StorageKey()指明
		field.String("phone_number").Unique(),
		field.String("refresh_token").Default(""),
		field.Int8("is_active").Default(1),
		field.Time("last_login").Optional().Nillable(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
// M:M relation, one user got many roles, and one role may can belong to mang users
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", Role.Type).Required(),                     // Required()使得创建user时， 必须指定role
		edge.To("department", Department.Type).Unique().Required(), // Unique()表明关系唯一, m or 1 to 1
	}
}
