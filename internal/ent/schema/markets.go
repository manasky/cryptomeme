package schema

import (
	"cryptolist/common"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Markets holds the schema definition for the Markets entity.
type Markets struct {
	ent.Schema
}

// Fields of the Markets.
func (Markets) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.JSON("markets", []*common.Market{}),
		field.Time("time"),
	}
}

// Edges of the Markets.
func (Markets) Edges() []ent.Edge {
	return nil
}

func (Markets) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
