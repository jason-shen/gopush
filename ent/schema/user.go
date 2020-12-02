package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("first_name"),
		field.String("last_name"),
		field.String("email").Unique(),
		field.String("password"),
		field.UUID("apikey", uuid.UUID{}).Default(uuid.New),
		field.String("jwttoken").Optional().Nillable(),
		field.Int8("activate_code"),
		field.Bool("activated").Default(false),
		field.Bool("locked").Default(false),
		field.Time("updated_at").Default(time.Now),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
