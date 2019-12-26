// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	ent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Dog.
func (Pet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("team", User.Type).
			Ref("team").
			Unique(),
		edge.From("owner", User.Type).
			Ref("pets").
			Unique(),
	}
}

func (Pet) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").
			Edges("owner"),
	}
}
