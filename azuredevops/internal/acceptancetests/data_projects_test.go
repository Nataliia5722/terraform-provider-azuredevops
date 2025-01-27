//go:build (all || core || data_sources || resource_project || data_projects) && (!data_sources || !exclude_data_projects)
// +build all core data_sources resource_project data_projects
// +build !data_sources !exclude_data_projects

package acceptancetests

import (
	"testing"

	"github.com/Nataliia5722/terraform-provider-azuredevops/azuredevops/internal/acceptancetests/testutils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAzureDevOpsProjects_DataSource_SingleProject(t *testing.T) {
	projectName := testutils.GenerateResourceName()
	projectData := testutils.HclProjectsDataSource(projectName)

	tfNode := "data.azuredevops_projects.project-list"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testutils.PreCheck(t, nil) },
		ProviderFactories: testutils.GetProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: projectData,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tfNode, "projects.#", "1"),
				),
			},
		},
	})
}

func TestAccAzureDevOpsProjects_DataSource_EmptyResult(t *testing.T) {
	projectData := testutils.HclProjectsDataSourceWithStateAndInvalidName()

	tfNode := "data.azuredevops_projects.project-list"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testutils.PreCheck(t, nil) },
		Providers: testutils.GetProviders(),
		Steps: []resource.TestStep{
			{
				Config: projectData,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tfNode, "projects.#", "0"),
				),
			},
		},
	})
}
