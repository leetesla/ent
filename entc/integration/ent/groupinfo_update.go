// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"strconv"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/ent/group"
	"github.com/facebookincubator/ent/entc/integration/ent/groupinfo"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
	"github.com/facebookincubator/ent/schema/field"
)

// GroupInfoUpdate is the builder for updating GroupInfo entities.
type GroupInfoUpdate struct {
	config
	desc          *string
	max_users     *int
	addmax_users  *int
	groups        map[string]struct{}
	removedGroups map[string]struct{}
	predicates    []predicate.GroupInfo
}

// Where adds a new predicate for the builder.
func (giu *GroupInfoUpdate) Where(ps ...predicate.GroupInfo) *GroupInfoUpdate {
	giu.predicates = append(giu.predicates, ps...)
	return giu
}

// SetDesc sets the desc field.
func (giu *GroupInfoUpdate) SetDesc(s string) *GroupInfoUpdate {
	giu.desc = &s
	return giu
}

// SetMaxUsers sets the max_users field.
func (giu *GroupInfoUpdate) SetMaxUsers(i int) *GroupInfoUpdate {
	giu.max_users = &i
	giu.addmax_users = nil
	return giu
}

// SetNillableMaxUsers sets the max_users field if the given value is not nil.
func (giu *GroupInfoUpdate) SetNillableMaxUsers(i *int) *GroupInfoUpdate {
	if i != nil {
		giu.SetMaxUsers(*i)
	}
	return giu
}

// AddMaxUsers adds i to max_users.
func (giu *GroupInfoUpdate) AddMaxUsers(i int) *GroupInfoUpdate {
	if giu.addmax_users == nil {
		giu.addmax_users = &i
	} else {
		*giu.addmax_users += i
	}
	return giu
}

// AddGroupIDs adds the groups edge to Group by ids.
func (giu *GroupInfoUpdate) AddGroupIDs(ids ...string) *GroupInfoUpdate {
	if giu.groups == nil {
		giu.groups = make(map[string]struct{})
	}
	for i := range ids {
		giu.groups[ids[i]] = struct{}{}
	}
	return giu
}

// AddGroups adds the groups edges to Group.
func (giu *GroupInfoUpdate) AddGroups(g ...*Group) *GroupInfoUpdate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return giu.AddGroupIDs(ids...)
}

// RemoveGroupIDs removes the groups edge to Group by ids.
func (giu *GroupInfoUpdate) RemoveGroupIDs(ids ...string) *GroupInfoUpdate {
	if giu.removedGroups == nil {
		giu.removedGroups = make(map[string]struct{})
	}
	for i := range ids {
		giu.removedGroups[ids[i]] = struct{}{}
	}
	return giu
}

// RemoveGroups removes groups edges to Group.
func (giu *GroupInfoUpdate) RemoveGroups(g ...*Group) *GroupInfoUpdate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return giu.RemoveGroupIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (giu *GroupInfoUpdate) Save(ctx context.Context) (int, error) {
	return giu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (giu *GroupInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := giu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (giu *GroupInfoUpdate) Exec(ctx context.Context) error {
	_, err := giu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (giu *GroupInfoUpdate) ExecX(ctx context.Context) {
	if err := giu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (giu *GroupInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   groupinfo.Table,
			Columns: groupinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: groupinfo.FieldID,
			},
		},
	}
	if ps := giu.predicates; len(ps) > 0 {
		spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value := giu.desc; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: groupinfo.FieldDesc,
		})
	}
	if value := giu.max_users; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: groupinfo.FieldMaxUsers,
		})
	}
	if value := giu.addmax_users; value != nil {
		spec.Fields.Add = append(spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: groupinfo.FieldMaxUsers,
		})
	}
	if nodes := giu.removedGroups; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   groupinfo.GroupsTable,
			Columns: []string{groupinfo.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: group.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			k, err := strconv.Atoi(k)
			if err != nil {
				return 0, err
			}
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges.Clear = append(spec.Edges.Clear, edge)
	}
	if nodes := giu.groups; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   groupinfo.GroupsTable,
			Columns: []string{groupinfo.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: group.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			k, err := strconv.Atoi(k)
			if err != nil {
				return 0, err
			}
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges.Add = append(spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, giu.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// GroupInfoUpdateOne is the builder for updating a single GroupInfo entity.
type GroupInfoUpdateOne struct {
	config
	id            string
	desc          *string
	max_users     *int
	addmax_users  *int
	groups        map[string]struct{}
	removedGroups map[string]struct{}
}

// SetDesc sets the desc field.
func (giuo *GroupInfoUpdateOne) SetDesc(s string) *GroupInfoUpdateOne {
	giuo.desc = &s
	return giuo
}

// SetMaxUsers sets the max_users field.
func (giuo *GroupInfoUpdateOne) SetMaxUsers(i int) *GroupInfoUpdateOne {
	giuo.max_users = &i
	giuo.addmax_users = nil
	return giuo
}

// SetNillableMaxUsers sets the max_users field if the given value is not nil.
func (giuo *GroupInfoUpdateOne) SetNillableMaxUsers(i *int) *GroupInfoUpdateOne {
	if i != nil {
		giuo.SetMaxUsers(*i)
	}
	return giuo
}

// AddMaxUsers adds i to max_users.
func (giuo *GroupInfoUpdateOne) AddMaxUsers(i int) *GroupInfoUpdateOne {
	if giuo.addmax_users == nil {
		giuo.addmax_users = &i
	} else {
		*giuo.addmax_users += i
	}
	return giuo
}

// AddGroupIDs adds the groups edge to Group by ids.
func (giuo *GroupInfoUpdateOne) AddGroupIDs(ids ...string) *GroupInfoUpdateOne {
	if giuo.groups == nil {
		giuo.groups = make(map[string]struct{})
	}
	for i := range ids {
		giuo.groups[ids[i]] = struct{}{}
	}
	return giuo
}

// AddGroups adds the groups edges to Group.
func (giuo *GroupInfoUpdateOne) AddGroups(g ...*Group) *GroupInfoUpdateOne {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return giuo.AddGroupIDs(ids...)
}

// RemoveGroupIDs removes the groups edge to Group by ids.
func (giuo *GroupInfoUpdateOne) RemoveGroupIDs(ids ...string) *GroupInfoUpdateOne {
	if giuo.removedGroups == nil {
		giuo.removedGroups = make(map[string]struct{})
	}
	for i := range ids {
		giuo.removedGroups[ids[i]] = struct{}{}
	}
	return giuo
}

// RemoveGroups removes groups edges to Group.
func (giuo *GroupInfoUpdateOne) RemoveGroups(g ...*Group) *GroupInfoUpdateOne {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return giuo.RemoveGroupIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (giuo *GroupInfoUpdateOne) Save(ctx context.Context) (*GroupInfo, error) {
	return giuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (giuo *GroupInfoUpdateOne) SaveX(ctx context.Context) *GroupInfo {
	gi, err := giuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return gi
}

// Exec executes the query on the entity.
func (giuo *GroupInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := giuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (giuo *GroupInfoUpdateOne) ExecX(ctx context.Context) {
	if err := giuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (giuo *GroupInfoUpdateOne) sqlSave(ctx context.Context) (gi *GroupInfo, err error) {
	spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   groupinfo.Table,
			Columns: groupinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Value:  giuo.id,
				Type:   field.TypeString,
				Column: groupinfo.FieldID,
			},
		},
	}
	if value := giuo.desc; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: groupinfo.FieldDesc,
		})
	}
	if value := giuo.max_users; value != nil {
		spec.Fields.Set = append(spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: groupinfo.FieldMaxUsers,
		})
	}
	if value := giuo.addmax_users; value != nil {
		spec.Fields.Add = append(spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: groupinfo.FieldMaxUsers,
		})
	}
	if nodes := giuo.removedGroups; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   groupinfo.GroupsTable,
			Columns: []string{groupinfo.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: group.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			k, err := strconv.Atoi(k)
			if err != nil {
				return nil, err
			}
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges.Clear = append(spec.Edges.Clear, edge)
	}
	if nodes := giuo.groups; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   groupinfo.GroupsTable,
			Columns: []string{groupinfo.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: group.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			k, err := strconv.Atoi(k)
			if err != nil {
				return nil, err
			}
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges.Add = append(spec.Edges.Add, edge)
	}
	gi = &GroupInfo{config: giuo.config}
	spec.Assign = gi.assignValues
	spec.ScanTypes = gi.scanValues()
	if err = sqlgraph.UpdateNode(ctx, giuo.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return gi, nil
}
