package schema

import "github.com/facebook/ent"

// Platform holds the schema definition for the Platform entity.
type Platform struct {
	ent.Schema
}

// Fields of the Platform.
func (Platform) Fields() []ent.Field {
	return nil
}

// Edges of the Platform.
func (Platform) Edges() []ent.Edge {
	return nil
}
