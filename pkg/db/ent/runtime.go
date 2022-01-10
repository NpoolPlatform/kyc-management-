// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/kyc-management/pkg/db/ent/kyc"
	"github.com/NpoolPlatform/kyc-management/pkg/db/ent/schema"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	kycFields := schema.Kyc{}.Fields()
	_ = kycFields
	// kycDescCreateAt is the schema descriptor for create_at field.
	kycDescCreateAt := kycFields[11].Descriptor()
	// kyc.DefaultCreateAt holds the default value on creation for the create_at field.
	kyc.DefaultCreateAt = kycDescCreateAt.Default.(func() uint32)
	// kycDescUpdateAt is the schema descriptor for update_at field.
	kycDescUpdateAt := kycFields[12].Descriptor()
	// kyc.DefaultUpdateAt holds the default value on creation for the update_at field.
	kyc.DefaultUpdateAt = kycDescUpdateAt.Default.(func() uint32)
	// kyc.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	kyc.UpdateDefaultUpdateAt = kycDescUpdateAt.UpdateDefault.(func() uint32)
	// kycDescID is the schema descriptor for id field.
	kycDescID := kycFields[0].Descriptor()
	// kyc.DefaultID holds the default value on creation for the id field.
	kyc.DefaultID = kycDescID.Default.(func() uuid.UUID)
}
