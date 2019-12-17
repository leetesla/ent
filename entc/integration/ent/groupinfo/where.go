// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package groupinfo

import (
	"strconv"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id string) predicate.GroupInfo {
	return predicate.GroupInfo(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		id, _ := strconv.Atoi(id)
		s.Where(sql.EQ(s.C(FieldID), id))
	},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		id, _ := strconv.Atoi(id)
		s.Where(sql.NEQ(s.C(FieldID), id))
	},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i], _ = strconv.Atoi(ids[i])
		}
		s.Where(sql.In(s.C(FieldID), v...))
	},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i], _ = strconv.Atoi(ids[i])
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		id, _ := strconv.Atoi(id)
		s.Where(sql.GT(s.C(FieldID), id))
	},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		id, _ := strconv.Atoi(id)
		s.Where(sql.GTE(s.C(FieldID), id))
	},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		id, _ := strconv.Atoi(id)
		s.Where(sql.LT(s.C(FieldID), id))
	},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		id, _ := strconv.Atoi(id)
		s.Where(sql.LTE(s.C(FieldID), id))
	},
	)
}

// Desc applies equality check predicate on the "desc" field. It's identical to DescEQ.
func Desc(v string) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDesc), v))
	},
	)
}

// MaxUsers applies equality check predicate on the "max_users" field. It's identical to MaxUsersEQ.
func MaxUsers(v int) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaxUsers), v))
	},
	)
}

// DescEQ applies the EQ predicate on the "desc" field.
func DescEQ(v string) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDesc), v))
	},
	)
}

// DescNEQ applies the NEQ predicate on the "desc" field.
func DescNEQ(v string) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDesc), v))
	},
	)
}

// DescIn applies the In predicate on the "desc" field.
func DescIn(vs ...string) predicate.GroupInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDesc), v...))
	},
	)
}

// DescNotIn applies the NotIn predicate on the "desc" field.
func DescNotIn(vs ...string) predicate.GroupInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDesc), v...))
	},
	)
}

// DescGT applies the GT predicate on the "desc" field.
func DescGT(v string) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDesc), v))
	},
	)
}

// DescGTE applies the GTE predicate on the "desc" field.
func DescGTE(v string) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDesc), v))
	},
	)
}

// DescLT applies the LT predicate on the "desc" field.
func DescLT(v string) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDesc), v))
	},
	)
}

// DescLTE applies the LTE predicate on the "desc" field.
func DescLTE(v string) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDesc), v))
	},
	)
}

// DescContains applies the Contains predicate on the "desc" field.
func DescContains(v string) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDesc), v))
	},
	)
}

// DescHasPrefix applies the HasPrefix predicate on the "desc" field.
func DescHasPrefix(v string) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDesc), v))
	},
	)
}

// DescHasSuffix applies the HasSuffix predicate on the "desc" field.
func DescHasSuffix(v string) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDesc), v))
	},
	)
}

// DescEqualFold applies the EqualFold predicate on the "desc" field.
func DescEqualFold(v string) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDesc), v))
	},
	)
}

// DescContainsFold applies the ContainsFold predicate on the "desc" field.
func DescContainsFold(v string) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDesc), v))
	},
	)
}

// MaxUsersEQ applies the EQ predicate on the "max_users" field.
func MaxUsersEQ(v int) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaxUsers), v))
	},
	)
}

// MaxUsersNEQ applies the NEQ predicate on the "max_users" field.
func MaxUsersNEQ(v int) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMaxUsers), v))
	},
	)
}

// MaxUsersIn applies the In predicate on the "max_users" field.
func MaxUsersIn(vs ...int) predicate.GroupInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMaxUsers), v...))
	},
	)
}

// MaxUsersNotIn applies the NotIn predicate on the "max_users" field.
func MaxUsersNotIn(vs ...int) predicate.GroupInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMaxUsers), v...))
	},
	)
}

// MaxUsersGT applies the GT predicate on the "max_users" field.
func MaxUsersGT(v int) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMaxUsers), v))
	},
	)
}

// MaxUsersGTE applies the GTE predicate on the "max_users" field.
func MaxUsersGTE(v int) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMaxUsers), v))
	},
	)
}

// MaxUsersLT applies the LT predicate on the "max_users" field.
func MaxUsersLT(v int) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMaxUsers), v))
	},
	)
}

// MaxUsersLTE applies the LTE predicate on the "max_users" field.
func MaxUsersLTE(v int) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMaxUsers), v))
	},
	)
}

// HasGroups applies the HasEdge predicate on the "groups" edge.
func HasGroups() predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, GroupsTable, GroupsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	},
	)
}

// HasGroupsWith applies the HasEdge predicate on the "groups" edge with a given conditions (other predicates).
func HasGroupsWith(preds ...predicate.Group) predicate.GroupInfo {
	return predicate.GroupInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, GroupsTable, GroupsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.GroupInfo) predicate.GroupInfo {
	return predicate.GroupInfo(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for _, p := range predicates {
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.GroupInfo) predicate.GroupInfo {
	return predicate.GroupInfo(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for i, p := range predicates {
				if i > 0 {
					s1.Or()
				}
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.GroupInfo) predicate.GroupInfo {
	return predicate.GroupInfo(
		func(s *sql.Selector) {
			p(s.Not())
		},
	)
}
