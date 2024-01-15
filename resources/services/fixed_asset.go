package services

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/razin99/cq-source-salesforce/client"
)

type FixedAssetModel struct {
	Id                  string `json:"id"`
	Project_Resource__c string `json:"project_resource__c"`
	Status__c           string `json:"status__c"`
	Type__c             string `json:"type__c"`
	Serial_Number__c    string `json:"serial_number__c"`
}

func FixedAssetTable() *schema.Table {
	return &schema.Table{
		Name: "salesforce_fixed_asset",
		Transform: transformers.TransformWithStruct(
			&FixedAssetModel{},
			transformers.WithPrimaryKeys("Id"),
		),
		Resolver: func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
			cl := meta.(*client.Client)
			return queryGetAll(
				"select id, project_resource__c, status__c, type__c, serial_number__c FROM fixed_asset__c",
				cl.SalesForce,
				res,
			)
		},
	}
}
