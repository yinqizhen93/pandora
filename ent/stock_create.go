// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"pandora/ent/stock"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StockCreate is the builder for creating a Stock entity.
type StockCreate struct {
	config
	mutation *StockMutation
	hooks    []Hook
}

// SetMarket sets the "market" field.
func (sc *StockCreate) SetMarket(s string) *StockCreate {
	sc.mutation.SetMarket(s)
	return sc
}

// SetCode sets the "code" field.
func (sc *StockCreate) SetCode(s string) *StockCreate {
	sc.mutation.SetCode(s)
	return sc
}

// SetName sets the "name" field.
func (sc *StockCreate) SetName(s string) *StockCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetDate sets the "date" field.
func (sc *StockCreate) SetDate(t time.Time) *StockCreate {
	sc.mutation.SetDate(t)
	return sc
}

// SetOpen sets the "open" field.
func (sc *StockCreate) SetOpen(f float32) *StockCreate {
	sc.mutation.SetOpen(f)
	return sc
}

// SetClose sets the "close" field.
func (sc *StockCreate) SetClose(f float32) *StockCreate {
	sc.mutation.SetClose(f)
	return sc
}

// SetHigh sets the "high" field.
func (sc *StockCreate) SetHigh(f float32) *StockCreate {
	sc.mutation.SetHigh(f)
	return sc
}

// SetLow sets the "low" field.
func (sc *StockCreate) SetLow(f float32) *StockCreate {
	sc.mutation.SetLow(f)
	return sc
}

// SetVolume sets the "volume" field.
func (sc *StockCreate) SetVolume(i int32) *StockCreate {
	sc.mutation.SetVolume(i)
	return sc
}

// SetOutstandingShare sets the "outstandingShare" field.
func (sc *StockCreate) SetOutstandingShare(i int32) *StockCreate {
	sc.mutation.SetOutstandingShare(i)
	return sc
}

// SetTurnover sets the "turnover" field.
func (sc *StockCreate) SetTurnover(f float32) *StockCreate {
	sc.mutation.SetTurnover(f)
	return sc
}

// Mutation returns the StockMutation object of the builder.
func (sc *StockCreate) Mutation() *StockMutation {
	return sc.mutation
}

// Save creates the Stock in the database.
func (sc *StockCreate) Save(ctx context.Context) (*Stock, error) {
	var (
		err  error
		node *Stock
	)
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StockCreate) SaveX(ctx context.Context) *Stock {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StockCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StockCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StockCreate) check() error {
	if _, ok := sc.mutation.Market(); !ok {
		return &ValidationError{Name: "market", err: errors.New(`ent: missing required field "Stock.market"`)}
	}
	if v, ok := sc.mutation.Market(); ok {
		if err := stock.MarketValidator(v); err != nil {
			return &ValidationError{Name: "market", err: fmt.Errorf(`ent: validator failed for field "Stock.market": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Stock.code"`)}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Stock.name"`)}
	}
	if _, ok := sc.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New(`ent: missing required field "Stock.date"`)}
	}
	if _, ok := sc.mutation.Open(); !ok {
		return &ValidationError{Name: "open", err: errors.New(`ent: missing required field "Stock.open"`)}
	}
	if _, ok := sc.mutation.Close(); !ok {
		return &ValidationError{Name: "close", err: errors.New(`ent: missing required field "Stock.close"`)}
	}
	if _, ok := sc.mutation.High(); !ok {
		return &ValidationError{Name: "high", err: errors.New(`ent: missing required field "Stock.high"`)}
	}
	if _, ok := sc.mutation.Low(); !ok {
		return &ValidationError{Name: "low", err: errors.New(`ent: missing required field "Stock.low"`)}
	}
	if _, ok := sc.mutation.Volume(); !ok {
		return &ValidationError{Name: "volume", err: errors.New(`ent: missing required field "Stock.volume"`)}
	}
	if _, ok := sc.mutation.OutstandingShare(); !ok {
		return &ValidationError{Name: "outstandingShare", err: errors.New(`ent: missing required field "Stock.outstandingShare"`)}
	}
	if _, ok := sc.mutation.Turnover(); !ok {
		return &ValidationError{Name: "turnover", err: errors.New(`ent: missing required field "Stock.turnover"`)}
	}
	return nil
}

func (sc *StockCreate) sqlSave(ctx context.Context) (*Stock, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sc *StockCreate) createSpec() (*Stock, *sqlgraph.CreateSpec) {
	var (
		_node = &Stock{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: stock.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: stock.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.Market(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: stock.FieldMarket,
		})
		_node.Market = value
	}
	if value, ok := sc.mutation.Code(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: stock.FieldCode,
		})
		_node.Code = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: stock.FieldName,
		})
		_node.Name = value
	}
	if value, ok := sc.mutation.Date(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: stock.FieldDate,
		})
		_node.Date = value
	}
	if value, ok := sc.mutation.Open(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: stock.FieldOpen,
		})
		_node.Open = value
	}
	if value, ok := sc.mutation.Close(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: stock.FieldClose,
		})
		_node.Close = value
	}
	if value, ok := sc.mutation.High(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: stock.FieldHigh,
		})
		_node.High = value
	}
	if value, ok := sc.mutation.Low(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: stock.FieldLow,
		})
		_node.Low = value
	}
	if value, ok := sc.mutation.Volume(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: stock.FieldVolume,
		})
		_node.Volume = value
	}
	if value, ok := sc.mutation.OutstandingShare(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: stock.FieldOutstandingShare,
		})
		_node.OutstandingShare = value
	}
	if value, ok := sc.mutation.Turnover(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: stock.FieldTurnover,
		})
		_node.Turnover = value
	}
	return _node, _spec
}

// StockCreateBulk is the builder for creating many Stock entities in bulk.
type StockCreateBulk struct {
	config
	builders []*StockCreate
}

// Save creates the Stock entities in the database.
func (scb *StockCreateBulk) Save(ctx context.Context) ([]*Stock, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Stock, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StockMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StockCreateBulk) SaveX(ctx context.Context) []*Stock {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StockCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StockCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
