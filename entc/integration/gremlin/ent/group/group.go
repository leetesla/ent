// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package group

import (
	"github.com/facebookincubator/ent/entc/integration/ent/schema"
)

const (
	// Label holds the string label denoting the group type in the database.
	Label = "group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldActive holds the string denoting the active vertex property in the database.
	FieldActive = "active"
	// FieldExpire holds the string denoting the expire vertex property in the database.
	FieldExpire = "expire"
	// FieldType holds the string denoting the type vertex property in the database.
	FieldType = "type"
	// FieldMaxUsers holds the string denoting the max_users vertex property in the database.
	FieldMaxUsers = "max_users"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"

	// FilesLabel holds the string label denoting the files edge type in the database.
	FilesLabel = "group_files"
	// BlockedLabel holds the string label denoting the blocked edge type in the database.
	BlockedLabel = "group_blocked"
	// UsersInverseLabel holds the string label denoting the users inverse edge type in the database.
	UsersInverseLabel = "user_groups"
	// InfoLabel holds the string label denoting the info edge type in the database.
	InfoLabel = "group_info"
)

var (
	fields = schema.Group{}.Fields()

	// descActive is the schema descriptor for active field.
	descActive = fields[0].Descriptor()
	// DefaultActive holds the default value on creation for the active field.
	DefaultActive = descActive.Default.(bool)

	// descType is the schema descriptor for type field.
	descType = fields[2].Descriptor()
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator = func() func(string) error {
		validators := descType.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(_type string) error {
			for _, fn := range fns {
				if err := fn(_type); err != nil {
					return err
				}
			}
			return nil
		}
	}()

	// descMaxUsers is the schema descriptor for max_users field.
	descMaxUsers = fields[3].Descriptor()
	// DefaultMaxUsers holds the default value on creation for the max_users field.
	DefaultMaxUsers = descMaxUsers.Default.(int)
	// MaxUsersValidator is a validator for the "max_users" field. It is called by the builders before save.
	MaxUsersValidator = descMaxUsers.Validators[0].(func(int) error)

	// descName is the schema descriptor for name field.
	descName = fields[4].Descriptor()
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator = func() func(string) error {
		validators := descName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
)
