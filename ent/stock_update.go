// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"pandora/ent/predicate"
	"pandora/ent/stock"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// StockUpdate is the builder for updating Stock entities.
type StockUpdate struct {
	config
	hooks    []Hook
	mutation *StockMutation
}

// Where appends a list predicates to the StockUpdate builder.
func (su *StockUpdate) Where(ps ...predicate.Stock) *StockUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetMarket sets the "market" field.
func (su *StockUpdate) SetMarket(s string) *StockUpdate {
	su.mutation.SetMarket(s)
	return su
}

// SetCode sets the "code" field.
func (su *StockUpdate) SetCode(s string) *StockUpdate {
	su.mutation.SetCode(s)
	return su
}

// SetName sets the "name" field.
func (su *StockUpdate) SetName(s string) *StockUpdate {
	su.mutation.SetName(s)
	return su
}

// SetDate sets the "date" field.
func (su *StockUpdate) SetDate(t time.Time) *StockUpdate {
	su.mutation.SetDate(t)
	return su
}

// SetOpen sets the "open" field.
func (su *StockUpdate) SetOpen(f float32) *StockUpdate {
	su.mutation.ResetOpen()
	su.mutation.SetOpen(f)
	return su
}

// AddOpen adds f to the "open" field.
func (su *StockUpdate) AddOpen(f float32) *StockUpdate {
	su.mutation.AddOpen(f)
	return su
}

// SetClose sets the "close" field.
func (su *StockUpdate) SetClose(f float32) *StockUpdate {
	su.mutation.ResetClose()
	su.mutation.SetClose(f)
	return su
}

// AddClose adds f to the "close" field.
func (su *StockUpdate) AddClose(f float32) *StockUpdate {
	su.mutation.AddClose(f)
	return su
}

// SetHigh sets the "high" field.
func (su *StockUpdate) SetHigh(f float32) *StockUpdate {
	su.mutation.ResetHigh()
	su.mutation.SetHigh(f)
	return su
}

// AddHigh adds f to the "high" field.
func (su *StockUpdate) AddHigh(f float32) *StockUpdate {
	su.mutation.AddHigh(f)
	return su
}

// SetLow sets the "low" field.
func (su *StockUpdate) SetLow(f float32) *StockUpdate {
	su.mutation.ResetLow()
	su.mutation.SetLow(f)
	return su
}

// AddLow adds f to the "low" field.
func (su *StockUpdate) AddLow(f float32) *StockUpdate {
	su.mutation.AddLow(f)
	return su
}

// SetVolume sets the "volume" field.
func (su *StockUpdate) SetVolume(i int32) *StockUpdate {
	su.mutation.ResetVolume()
	su.mutation.SetVolume(i)
	return su
}

// AddVolume adds i to the "volume" field.
func (su *StockUpdate) AddVolume(i int32) *StockUpdate {
	su.mutation.AddVolume(i)
	return su
}

// SetOutstandingShare sets the "outstandingShare" field.
func (su *StockUpdate) SetOutstandingShare(i int32) *StockUpdate {
	su.mutation.ResetOutstandingShare()
	su.mutation.SetOutstandingShare(i)
	return su
}

// AddOutstandingShare adds i to the "outstandingShare" field.
func (su *StockUpdate) AddOutstandingShare(i int32) *StockUpdate {
	su.mutation.AddOutstandingShare(i)
	return su
}

// SetTurnover sets the "turnover" field.
func (su *StockUpdate) SetTurnover(f float32) *StockUpdate {
	su.mutation.ResetTurnover()
	su.mutation.SetTurnover(f)
	return su
}

// AddTurnover adds f to the "turnover" field.
func (su *StockUpdate) AddTurnover(f float32) *StockUpdate {
	su.mutation.AddTurnover(f)
	return su
}

// Mutation returns the StockMutation object of the builder.
func (su *StockUpdate) Mutation() *StockMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StockUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *StockUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StockUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StockUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *StockUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   stock.Table,
			Columns: stock.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: stock.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Market(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: stock.FieldMarket,
		})
	}
	if value, ok := su.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: stock.FieldCode,
		})
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: stock.FieldName,
		})
	}
	if value, ok := su.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: stock.FieldDate,
		})
	}
	if value, ok := su.mutation.Open(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: stock.FieldOpen,
		})
	}
	if value, ok := su.mutation.AddedOpen(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: stock.FieldOpen,
		})
	}
	if value, ok := su.mutation.Close(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: stock.FieldClose,
		})
	}
	if value, ok := su.mutation.AddedClose(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: stock.FieldClose,
		})
	}
	if value, ok := su.mutation.High(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: stock.FieldHigh,
		})
	}
	if value, ok := su.mutation.AddedHigh(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: stock.FieldHigh,
		})
	}
	if value, ok := su.mutation.Low(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: stock.FieldLow,
		})
	}
	if value, ok := su.mutation.AddedLow(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: stock.FieldLow,
		})
	}
	if value, ok := su.mutation.Volume(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: stock.FieldVolume,
		})
	}
	if value, ok := su.mutation.AddedVolume(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: stock.FieldVolume,
		})
	}
	if value, ok := su.mutation.OutstandingShare(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: stock.FieldOutstandingShare,
		})
	}
	if value, ok := su.mutation.AddedOutstandingShare(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: stock.FieldOutstandingShare,
		})
	}
	if value, ok := su.mutation.Turnover(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: stock.FieldTurnover,
		})
	}
	if value, ok := su.mutation.AddedTurnover(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: stock.FieldTurnover,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{stock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// StockUpdateOne is the builder for updating a single Stock entity.
type StockUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StockMutation
}

// SetMarket sets the "market" field.
func (suo *StockUpdateOne) SetMarket(s string) *StockUpdateOne {
	suo.mutation.SetMarket(s)
	return suo
}

// SetCode sets the "code" field.
func (suo *StockUpdateOne) SetCode(s string) *StockUpdateOne {
	suo.mutation.SetCode(s)
	return suo
}

// SetName sets the "name" field.
func (suo *StockUpdateOne) SetName(s string) *StockUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetDate sets the "date" field.
func (suo *StockUpdateOne) SetDate(t time.Time) *StockUpdateOne {
	suo.mutation.SetDate(t)
	return suo
}

// SetOpen sets the "open" field.
func (suo *StockUpdateOne) SetOpen(f float32) *StockUpdateOne {
	suo.mutation.ResetOpen()
	suo.mutation.SetOpen(f)
	return suo
}

// AddOpen adds f to the "open" field.
func (suo *StockUpdateOne) AddOpen(f float32) *StockUpdateOne {
	suo.mutation.AddOpen(f)
	return suo
}

// SetClose sets the "close" field.
func (suo *StockUpdateOne) SetClose(f float32) *StockUpdateOne {
	suo.mutation.ResetClose()
	suo.mutation.SetClose(f)
	return suo
}

// AddClose adds f to the "close" field.
func (suo *StockUpdateOne) AddClose(f float32) *StockUpdateOne {
	suo.mutation.AddClose(f)
	return suo
}

// SetHigh sets the "high" field.
func (suo *StockUpdateOne) SetHigh(f float32) *StockUpdateOne {
	suo.mutation.ResetHigh()
	suo.mutation.SetHigh(f)
	return suo
}

// AddHigh adds f to the "high" field.
func (suo *StockUpdateOne) AddHigh(f float32) *StockUpdateOne {
	suo.mutation.AddHigh(f)
	return suo
}

// SetLow sets the "low" field.
func (suo *StockUpdateOne) SetLow(f float32) *StockUpdateOne {
	suo.mutation.ResetLow()
	suo.mutation.SetLow(f)
	return suo
}

// AddLow adds f to the "low" field.
func (suo *StockUpdateOne) AddLow(f float32) *StockUpdateOne {
	suo.mutation.AddLow(f)
	return suo
}

// SetVolume sets the "volume" field.
func (suo *StockUpdateOne) SetVolume(i int32) *StockUpdateOne {
	suo.mutation.ResetVolume()
	suo.mutation.SetVolume(i)
	return suo
}

// AddVolume adds i to the "volume" field.
func (suo *StockUpdateOne) AddVolume(i int32) *StockUpdateOne {
	suo.mutation.AddVolume(i)
	return suo
}

// SetOutstandingShare sets the "outstandingShare" field.
func (suo *StockUpdateOne) SetOutstandingShare(i int32) *StockUpdateOne {
	suo.mutation.ResetOutstandingShare()
	suo.mutation.SetOutstandingShare(i)
	return suo
}

// AddOutstandingShare adds i to the "outstandingShare" field.
func (suo *StockUpdateOne) AddOutstandingShare(i int32) *StockUpdateOne {
	suo.mutation.AddOutstandingShare(i)
	return suo
}

// SetTurnover sets the "turnover" field.
func (suo *StockUpdateOne) SetTurnover(f float32) *StockUpdateOne {
	suo.mutation.ResetTurnover()
	suo.mutation.SetTurnover(f)
	return suo
}

// AddTurnover adds f to the "turnover" field.
func (suo *StockUpdateOne) AddTurnover(f float32) *StockUpdateOne {
	suo.mutation.AddTurnover(f)
	return suo
}

// Mutation returns the StockMutation object of the builder.
func (suo *StockUpdateOne) Mutation() *StockMutation {
	return suo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StockUpdateOne) Select(field string, fields ...string) *StockUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Stock entity.
func (suo *StockUpdateOne) Save(ctx context.Context) (*Stock, error) {
	var (
		err  error
		node *Stock
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StockUpdateOne) SaveX(ctx context.Context) *Stock {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StockUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StockUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *StockUpdateOne) sqlSave(ctx context.Context) (_node *Stock, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   stock.Table,
			Columns: stock.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: stock.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Stock.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, stock.FieldID)
		for _, f := range fields {
			if !stock.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != stock.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Market(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: stock.FieldMarket,
		})
	}
	if value, ok := suo.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: stock.FieldCode,
		})
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: stock.FieldName,
		})
	}
	if value, ok := suo.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: stock.FieldDate,
		})
	}
	if value, ok := suo.mutation.Open(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: stock.FieldOpen,
		})
	}
	if value, ok := suo.mutation.AddedOpen(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: stock.FieldOpen,
		})
	}
	if value, ok := suo.mutation.Close(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: stock.FieldClose,
		})
	}
	if value, ok := suo.mutation.AddedClose(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: stock.FieldClose,
		})
	}
	if value, ok := suo.mutation.High(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: stock.FieldHigh,
		})
	}
	if value, ok := suo.mutation.AddedHigh(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: stock.FieldHigh,
		})
	}
	if value, ok := suo.mutation.Low(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: stock.FieldLow,
		})
	}
	if value, ok := suo.mutation.AddedLow(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: stock.FieldLow,
		})
	}
	if value, ok := suo.mutation.Volume(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: stock.FieldVolume,
		})
	}
	if value, ok := suo.mutation.AddedVolume(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: stock.FieldVolume,
		})
	}
	if value, ok := suo.mutation.OutstandingShare(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: stock.FieldOutstandingShare,
		})
	}
	if value, ok := suo.mutation.AddedOutstandingShare(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: stock.FieldOutstandingShare,
		})
	}
	if value, ok := suo.mutation.Turnover(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: stock.FieldTurnover,
		})
	}
	if value, ok := suo.mutation.AddedTurnover(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: stock.FieldTurnover,
		})
	}
	_node = &Stock{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{stock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
