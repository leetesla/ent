// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/p"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/filetype"
)

// FileTypeCreate is the builder for creating a FileType entity.
type FileTypeCreate struct {
	config
	name  *string
	files map[string]struct{}
}

// SetName sets the name field.
func (ftc *FileTypeCreate) SetName(s string) *FileTypeCreate {
	ftc.name = &s
	return ftc
}

// AddFileIDs adds the files edge to File by ids.
func (ftc *FileTypeCreate) AddFileIDs(ids ...string) *FileTypeCreate {
	if ftc.files == nil {
		ftc.files = make(map[string]struct{})
	}
	for i := range ids {
		ftc.files[ids[i]] = struct{}{}
	}
	return ftc
}

// AddFiles adds the files edges to File.
func (ftc *FileTypeCreate) AddFiles(f ...*File) *FileTypeCreate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftc.AddFileIDs(ids...)
}

// Save creates the FileType in the database.
func (ftc *FileTypeCreate) Save(ctx context.Context) (*FileType, error) {
	if ftc.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	return ftc.gremlinSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (ftc *FileTypeCreate) SaveX(ctx context.Context) *FileType {
	v, err := ftc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ftc *FileTypeCreate) gremlinSave(ctx context.Context) (*FileType, error) {
	res := &gremlin.Response{}
	query, bindings := ftc.gremlin().Query()
	if err := ftc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	ft := &FileType{config: ftc.config}
	if err := ft.FromResponse(res); err != nil {
		return nil, err
	}
	return ft, nil
}

func (ftc *FileTypeCreate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.AddV(filetype.Label)
	if ftc.name != nil {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(filetype.Label, filetype.FieldName, *ftc.name).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(filetype.Label, filetype.FieldName, *ftc.name)),
		})
		v.Property(dsl.Single, filetype.FieldName, *ftc.name)
	}
	for id := range ftc.files {
		v.AddE(filetype.FilesLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(filetype.FilesLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(filetype.Label, filetype.FilesLabel, id)),
		})
	}
	if len(constraints) == 0 {
		return v.ValueMap(true)
	}
	tr := constraints[0].pred.Coalesce(constraints[0].test, v.ValueMap(true))
	for _, cr := range constraints[1:] {
		tr = cr.pred.Coalesce(cr.test, tr)
	}
	return tr
}
