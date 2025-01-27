package repository

import (
	"github.com/Nataliia5722/azure-devops-go-api/azuredevops/v7/policy"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceRepositoryReservedNames() *schema.Resource {
	resource := genBasePolicyResource(&policyCrudArgs{
		FlattenFunc: reservedNamesFlattenFunc,
		ExpandFunc:  reservedNamesExpandFunc,
		PolicyType:  ReservedNames,
	})
	return resource
}

func reservedNamesFlattenFunc(d *schema.ResourceData, policyConfig *policy.PolicyConfiguration, projectID *string) error {
	err := baseFlattenFunc(d, policyConfig, projectID)
	if err != nil {
		return err
	}
	return nil
}

func reservedNamesExpandFunc(d *schema.ResourceData, typeID uuid.UUID) (*policy.PolicyConfiguration, *string, error) {
	policyConfig, projectID, err := baseExpandFunc(d, typeID)
	if err != nil {
		return nil, nil, err
	}
	return policyConfig, projectID, nil
}
