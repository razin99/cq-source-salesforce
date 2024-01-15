package services

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/razin99/cq-source-salesforce/client"
)

type KrowLocationModel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func KrowLocationTable() *schema.Table {
	return &schema.Table{
		Name: "salesforce_krow_location",
		Transform: transformers.TransformWithStruct(
			&KrowLocationModel{},
			transformers.WithPrimaryKeys("Id"),
		),
		Resolver: func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
			cl := meta.(*client.Client)
			return queryGetAll(
				"select id, name from krow__location__c",
				cl.SalesForce,
				res,
			)
		},
	}
}
