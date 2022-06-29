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
		field.String("refreshToken").StorageKey("refresh_token"),
		field.Time("createdAt").StorageKey("created_at").Default(time.Now).Immutable(),
		field.Time("updatedAt").StorageKey("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", Role.Type),
	}
}
