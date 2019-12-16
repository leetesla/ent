// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package entv1

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv1/user"
	"github.com/facebookincubator/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	age      *int32
	name     *string
	nickname *string
	address  *string
	renamed  *string
	blob     *[]byte
	state    *user.State
}

// SetAge sets the age field.
func (uc *UserCreate) SetAge(i int32) *UserCreate {
	uc.age = &i
	return uc
}

// SetName sets the name field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.name = &s
	return uc
}

// SetNickname sets the nickname field.
func (uc *UserCreate) SetNickname(s string) *UserCreate {
	uc.nickname = &s
	return uc
}

// SetAddress sets the address field.
func (uc *UserCreate) SetAddress(s string) *UserCreate {
	uc.address = &s
	return uc
}

// SetNillableAddress sets the address field if the given value is not nil.
func (uc *UserCreate) SetNillableAddress(s *string) *UserCreate {
	if s != nil {
		uc.SetAddress(*s)
	}
	return uc
}

// SetRenamed sets the renamed field.
func (uc *UserCreate) SetRenamed(s string) *UserCreate {
	uc.renamed = &s
	return uc
}

// SetNillableRenamed sets the renamed field if the given value is not nil.
func (uc *UserCreate) SetNillableRenamed(s *string) *UserCreate {
	if s != nil {
		uc.SetRenamed(*s)
	}
	return uc
}

// SetBlob sets the blob field.
func (uc *UserCreate) SetBlob(b []byte) *UserCreate {
	uc.blob = &b
	return uc
}

// SetState sets the state field.
func (uc *UserCreate) SetState(u user.State) *UserCreate {
	uc.state = &u
	return uc
}

// SetNillableState sets the state field if the given value is not nil.
func (uc *UserCreate) SetNillableState(u *user.State) *UserCreate {
	if u != nil {
		uc.SetState(*u)
	}
	return uc
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if uc.age == nil {
		return nil, errors.New("entv1: missing required field \"age\"")
	}
	if uc.name == nil {
		return nil, errors.New("entv1: missing required field \"name\"")
	}
	if err := user.NameValidator(*uc.name); err != nil {
		return nil, fmt.Errorf("entv1: validator failed for field \"name\": %v", err)
	}
	if uc.nickname == nil {
		return nil, errors.New("entv1: missing required field \"nickname\"")
	}
	if uc.state != nil {
		if err := user.StateValidator(*uc.state); err != nil {
			return nil, fmt.Errorf("entv1: validator failed for field \"state\": %v", err)
		}
	}
	return uc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	var (
		u    = &User{config: uc.config}
		spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if value := uc.age; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  *value,
			Column: user.FieldAge,
		})
		u.Age = *value
	}
	if value := uc.name; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldName,
		})
		u.Name = *value
	}
	if value := uc.nickname; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldNickname,
		})
		u.Nickname = *value
	}
	if value := uc.address; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldAddress,
		})
		u.Address = *value
	}
	if value := uc.renamed; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldRenamed,
		})
		u.Renamed = *value
	}
	if value := uc.blob; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  *value,
			Column: user.FieldBlob,
		})
		u.Blob = *value
	}
	if value := uc.state; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  *value,
			Column: user.FieldState,
		})
		u.State = *value
	}
	if err := sqlgraph.CreateNode(ctx, uc.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := spec.ID.Value.(int64)
	u.ID = int(id)
	return u, nil
}
