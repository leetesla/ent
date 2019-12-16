// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/json/ent/user"
	"github.com/facebookincubator/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	url     **url.URL
	raw     *json.RawMessage
	dirs    *[]http.Dir
	ints    *[]int
	floats  *[]float64
	strings *[]string
}

// SetURL sets the url field.
func (uc *UserCreate) SetURL(u *url.URL) *UserCreate {
	uc.url = &u
	return uc
}

// SetRaw sets the raw field.
func (uc *UserCreate) SetRaw(jm json.RawMessage) *UserCreate {
	uc.raw = &jm
	return uc
}

// SetDirs sets the dirs field.
func (uc *UserCreate) SetDirs(h []http.Dir) *UserCreate {
	uc.dirs = &h
	return uc
}

// SetInts sets the ints field.
func (uc *UserCreate) SetInts(i []int) *UserCreate {
	uc.ints = &i
	return uc
}

// SetFloats sets the floats field.
func (uc *UserCreate) SetFloats(f []float64) *UserCreate {
	uc.floats = &f
	return uc
}

// SetStrings sets the strings field.
func (uc *UserCreate) SetStrings(s []string) *UserCreate {
	uc.strings = &s
	return uc
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
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
	if value := uc.url; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  *value,
			Column: user.FieldURL,
		})
		u.URL = *value
	}
	if value := uc.raw; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  *value,
			Column: user.FieldRaw,
		})
		u.Raw = *value
	}
	if value := uc.dirs; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  *value,
			Column: user.FieldDirs,
		})
		u.Dirs = *value
	}
	if value := uc.ints; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  *value,
			Column: user.FieldInts,
		})
		u.Ints = *value
	}
	if value := uc.floats; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  *value,
			Column: user.FieldFloats,
		})
		u.Floats = *value
	}
	if value := uc.strings; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  *value,
			Column: user.FieldStrings,
		})
		u.Strings = *value
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
