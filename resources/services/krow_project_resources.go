package services

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/razin99/cq-source-salesforce/client"
)

type KrowProjectResourcesModel struct {
	Id                       string `json:"id"`
	Name                     string `json:"name"`
	User_Email__c            string `json:"user_email__c"`
	Employment_Start_Date__c string `json:"employment_start_date__c"`
	Employment_End_Date__c   string `json:"employment_end_date__c"`
	Legal_Name__c            string `json:"legal_name__c"`
	Krow__Team__c            string `json:"krow__team__c"`
	Krow__Active__c          bool   `json:"krow__active__c"`
	Krow__Location__c        string `json:"krow__location__c"`
}

func KrowProjectResourcesTable() *schema.Table {
	return &schema.Table{
		Name: "salesforce_krow_project_resources",
		Transform: transformers.TransformWithStruct(
			&KrowProjectResourcesModel{},
			transformers.WithPrimaryKeys("Id"),
		),
		Resolver: func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
			cl := meta.(*client.Client)
			return queryGetAll(
				"select id,name,user_email__c,employment_end_date__c,legal_name__c,employment_start_date__c,krow__team__c,krow__active__c,krow__location__c from krow__project_resources__c",
				cl.SalesForce,
				res,
			)
		},
	}
}
