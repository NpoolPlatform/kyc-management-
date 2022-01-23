// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kyc-management/pkg/db/ent/kyc"
	"github.com/NpoolPlatform/kyc-management/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// KycUpdate is the builder for updating Kyc entities.
type KycUpdate struct {
	config
	hooks    []Hook
	mutation *KycMutation
}

// Where appends a list predicates to the KycUpdate builder.
func (ku *KycUpdate) Where(ps ...predicate.Kyc) *KycUpdate {
	ku.mutation.Where(ps...)
	return ku
}

// SetUserID sets the "user_id" field.
func (ku *KycUpdate) SetUserID(u uuid.UUID) *KycUpdate {
	ku.mutation.SetUserID(u)
	return ku
}

// SetAppID sets the "app_id" field.
func (ku *KycUpdate) SetAppID(u uuid.UUID) *KycUpdate {
	ku.mutation.SetAppID(u)
	return ku
}

// SetCardType sets the "card_type" field.
func (ku *KycUpdate) SetCardType(s string) *KycUpdate {
	ku.mutation.SetCardType(s)
	return ku
}

// SetCardID sets the "card_id" field.
func (ku *KycUpdate) SetCardID(s string) *KycUpdate {
	ku.mutation.SetCardID(s)
	return ku
}

// SetFrontCardImg sets the "front_card_img" field.
func (ku *KycUpdate) SetFrontCardImg(s string) *KycUpdate {
	ku.mutation.SetFrontCardImg(s)
	return ku
}

// SetBackCardImg sets the "back_card_img" field.
func (ku *KycUpdate) SetBackCardImg(s string) *KycUpdate {
	ku.mutation.SetBackCardImg(s)
	return ku
}

// SetUserHandingCardImg sets the "user_handing_card_img" field.
func (ku *KycUpdate) SetUserHandingCardImg(s string) *KycUpdate {
	ku.mutation.SetUserHandingCardImg(s)
	return ku
}

// SetCreateAt sets the "create_at" field.
func (ku *KycUpdate) SetCreateAt(u uint32) *KycUpdate {
	ku.mutation.ResetCreateAt()
	ku.mutation.SetCreateAt(u)
	return ku
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (ku *KycUpdate) SetNillableCreateAt(u *uint32) *KycUpdate {
	if u != nil {
		ku.SetCreateAt(*u)
	}
	return ku
}

// AddCreateAt adds u to the "create_at" field.
func (ku *KycUpdate) AddCreateAt(u int32) *KycUpdate {
	ku.mutation.AddCreateAt(u)
	return ku
}

// SetUpdateAt sets the "update_at" field.
func (ku *KycUpdate) SetUpdateAt(u uint32) *KycUpdate {
	ku.mutation.ResetUpdateAt()
	ku.mutation.SetUpdateAt(u)
	return ku
}

// AddUpdateAt adds u to the "update_at" field.
func (ku *KycUpdate) AddUpdateAt(u int32) *KycUpdate {
	ku.mutation.AddUpdateAt(u)
	return ku
}

// Mutation returns the KycMutation object of the builder.
func (ku *KycUpdate) Mutation() *KycMutation {
	return ku.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ku *KycUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ku.defaults()
	if len(ku.hooks) == 0 {
		affected, err = ku.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KycMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ku.mutation = mutation
			affected, err = ku.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ku.hooks) - 1; i >= 0; i-- {
			if ku.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ku.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ku.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ku *KycUpdate) SaveX(ctx context.Context) int {
	affected, err := ku.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ku *KycUpdate) Exec(ctx context.Context) error {
	_, err := ku.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ku *KycUpdate) ExecX(ctx context.Context) {
	if err := ku.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ku *KycUpdate) defaults() {
	if _, ok := ku.mutation.UpdateAt(); !ok {
		v := kyc.UpdateDefaultUpdateAt()
		ku.mutation.SetUpdateAt(v)
	}
}

func (ku *KycUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   kyc.Table,
			Columns: kyc.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: kyc.FieldID,
			},
		},
	}
	if ps := ku.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ku.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: kyc.FieldUserID,
		})
	}
	if value, ok := ku.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: kyc.FieldAppID,
		})
	}
	if value, ok := ku.mutation.CardType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldCardType,
		})
	}
	if value, ok := ku.mutation.CardID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldCardID,
		})
	}
	if value, ok := ku.mutation.FrontCardImg(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldFrontCardImg,
		})
	}
	if value, ok := ku.mutation.BackCardImg(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldBackCardImg,
		})
	}
	if value, ok := ku.mutation.UserHandingCardImg(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldUserHandingCardImg,
		})
	}
	if value, ok := ku.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kyc.FieldCreateAt,
		})
	}
	if value, ok := ku.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kyc.FieldCreateAt,
		})
	}
	if value, ok := ku.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kyc.FieldUpdateAt,
		})
	}
	if value, ok := ku.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kyc.FieldUpdateAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ku.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{kyc.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// KycUpdateOne is the builder for updating a single Kyc entity.
type KycUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *KycMutation
}

// SetUserID sets the "user_id" field.
func (kuo *KycUpdateOne) SetUserID(u uuid.UUID) *KycUpdateOne {
	kuo.mutation.SetUserID(u)
	return kuo
}

// SetAppID sets the "app_id" field.
func (kuo *KycUpdateOne) SetAppID(u uuid.UUID) *KycUpdateOne {
	kuo.mutation.SetAppID(u)
	return kuo
}

// SetCardType sets the "card_type" field.
func (kuo *KycUpdateOne) SetCardType(s string) *KycUpdateOne {
	kuo.mutation.SetCardType(s)
	return kuo
}

// SetCardID sets the "card_id" field.
func (kuo *KycUpdateOne) SetCardID(s string) *KycUpdateOne {
	kuo.mutation.SetCardID(s)
	return kuo
}

// SetFrontCardImg sets the "front_card_img" field.
func (kuo *KycUpdateOne) SetFrontCardImg(s string) *KycUpdateOne {
	kuo.mutation.SetFrontCardImg(s)
	return kuo
}

// SetBackCardImg sets the "back_card_img" field.
func (kuo *KycUpdateOne) SetBackCardImg(s string) *KycUpdateOne {
	kuo.mutation.SetBackCardImg(s)
	return kuo
}

// SetUserHandingCardImg sets the "user_handing_card_img" field.
func (kuo *KycUpdateOne) SetUserHandingCardImg(s string) *KycUpdateOne {
	kuo.mutation.SetUserHandingCardImg(s)
	return kuo
}

// SetCreateAt sets the "create_at" field.
func (kuo *KycUpdateOne) SetCreateAt(u uint32) *KycUpdateOne {
	kuo.mutation.ResetCreateAt()
	kuo.mutation.SetCreateAt(u)
	return kuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (kuo *KycUpdateOne) SetNillableCreateAt(u *uint32) *KycUpdateOne {
	if u != nil {
		kuo.SetCreateAt(*u)
	}
	return kuo
}

// AddCreateAt adds u to the "create_at" field.
func (kuo *KycUpdateOne) AddCreateAt(u int32) *KycUpdateOne {
	kuo.mutation.AddCreateAt(u)
	return kuo
}

// SetUpdateAt sets the "update_at" field.
func (kuo *KycUpdateOne) SetUpdateAt(u uint32) *KycUpdateOne {
	kuo.mutation.ResetUpdateAt()
	kuo.mutation.SetUpdateAt(u)
	return kuo
}

// AddUpdateAt adds u to the "update_at" field.
func (kuo *KycUpdateOne) AddUpdateAt(u int32) *KycUpdateOne {
	kuo.mutation.AddUpdateAt(u)
	return kuo
}

// Mutation returns the KycMutation object of the builder.
func (kuo *KycUpdateOne) Mutation() *KycMutation {
	return kuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (kuo *KycUpdateOne) Select(field string, fields ...string) *KycUpdateOne {
	kuo.fields = append([]string{field}, fields...)
	return kuo
}

// Save executes the query and returns the updated Kyc entity.
func (kuo *KycUpdateOne) Save(ctx context.Context) (*Kyc, error) {
	var (
		err  error
		node *Kyc
	)
	kuo.defaults()
	if len(kuo.hooks) == 0 {
		node, err = kuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KycMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			kuo.mutation = mutation
			node, err = kuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(kuo.hooks) - 1; i >= 0; i-- {
			if kuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (kuo *KycUpdateOne) SaveX(ctx context.Context) *Kyc {
	node, err := kuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (kuo *KycUpdateOne) Exec(ctx context.Context) error {
	_, err := kuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kuo *KycUpdateOne) ExecX(ctx context.Context) {
	if err := kuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kuo *KycUpdateOne) defaults() {
	if _, ok := kuo.mutation.UpdateAt(); !ok {
		v := kyc.UpdateDefaultUpdateAt()
		kuo.mutation.SetUpdateAt(v)
	}
}

func (kuo *KycUpdateOne) sqlSave(ctx context.Context) (_node *Kyc, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   kyc.Table,
			Columns: kyc.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: kyc.FieldID,
			},
		},
	}
	id, ok := kuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Kyc.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := kuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, kyc.FieldID)
		for _, f := range fields {
			if !kyc.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != kyc.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := kuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: kyc.FieldUserID,
		})
	}
	if value, ok := kuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: kyc.FieldAppID,
		})
	}
	if value, ok := kuo.mutation.CardType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldCardType,
		})
	}
	if value, ok := kuo.mutation.CardID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldCardID,
		})
	}
	if value, ok := kuo.mutation.FrontCardImg(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldFrontCardImg,
		})
	}
	if value, ok := kuo.mutation.BackCardImg(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldBackCardImg,
		})
	}
	if value, ok := kuo.mutation.UserHandingCardImg(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldUserHandingCardImg,
		})
	}
	if value, ok := kuo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kyc.FieldCreateAt,
		})
	}
	if value, ok := kuo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kyc.FieldCreateAt,
		})
	}
	if value, ok := kuo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kyc.FieldUpdateAt,
		})
	}
	if value, ok := kuo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kyc.FieldUpdateAt,
		})
	}
	_node = &Kyc{config: kuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, kuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{kyc.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
