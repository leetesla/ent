// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
)

// Group is the model entity for the Group schema.
type Group struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// FromRows scans the sql response data into Group.
func (gr *Group) FromRows(rows *sql.Rows) error {
	var scangr struct {
		ID   int
		Name sql.NullString
	}
	// the order here should be the same as in the `group.Columns`.
	if err := rows.Scan(
		&scangr.ID,
		&scangr.Name,
	); err != nil {
		return err
	}
	gr.ID = scangr.ID
	gr.Name = scangr.Name.String
	return nil
}

// QueryUsers queries the users edge of the Group.
func (gr *Group) QueryUsers() *UserQuery {
	return (&GroupClient{gr.config}).QueryUsers(gr)
}

// Update returns a builder for updating this Group.
// Note that, you need to call Group.Unwrap() before calling this method, if this Group
// was returned from a transaction, and the transaction was committed or rolled back.
func (gr *Group) Update() *GroupUpdateOne {
	return (&GroupClient{gr.config}).UpdateOne(gr)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (gr *Group) Unwrap() *Group {
	tx, ok := gr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Group is not a transactional entity")
	}
	gr.config.driver = tx.drv
	return gr
}

// String implements the fmt.Stringer.
func (gr *Group) String() string {
	var builder strings.Builder
	builder.WriteString("Group(")
	builder.WriteString(fmt.Sprintf("id=%v", gr.ID))
	builder.WriteString(", name=")
	builder.WriteString(gr.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Groups is a parsable slice of Group.
type Groups []*Group

// FromRows scans the sql response data into Groups.
func (gr *Groups) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		scangr := &Group{}
		if err := scangr.FromRows(rows); err != nil {
			return err
		}
		*gr = append(*gr, scangr)
	}
	return nil
}

func (gr Groups) config(cfg config) {
	for _i := range gr {
		gr[_i].config = cfg
	}
}
