//go:build all || core || data_projects
// +build all core data_projects

package taskagent

// The tests in this file use the mock clients in mock_client.go to mock out
// the Azure DevOps client operations.

import (
	"context"
	"testing"

	"github.com/Nataliia5722/azure-devops-go-api/azuredevops/v7/taskagent"
	"github.com/Nataliia5722/terraform-provider-azuredevops/azdosdkmocks"
	"github.com/Nataliia5722/terraform-provider-azuredevops/azuredevops/internal/client"
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"
)

func TestDataSourceAgentPool_Read_TestAgentPoolNotFound(t *testing.T) {
	agentPoolListEmpty := []taskagent.TaskAgentPool{}
	name := "nonexistentAgentPool"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	taskAgentClient := azdosdkmocks.NewMockTaskagentClient(ctrl)
	clients := &client.AggregatedClient{
		TaskAgentClient: taskAgentClient,
		Ctx:             context.Background(),
	}

	taskAgentClient.
		EXPECT().
		GetAgentPools(clients.Ctx, taskagent.GetAgentPoolsArgs{
			PoolName: &name,
		}).
		Return(&agentPoolListEmpty, nil).
		Times(1)

	resourceData := schema.TestResourceDataRaw(t, DataAgentPool().Schema, nil)
	resourceData.Set("name", &name)
	err := dataSourceAgentPoolRead(resourceData, clients)
	require.Contains(t, err.Error(), "Unable to find agent pool")
}

func TestDataSourceAgentPool_Read_TestMultipleAgentPoolsFound(t *testing.T) {
	agentPoolList := []taskagent.TaskAgentPool{{}, {}}
	name := "multiplePools"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	taskAgentClient := azdosdkmocks.NewMockTaskagentClient(ctrl)
	clients := &client.AggregatedClient{
		TaskAgentClient: taskAgentClient,
		Ctx:             context.Background(),
	}

	taskAgentClient.
		EXPECT().
		GetAgentPools(clients.Ctx, taskagent.GetAgentPoolsArgs{
			PoolName: &name,
		}).
		Return(&agentPoolList, nil).
		Times(1)

	resourceData := schema.TestResourceDataRaw(t, DataAgentPool().Schema, nil)
	resourceData.Set("name", &name)
	err := dataSourceAgentPoolRead(resourceData, clients)
	require.Contains(t, err.Error(), "Found multiple agent pools for name")
}
