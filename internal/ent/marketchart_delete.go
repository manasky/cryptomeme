// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"github.com/manasky/cryptomeme/internal/ent/marketchart"
	"github.com/manasky/cryptomeme/internal/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MarketChartDelete is the builder for deleting a MarketChart entity.
type MarketChartDelete struct {
	config
	hooks    []Hook
	mutation *MarketChartMutation
}

// Where appends a list predicates to the MarketChartDelete builder.
func (mcd *MarketChartDelete) Where(ps ...predicate.MarketChart) *MarketChartDelete {
	mcd.mutation.Where(ps...)
	return mcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mcd *MarketChartDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mcd.hooks) == 0 {
		affected, err = mcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MarketChartMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mcd.mutation = mutation
			affected, err = mcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mcd.hooks) - 1; i >= 0; i-- {
			if mcd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcd *MarketChartDelete) ExecX(ctx context.Context) int {
	n, err := mcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mcd *MarketChartDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: marketchart.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: marketchart.FieldID,
			},
		},
	}
	if ps := mcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, mcd.driver, _spec)
}

// MarketChartDeleteOne is the builder for deleting a single MarketChart entity.
type MarketChartDeleteOne struct {
	mcd *MarketChartDelete
}

// Exec executes the deletion query.
func (mcdo *MarketChartDeleteOne) Exec(ctx context.Context) error {
	n, err := mcdo.mcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{marketchart.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mcdo *MarketChartDeleteOne) ExecX(ctx context.Context) {
	mcdo.mcd.ExecX(ctx)
}
