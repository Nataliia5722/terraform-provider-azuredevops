//go:build (all || resource_policy_path_lenght) && !resource_policy_path_lenght
// +build all resource_policy_path_lenght
// +build !resource_policy_path_lenght

package acceptancetests

import (
	"fmt"
	"testing"

	"github.com/Nataliia5722/terraform-provider-azuredevops/azuredevops/internal/acceptancetests/testutils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const pathLengthTfNode = "azuredevops_repository_policy_max_path_length.p"

func TestAccPolicyPathLength(t *testing.T) {
	testutils.RunTestsInSequence(t, map[string]map[string]func(t *testing.T){
		"RepositoryPolicies": {
			"basic":  testAccRepoPolicyPathLengthBasic,
			"update": testAccRepoPolicyPathLengthUpdate,
		},
		"ProjectPolicies": {
			"basic":  testAccProjectPolicyPathLengthBasic,
			"update": testAccProjectPolicyPathLengthUpdate,
		},
	})
}

func testAccRepoPolicyPathLengthBasic(t *testing.T) {
	projectName := testutils.GenerateResourceName()
	repoName := testutils.GenerateResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testutils.PreCheck(t, nil) },
		Providers: testutils.GetProviders(),
		Steps: []resource.TestStep{
			{
				Config: hclRepoPolicyPathLengthBasic(projectName, repoName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(pathLengthTfNode, "enabled", "true"),
					resource.TestCheckResourceAttr(pathLengthTfNode, "max_path_length", "500"),
				),
			}, {
				ResourceName:      pathLengthTfNode,
				ImportStateIdFunc: testutils.ComputeProjectQualifiedResourceImportID(pathLengthTfNode),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccRepoPolicyPathLengthUpdate(t *testing.T) {
	projectName := testutils.GenerateResourceName()
	repoName := testutils.GenerateResourceName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testutils.PreCheck(t, nil) },
		Providers: testutils.GetProviders(),
		Steps: []resource.TestStep{
			{
				Config: hclRepoPolicyPathLengthBasic(projectName, repoName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(pathLengthTfNode, "enabled", "true"),
				),
			}, {
				Config: hclRepoPolicyPathLengthUpdate(projectName, repoName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(pathLengthTfNode, "enabled", "true"),
					resource.TestCheckResourceAttr(pathLengthTfNode, "max_path_length", "1000"),
				),
			}, {
				ResourceName:      pathLengthTfNode,
				ImportStateIdFunc: testutils.ComputeProjectQualifiedResourceImportID(pathLengthTfNode),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccProjectPolicyPathLengthBasic(t *testing.T) {
	projectName := testutils.GenerateResourceName()
	repoName := testutils.GenerateResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testutils.PreCheck(t, nil) },
		Providers: testutils.GetProviders(),
		Steps: []resource.TestStep{
			{
				Config: hclProjectPolicyPathLengthBasic(projectName, repoName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(pathLengthTfNode, "enabled", "true"),
					resource.TestCheckResourceAttr(pathLengthTfNode, "max_path_length", "500"),
				),
			}, {
				ResourceName:      pathLengthTfNode,
				ImportStateIdFunc: testutils.ComputeProjectQualifiedResourceImportID(pathLengthTfNode),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccProjectPolicyPathLengthUpdate(t *testing.T) {
	projectName := testutils.GenerateResourceName()
	repoName := testutils.GenerateResourceName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testutils.PreCheck(t, nil) },
		Providers: testutils.GetProviders(),
		Steps: []resource.TestStep{
			{
				Config: hclProjectPolicyPathLengthBasic(projectName, repoName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(pathLengthTfNode, "enabled", "true"),
				),
			}, {
				Config: hclProjectPolicyPathLengthUpdate(projectName, repoName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(pathLengthTfNode, "enabled", "true"),
					resource.TestCheckResourceAttr(pathLengthTfNode, "max_path_length", "1000"),
				),
			}, {
				ResourceName:      pathLengthTfNode,
				ImportStateIdFunc: testutils.ComputeProjectQualifiedResourceImportID(pathLengthTfNode),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func hclPolicyPathLengthResourceTemplate(projectName string, repoName string) string {
	return fmt.Sprintf(`
resource "azuredevops_project" "p" {
  name               = "%s"
  description        = "Test Project Description"
  visibility         = "private"
  version_control    = "Git"
  work_item_template = "Agile"
}

resource "azuredevops_git_repository" "r" {
  project_id = azuredevops_project.p.id
  name       = "%s"
  initialization {
    init_type = "Clean"
  }
}
`, projectName, repoName)
}

func hclRepoPolicyPathLengthBasic(projectName string, repoName string) string {
	projectAndRepo := hclPolicyPathLengthResourceTemplate(projectName, repoName)
	return fmt.Sprintf(`%s %s`, projectAndRepo, `
resource "azuredevops_repository_policy_max_path_length" "p" {
  project_id = azuredevops_project.p.id
  enabled  = true
  blocking = true
  max_path_length = 500
  repository_ids  = [azuredevops_git_repository.r.id]
}
`)
}

func hclRepoPolicyPathLengthUpdate(projectName string, repoName string) string {
	projectAndRepo := hclPolicyPathLengthResourceTemplate(projectName, repoName)
	return fmt.Sprintf(`%s %s`, projectAndRepo, `
resource "azuredevops_repository_policy_max_path_length" "p" {
  project_id = azuredevops_project.p.id
  enabled  = true
  blocking = true
  max_path_length = 1000
  repository_ids  = [azuredevops_git_repository.r.id]
}
`)
}

func hclProjectPolicyPathLengthBasic(projectName string, repoName string) string {
	projectAndRepo := hclPolicyPathLengthResourceTemplate(projectName, repoName)
	return fmt.Sprintf(`%s %s`, projectAndRepo, `
resource "azuredevops_repository_policy_max_path_length" "p" {
  project_id = azuredevops_project.p.id
  enabled  = true
  blocking = true
  max_path_length = 500
  depends_on = [azuredevops_git_repository.r]
}
`)
}

func hclProjectPolicyPathLengthUpdate(projectName string, repoName string) string {
	projectAndRepo := hclPolicyPathLengthResourceTemplate(projectName, repoName)
	return fmt.Sprintf(`%s %s`, projectAndRepo, `
resource "azuredevops_repository_policy_max_path_length" "p" {
  project_id = azuredevops_project.p.id
  enabled  = true
  blocking = true
  max_path_length = 1000
  depends_on = [azuredevops_git_repository.r]
}
`)
}
