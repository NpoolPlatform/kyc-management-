// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// KycsColumns holds the columns for the "kycs" table.
	KycsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "first_name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString},
		{Name: "region", Type: field.TypeString},
		{Name: "card_type", Type: field.TypeString},
		{Name: "card_id", Type: field.TypeString},
		{Name: "front_card_img", Type: field.TypeString},
		{Name: "back_card_img", Type: field.TypeString},
		{Name: "user_handling_card_img", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
	}
	// KycsTable holds the schema information for the "kycs" table.
	KycsTable = &schema.Table{
		Name:       "kycs",
		Columns:    KycsColumns,
		PrimaryKey: []*schema.Column{KycsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "kyc_user_id_app_id",
				Unique:  false,
				Columns: []*schema.Column{KycsColumns[1], KycsColumns[2]},
			},
			{
				Name:    "kyc_card_id_card_type_app_id",
				Unique:  false,
				Columns: []*schema.Column{KycsColumns[7], KycsColumns[6], KycsColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		KycsTable,
	}
)

func init() {
}
