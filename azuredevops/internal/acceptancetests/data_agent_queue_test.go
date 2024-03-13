//go:build (all || data_sources || data_agent_queue) && (!exclude_data_sources || !exclude_data_agent_queue)
// +build all data_sources data_agent_queue
// +build !exclude_data_sources !exclude_data_agent_queue

package acceptancetests

import (
	"testing"

	"github.com/Nataliia5722/terraform-provider-azuredevops/azuredevops/internal/acceptancetests/testutils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAgentQueue_DataSource(t *testing.T) {
	projectName := testutils.GenerateResourceName()
	agentQueueName := "Azure Pipelines"
	agentQueueData := testutils.HclAgentQueueDataSource(projectName, agentQueueName)

	tfNode := "data.azuredevops_agent_queue.queue"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testutils.PreCheck(t, nil) },
		Providers: testutils.GetProviders(),
		Steps: []resource.TestStep{
			{
				Config: agentQueueData,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(tfNode, "id"),
					resource.TestCheckResourceAttr(tfNode, "name", agentQueueName),
					resource.TestCheckResourceAttrSet(tfNode, "project_id"),
					resource.TestCheckResourceAttrSet(tfNode, "agent_pool_id"),
				),
			},
		},
	})
}
