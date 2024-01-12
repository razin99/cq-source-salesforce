package services

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/razin99/cq-source-salesforce/client"
)

type FixedAssetModel struct {
	ID              string `json:"id"`
	ProjectResource string `json:"project_resource__c"`
	Status          string `json:"status__c"`
	Type            string `json:"type__c"`
	SerialNumber    string `json:"serial_number__c"`
}

func FixedAssetTable() *schema.Table {
	return &schema.Table{
		Name:      "salesforce_fixed_asset",
		Transform: transformers.TransformWithStruct(&FixedAssetModel{}, transformers.WithPrimaryKeys("ID")),
		Resolver: func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
			cl := meta.(*client.Client)
			return queryGetAll(
				"select id,project_resource__c, status__c, type__c, serial_number__c FROM fixed_asset__c",
				cl.SalesForce,
				res,
			)
		},
	}
}
