package schema

import "github.com/facebook/ent"

// App holds the schema definition for the App entity.
type App struct {
	ent.Schema
}

// Fields of the App.
func (App) Fields() []ent.Field {
	return nil
}

// Edges of the App.
func (App) Edges() []ent.Edge {
	return nil
}
