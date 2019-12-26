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

// City is the model entity for the City schema.
type City struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// FromRows scans the sql response data into City.
func (c *City) FromRows(rows *sql.Rows) error {
	var scanc struct {
		ID   int
		Name sql.NullString
	}
	// the order here should be the same as in the `city.Columns`.
	if err := rows.Scan(
		&scanc.ID,
		&scanc.Name,
	); err != nil {
		return err
	}
	c.ID = scanc.ID
	c.Name = scanc.Name.String
	return nil
}

// QueryStreets queries the streets edge of the City.
func (c *City) QueryStreets() *StreetQuery {
	return (&CityClient{c.config}).QueryStreets(c)
}

// Update returns a builder for updating this City.
// Note that, you need to call City.Unwrap() before calling this method, if this City
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *City) Update() *CityUpdateOne {
	return (&CityClient{c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *City) Unwrap() *City {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: City is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *City) String() string {
	var builder strings.Builder
	builder.WriteString("City(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", name=")
	builder.WriteString(c.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Cities is a parsable slice of City.
type Cities []*City

// FromRows scans the sql response data into Cities.
func (c *Cities) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		scanc := &City{}
		if err := scanc.FromRows(rows); err != nil {
			return err
		}
		*c = append(*c, scanc)
	}
	return nil
}

func (c Cities) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
