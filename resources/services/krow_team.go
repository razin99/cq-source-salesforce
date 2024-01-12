package services

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/razin99/cq-source-salesforce/client"
)

type KrowTeamModel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func KrowTeamTable() *schema.Table {
	return &schema.Table{
		Name:      "salesforce_krow_team",
		Transform: transformers.TransformWithStruct(&KrowTeamModel{}, transformers.WithPrimaryKeys("ID")),
		Resolver: func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
			cl := meta.(*client.Client)
			return queryGetAll("select id, name from krow__team__c", cl.SalesForce, res)
		},
	}
}
