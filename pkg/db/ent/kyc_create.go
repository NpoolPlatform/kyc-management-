// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/kyc-management/pkg/db/ent/kyc"
	"github.com/google/uuid"
)

// KycCreate is the builder for creating a Kyc entity.
type KycCreate struct {
	config
	mutation *KycMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetUserID sets the "user_id" field.
func (kc *KycCreate) SetUserID(u uuid.UUID) *KycCreate {
	kc.mutation.SetUserID(u)
	return kc
}

// SetAppID sets the "app_id" field.
func (kc *KycCreate) SetAppID(u uuid.UUID) *KycCreate {
	kc.mutation.SetAppID(u)
	return kc
}

// SetCardType sets the "card_type" field.
func (kc *KycCreate) SetCardType(s string) *KycCreate {
	kc.mutation.SetCardType(s)
	return kc
}

// SetCardID sets the "card_id" field.
func (kc *KycCreate) SetCardID(s string) *KycCreate {
	kc.mutation.SetCardID(s)
	return kc
}

// SetFrontCardImg sets the "front_card_img" field.
func (kc *KycCreate) SetFrontCardImg(s string) *KycCreate {
	kc.mutation.SetFrontCardImg(s)
	return kc
}

// SetBackCardImg sets the "back_card_img" field.
func (kc *KycCreate) SetBackCardImg(s string) *KycCreate {
	kc.mutation.SetBackCardImg(s)
	return kc
}

// SetUserHandingCardImg sets the "user_handing_card_img" field.
func (kc *KycCreate) SetUserHandingCardImg(s string) *KycCreate {
	kc.mutation.SetUserHandingCardImg(s)
	return kc
}

// SetCreateAt sets the "create_at" field.
func (kc *KycCreate) SetCreateAt(u uint32) *KycCreate {
	kc.mutation.SetCreateAt(u)
	return kc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (kc *KycCreate) SetNillableCreateAt(u *uint32) *KycCreate {
	if u != nil {
		kc.SetCreateAt(*u)
	}
	return kc
}

// SetUpdateAt sets the "update_at" field.
func (kc *KycCreate) SetUpdateAt(u uint32) *KycCreate {
	kc.mutation.SetUpdateAt(u)
	return kc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (kc *KycCreate) SetNillableUpdateAt(u *uint32) *KycCreate {
	if u != nil {
		kc.SetUpdateAt(*u)
	}
	return kc
}

// SetID sets the "id" field.
func (kc *KycCreate) SetID(u uuid.UUID) *KycCreate {
	kc.mutation.SetID(u)
	return kc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (kc *KycCreate) SetNillableID(u *uuid.UUID) *KycCreate {
	if u != nil {
		kc.SetID(*u)
	}
	return kc
}

// Mutation returns the KycMutation object of the builder.
func (kc *KycCreate) Mutation() *KycMutation {
	return kc.mutation
}

// Save creates the Kyc in the database.
func (kc *KycCreate) Save(ctx context.Context) (*Kyc, error) {
	var (
		err  error
		node *Kyc
	)
	kc.defaults()
	if len(kc.hooks) == 0 {
		if err = kc.check(); err != nil {
			return nil, err
		}
		node, err = kc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*KycMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = kc.check(); err != nil {
				return nil, err
			}
			kc.mutation = mutation
			if node, err = kc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(kc.hooks) - 1; i >= 0; i-- {
			if kc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = kc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, kc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (kc *KycCreate) SaveX(ctx context.Context) *Kyc {
	v, err := kc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kc *KycCreate) Exec(ctx context.Context) error {
	_, err := kc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kc *KycCreate) ExecX(ctx context.Context) {
	if err := kc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kc *KycCreate) defaults() {
	if _, ok := kc.mutation.CreateAt(); !ok {
		v := kyc.DefaultCreateAt()
		kc.mutation.SetCreateAt(v)
	}
	if _, ok := kc.mutation.UpdateAt(); !ok {
		v := kyc.DefaultUpdateAt()
		kc.mutation.SetUpdateAt(v)
	}
	if _, ok := kc.mutation.ID(); !ok {
		v := kyc.DefaultID()
		kc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kc *KycCreate) check() error {
	if _, ok := kc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Kyc.user_id"`)}
	}
	if _, ok := kc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "Kyc.app_id"`)}
	}
	if _, ok := kc.mutation.CardType(); !ok {
		return &ValidationError{Name: "card_type", err: errors.New(`ent: missing required field "Kyc.card_type"`)}
	}
	if _, ok := kc.mutation.CardID(); !ok {
		return &ValidationError{Name: "card_id", err: errors.New(`ent: missing required field "Kyc.card_id"`)}
	}
	if _, ok := kc.mutation.FrontCardImg(); !ok {
		return &ValidationError{Name: "front_card_img", err: errors.New(`ent: missing required field "Kyc.front_card_img"`)}
	}
	if _, ok := kc.mutation.BackCardImg(); !ok {
		return &ValidationError{Name: "back_card_img", err: errors.New(`ent: missing required field "Kyc.back_card_img"`)}
	}
	if _, ok := kc.mutation.UserHandingCardImg(); !ok {
		return &ValidationError{Name: "user_handing_card_img", err: errors.New(`ent: missing required field "Kyc.user_handing_card_img"`)}
	}
	if _, ok := kc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "Kyc.create_at"`)}
	}
	if _, ok := kc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "Kyc.update_at"`)}
	}
	return nil
}

func (kc *KycCreate) sqlSave(ctx context.Context) (*Kyc, error) {
	_node, _spec := kc.createSpec()
	if err := sqlgraph.CreateNode(ctx, kc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (kc *KycCreate) createSpec() (*Kyc, *sqlgraph.CreateSpec) {
	var (
		_node = &Kyc{config: kc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: kyc.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: kyc.FieldID,
			},
		}
	)
	_spec.OnConflict = kc.conflict
	if id, ok := kc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := kc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: kyc.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := kc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: kyc.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := kc.mutation.CardType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldCardType,
		})
		_node.CardType = value
	}
	if value, ok := kc.mutation.CardID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldCardID,
		})
		_node.CardID = value
	}
	if value, ok := kc.mutation.FrontCardImg(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldFrontCardImg,
		})
		_node.FrontCardImg = value
	}
	if value, ok := kc.mutation.BackCardImg(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldBackCardImg,
		})
		_node.BackCardImg = value
	}
	if value, ok := kc.mutation.UserHandingCardImg(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: kyc.FieldUserHandingCardImg,
		})
		_node.UserHandingCardImg = value
	}
	if value, ok := kc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kyc.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := kc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: kyc.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Kyc.Create().
//		SetUserID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.KycUpsert) {
//			SetUserID(v+v).
//		}).
//		Exec(ctx)
//
func (kc *KycCreate) OnConflict(opts ...sql.ConflictOption) *KycUpsertOne {
	kc.conflict = opts
	return &KycUpsertOne{
		create: kc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Kyc.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (kc *KycCreate) OnConflictColumns(columns ...string) *KycUpsertOne {
	kc.conflict = append(kc.conflict, sql.ConflictColumns(columns...))
	return &KycUpsertOne{
		create: kc,
	}
}

type (
	// KycUpsertOne is the builder for "upsert"-ing
	//  one Kyc node.
	KycUpsertOne struct {
		create *KycCreate
	}

	// KycUpsert is the "OnConflict" setter.
	KycUpsert struct {
		*sql.UpdateSet
	}
)

// SetUserID sets the "user_id" field.
func (u *KycUpsert) SetUserID(v uuid.UUID) *KycUpsert {
	u.Set(kyc.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *KycUpsert) UpdateUserID() *KycUpsert {
	u.SetExcluded(kyc.FieldUserID)
	return u
}

// SetAppID sets the "app_id" field.
func (u *KycUpsert) SetAppID(v uuid.UUID) *KycUpsert {
	u.Set(kyc.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *KycUpsert) UpdateAppID() *KycUpsert {
	u.SetExcluded(kyc.FieldAppID)
	return u
}

// SetCardType sets the "card_type" field.
func (u *KycUpsert) SetCardType(v string) *KycUpsert {
	u.Set(kyc.FieldCardType, v)
	return u
}

// UpdateCardType sets the "card_type" field to the value that was provided on create.
func (u *KycUpsert) UpdateCardType() *KycUpsert {
	u.SetExcluded(kyc.FieldCardType)
	return u
}

// SetCardID sets the "card_id" field.
func (u *KycUpsert) SetCardID(v string) *KycUpsert {
	u.Set(kyc.FieldCardID, v)
	return u
}

// UpdateCardID sets the "card_id" field to the value that was provided on create.
func (u *KycUpsert) UpdateCardID() *KycUpsert {
	u.SetExcluded(kyc.FieldCardID)
	return u
}

// SetFrontCardImg sets the "front_card_img" field.
func (u *KycUpsert) SetFrontCardImg(v string) *KycUpsert {
	u.Set(kyc.FieldFrontCardImg, v)
	return u
}

// UpdateFrontCardImg sets the "front_card_img" field to the value that was provided on create.
func (u *KycUpsert) UpdateFrontCardImg() *KycUpsert {
	u.SetExcluded(kyc.FieldFrontCardImg)
	return u
}

// SetBackCardImg sets the "back_card_img" field.
func (u *KycUpsert) SetBackCardImg(v string) *KycUpsert {
	u.Set(kyc.FieldBackCardImg, v)
	return u
}

// UpdateBackCardImg sets the "back_card_img" field to the value that was provided on create.
func (u *KycUpsert) UpdateBackCardImg() *KycUpsert {
	u.SetExcluded(kyc.FieldBackCardImg)
	return u
}

// SetUserHandingCardImg sets the "user_handing_card_img" field.
func (u *KycUpsert) SetUserHandingCardImg(v string) *KycUpsert {
	u.Set(kyc.FieldUserHandingCardImg, v)
	return u
}

// UpdateUserHandingCardImg sets the "user_handing_card_img" field to the value that was provided on create.
func (u *KycUpsert) UpdateUserHandingCardImg() *KycUpsert {
	u.SetExcluded(kyc.FieldUserHandingCardImg)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *KycUpsert) SetCreateAt(v uint32) *KycUpsert {
	u.Set(kyc.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *KycUpsert) UpdateCreateAt() *KycUpsert {
	u.SetExcluded(kyc.FieldCreateAt)
	return u
}

// AddCreateAt adds v to the "create_at" field.
func (u *KycUpsert) AddCreateAt(v uint32) *KycUpsert {
	u.Add(kyc.FieldCreateAt, v)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *KycUpsert) SetUpdateAt(v uint32) *KycUpsert {
	u.Set(kyc.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *KycUpsert) UpdateUpdateAt() *KycUpsert {
	u.SetExcluded(kyc.FieldUpdateAt)
	return u
}

// AddUpdateAt adds v to the "update_at" field.
func (u *KycUpsert) AddUpdateAt(v uint32) *KycUpsert {
	u.Add(kyc.FieldUpdateAt, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Kyc.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(kyc.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *KycUpsertOne) UpdateNewValues() *KycUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(kyc.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Kyc.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *KycUpsertOne) Ignore() *KycUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *KycUpsertOne) DoNothing() *KycUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the KycCreate.OnConflict
// documentation for more info.
func (u *KycUpsertOne) Update(set func(*KycUpsert)) *KycUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&KycUpsert{UpdateSet: update})
	}))
	return u
}

// SetUserID sets the "user_id" field.
func (u *KycUpsertOne) SetUserID(v uuid.UUID) *KycUpsertOne {
	return u.Update(func(s *KycUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *KycUpsertOne) UpdateUserID() *KycUpsertOne {
	return u.Update(func(s *KycUpsert) {
		s.UpdateUserID()
	})
}

// SetAppID sets the "app_id" field.
func (u *KycUpsertOne) SetAppID(v uuid.UUID) *KycUpsertOne {
	return u.Update(func(s *KycUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *KycUpsertOne) UpdateAppID() *KycUpsertOne {
	return u.Update(func(s *KycUpsert) {
		s.UpdateAppID()
	})
}

// SetCardType sets the "card_type" field.
func (u *KycUpsertOne) SetCardType(v string) *KycUpsertOne {
	return u.Update(func(s *KycUpsert) {
		s.SetCardType(v)
	})
}

// UpdateCardType sets the "card_type" field to the value that was provided on create.
func (u *KycUpsertOne) UpdateCardType() *KycUpsertOne {
	return u.Update(func(s *KycUpsert) {
		s.UpdateCardType()
	})
}

// SetCardID sets the "card_id" field.
func (u *KycUpsertOne) SetCardID(v string) *KycUpsertOne {
	return u.Update(func(s *KycUpsert) {
		s.SetCardID(v)
	})
}

// UpdateCardID sets the "card_id" field to the value that was provided on create.
func (u *KycUpsertOne) UpdateCardID() *KycUpsertOne {
	return u.Update(func(s *KycUpsert) {
		s.UpdateCardID()
	})
}

// SetFrontCardImg sets the "front_card_img" field.
func (u *KycUpsertOne) SetFrontCardImg(v string) *KycUpsertOne {
	return u.Update(func(s *KycUpsert) {
		s.SetFrontCardImg(v)
	})
}

// UpdateFrontCardImg sets the "front_card_img" field to the value that was provided on create.
func (u *KycUpsertOne) UpdateFrontCardImg() *KycUpsertOne {
	return u.Update(func(s *KycUpsert) {
		s.UpdateFrontCardImg()
	})
}

// SetBackCardImg sets the "back_card_img" field.
func (u *KycUpsertOne) SetBackCardImg(v string) *KycUpsertOne {
	return u.Update(func(s *KycUpsert) {
		s.SetBackCardImg(v)
	})
}

// UpdateBackCardImg sets the "back_card_img" field to the value that was provided on create.
func (u *KycUpsertOne) UpdateBackCardImg() *KycUpsertOne {
	return u.Update(func(s *KycUpsert) {
		s.UpdateBackCardImg()
	})
}

// SetUserHandingCardImg sets the "user_handing_card_img" field.
func (u *KycUpsertOne) SetUserHandingCardImg(v string) *KycUpsertOne {
	return u.Update(func(s *KycUpsert) {
		s.SetUserHandingCardImg(v)
	})
}

// UpdateUserHandingCardImg sets the "user_handing_card_img" field to the value that was provided on create.
func (u *KycUpsertOne) UpdateUserHandingCardImg() *KycUpsertOne {
	return u.Update(func(s *KycUpsert) {
		s.UpdateUserHandingCardImg()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *KycUpsertOne) SetCreateAt(v uint32) *KycUpsertOne {
	return u.Update(func(s *KycUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *KycUpsertOne) AddCreateAt(v uint32) *KycUpsertOne {
	return u.Update(func(s *KycUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *KycUpsertOne) UpdateCreateAt() *KycUpsertOne {
	return u.Update(func(s *KycUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *KycUpsertOne) SetUpdateAt(v uint32) *KycUpsertOne {
	return u.Update(func(s *KycUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *KycUpsertOne) AddUpdateAt(v uint32) *KycUpsertOne {
	return u.Update(func(s *KycUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *KycUpsertOne) UpdateUpdateAt() *KycUpsertOne {
	return u.Update(func(s *KycUpsert) {
		s.UpdateUpdateAt()
	})
}

// Exec executes the query.
func (u *KycUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for KycCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *KycUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *KycUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: KycUpsertOne.ID is not supported by MySQL driver. Use KycUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *KycUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// KycCreateBulk is the builder for creating many Kyc entities in bulk.
type KycCreateBulk struct {
	config
	builders []*KycCreate
	conflict []sql.ConflictOption
}

// Save creates the Kyc entities in the database.
func (kcb *KycCreateBulk) Save(ctx context.Context) ([]*Kyc, error) {
	specs := make([]*sqlgraph.CreateSpec, len(kcb.builders))
	nodes := make([]*Kyc, len(kcb.builders))
	mutators := make([]Mutator, len(kcb.builders))
	for i := range kcb.builders {
		func(i int, root context.Context) {
			builder := kcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*KycMutation)
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
					_, err = mutators[i+1].Mutate(root, kcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = kcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kcb.driver, spec); err != nil {
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
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, kcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (kcb *KycCreateBulk) SaveX(ctx context.Context) []*Kyc {
	v, err := kcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kcb *KycCreateBulk) Exec(ctx context.Context) error {
	_, err := kcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kcb *KycCreateBulk) ExecX(ctx context.Context) {
	if err := kcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Kyc.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.KycUpsert) {
//			SetUserID(v+v).
//		}).
//		Exec(ctx)
//
func (kcb *KycCreateBulk) OnConflict(opts ...sql.ConflictOption) *KycUpsertBulk {
	kcb.conflict = opts
	return &KycUpsertBulk{
		create: kcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Kyc.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (kcb *KycCreateBulk) OnConflictColumns(columns ...string) *KycUpsertBulk {
	kcb.conflict = append(kcb.conflict, sql.ConflictColumns(columns...))
	return &KycUpsertBulk{
		create: kcb,
	}
}

// KycUpsertBulk is the builder for "upsert"-ing
// a bulk of Kyc nodes.
type KycUpsertBulk struct {
	create *KycCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Kyc.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(kyc.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *KycUpsertBulk) UpdateNewValues() *KycUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(kyc.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Kyc.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *KycUpsertBulk) Ignore() *KycUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *KycUpsertBulk) DoNothing() *KycUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the KycCreateBulk.OnConflict
// documentation for more info.
func (u *KycUpsertBulk) Update(set func(*KycUpsert)) *KycUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&KycUpsert{UpdateSet: update})
	}))
	return u
}

// SetUserID sets the "user_id" field.
func (u *KycUpsertBulk) SetUserID(v uuid.UUID) *KycUpsertBulk {
	return u.Update(func(s *KycUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *KycUpsertBulk) UpdateUserID() *KycUpsertBulk {
	return u.Update(func(s *KycUpsert) {
		s.UpdateUserID()
	})
}

// SetAppID sets the "app_id" field.
func (u *KycUpsertBulk) SetAppID(v uuid.UUID) *KycUpsertBulk {
	return u.Update(func(s *KycUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *KycUpsertBulk) UpdateAppID() *KycUpsertBulk {
	return u.Update(func(s *KycUpsert) {
		s.UpdateAppID()
	})
}

// SetCardType sets the "card_type" field.
func (u *KycUpsertBulk) SetCardType(v string) *KycUpsertBulk {
	return u.Update(func(s *KycUpsert) {
		s.SetCardType(v)
	})
}

// UpdateCardType sets the "card_type" field to the value that was provided on create.
func (u *KycUpsertBulk) UpdateCardType() *KycUpsertBulk {
	return u.Update(func(s *KycUpsert) {
		s.UpdateCardType()
	})
}

// SetCardID sets the "card_id" field.
func (u *KycUpsertBulk) SetCardID(v string) *KycUpsertBulk {
	return u.Update(func(s *KycUpsert) {
		s.SetCardID(v)
	})
}

// UpdateCardID sets the "card_id" field to the value that was provided on create.
func (u *KycUpsertBulk) UpdateCardID() *KycUpsertBulk {
	return u.Update(func(s *KycUpsert) {
		s.UpdateCardID()
	})
}

// SetFrontCardImg sets the "front_card_img" field.
func (u *KycUpsertBulk) SetFrontCardImg(v string) *KycUpsertBulk {
	return u.Update(func(s *KycUpsert) {
		s.SetFrontCardImg(v)
	})
}

// UpdateFrontCardImg sets the "front_card_img" field to the value that was provided on create.
func (u *KycUpsertBulk) UpdateFrontCardImg() *KycUpsertBulk {
	return u.Update(func(s *KycUpsert) {
		s.UpdateFrontCardImg()
	})
}

// SetBackCardImg sets the "back_card_img" field.
func (u *KycUpsertBulk) SetBackCardImg(v string) *KycUpsertBulk {
	return u.Update(func(s *KycUpsert) {
		s.SetBackCardImg(v)
	})
}

// UpdateBackCardImg sets the "back_card_img" field to the value that was provided on create.
func (u *KycUpsertBulk) UpdateBackCardImg() *KycUpsertBulk {
	return u.Update(func(s *KycUpsert) {
		s.UpdateBackCardImg()
	})
}

// SetUserHandingCardImg sets the "user_handing_card_img" field.
func (u *KycUpsertBulk) SetUserHandingCardImg(v string) *KycUpsertBulk {
	return u.Update(func(s *KycUpsert) {
		s.SetUserHandingCardImg(v)
	})
}

// UpdateUserHandingCardImg sets the "user_handing_card_img" field to the value that was provided on create.
func (u *KycUpsertBulk) UpdateUserHandingCardImg() *KycUpsertBulk {
	return u.Update(func(s *KycUpsert) {
		s.UpdateUserHandingCardImg()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *KycUpsertBulk) SetCreateAt(v uint32) *KycUpsertBulk {
	return u.Update(func(s *KycUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *KycUpsertBulk) AddCreateAt(v uint32) *KycUpsertBulk {
	return u.Update(func(s *KycUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *KycUpsertBulk) UpdateCreateAt() *KycUpsertBulk {
	return u.Update(func(s *KycUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *KycUpsertBulk) SetUpdateAt(v uint32) *KycUpsertBulk {
	return u.Update(func(s *KycUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *KycUpsertBulk) AddUpdateAt(v uint32) *KycUpsertBulk {
	return u.Update(func(s *KycUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *KycUpsertBulk) UpdateUpdateAt() *KycUpsertBulk {
	return u.Update(func(s *KycUpsert) {
		s.UpdateUpdateAt()
	})
}

// Exec executes the query.
func (u *KycUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the KycCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for KycCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *KycUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
