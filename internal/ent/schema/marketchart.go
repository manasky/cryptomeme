package schema

import (
	"cryptolist/common"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// MarketChart holds the schema definition for the MarketChart entity.
type MarketChart struct {
	ent.Schema
}

// Fields of the MarketChart.
func (MarketChart) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name"),
		field.String("currency"),
		field.JSON("chart", common.MarketChart{}),
	}
}

// Edges of the MarketChart.
func (MarketChart) Edges() []ent.Edge {
	return nil
}

func (MarketChart) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

//func (MarketChart) Indexes() []ent.Index {
//	return []ent.Index{
//		index.Fields("name", "currency"),
//	}
//}