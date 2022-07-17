// Code generated by entc, DO NOT EDIT.

package ent

import (
	"pandora/ent/casbinrule"
	"pandora/ent/schema"
	"pandora/ent/task"
	"pandora/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	casbinruleFields := schema.CasbinRule{}.Fields()
	_ = casbinruleFields
	// casbinruleDescPtype is the schema descriptor for Ptype field.
	casbinruleDescPtype := casbinruleFields[0].Descriptor()
	// casbinrule.DefaultPtype holds the default value on creation for the Ptype field.
	casbinrule.DefaultPtype = casbinruleDescPtype.Default.(string)
	// casbinruleDescV0 is the schema descriptor for V0 field.
	casbinruleDescV0 := casbinruleFields[1].Descriptor()
	// casbinrule.DefaultV0 holds the default value on creation for the V0 field.
	casbinrule.DefaultV0 = casbinruleDescV0.Default.(string)
	// casbinruleDescV1 is the schema descriptor for V1 field.
	casbinruleDescV1 := casbinruleFields[2].Descriptor()
	// casbinrule.DefaultV1 holds the default value on creation for the V1 field.
	casbinrule.DefaultV1 = casbinruleDescV1.Default.(string)
	// casbinruleDescV2 is the schema descriptor for V2 field.
	casbinruleDescV2 := casbinruleFields[3].Descriptor()
	// casbinrule.DefaultV2 holds the default value on creation for the V2 field.
	casbinrule.DefaultV2 = casbinruleDescV2.Default.(string)
	// casbinruleDescV3 is the schema descriptor for V3 field.
	casbinruleDescV3 := casbinruleFields[4].Descriptor()
	// casbinrule.DefaultV3 holds the default value on creation for the V3 field.
	casbinrule.DefaultV3 = casbinruleDescV3.Default.(string)
	// casbinruleDescV4 is the schema descriptor for V4 field.
	casbinruleDescV4 := casbinruleFields[5].Descriptor()
	// casbinrule.DefaultV4 holds the default value on creation for the V4 field.
	casbinrule.DefaultV4 = casbinruleDescV4.Default.(string)
	// casbinruleDescV5 is the schema descriptor for V5 field.
	casbinruleDescV5 := casbinruleFields[6].Descriptor()
	// casbinrule.DefaultV5 holds the default value on creation for the V5 field.
	casbinrule.DefaultV5 = casbinruleDescV5.Default.(string)
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescCreatedAt is the schema descriptor for createdAt field.
	taskDescCreatedAt := taskFields[7].Descriptor()
	// task.DefaultCreatedAt holds the default value on creation for the createdAt field.
	task.DefaultCreatedAt = taskDescCreatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescRefreshToken is the schema descriptor for refreshToken field.
	userDescRefreshToken := userFields[3].Descriptor()
	// user.DefaultRefreshToken holds the default value on creation for the refreshToken field.
	user.DefaultRefreshToken = userDescRefreshToken.Default.(string)
	// userDescCreatedAt is the schema descriptor for createdAt field.
	userDescCreatedAt := userFields[4].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the createdAt field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updatedAt field.
	userDescUpdatedAt := userFields[5].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}
