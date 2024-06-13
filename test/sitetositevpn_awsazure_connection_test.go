package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"

	"encoding/json"
	"log"
)

type Telemetry struct {
	Status string `json:"status"`
}

func Test_SiteToSiteVPN_AWSAzure_Connection_Plan(t *testing.T) {
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../infra",
		VarFiles:     []string{"terraform.tfvars"}, // -var-file flag
		NoColor:      true,                         // Disable color in Terraform commands so its easier to parse stdout/stderr
	})

	terraform.InitAndPlan(t, terraformOptions)
	vpnConnectionOutput := terraform.OutputJson(t, terraformOptions, "vpn_s2s_aws_azure_debug")

}

func Test_SiteToSiteVPN_AWSAzure_Connection_Apply(t *testing.T) {
	//t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../infra",
		VarFiles:     []string{"terraform.tfvars"},
		NoColor:      true,
	})

	// Clean up resource at the end
	//defer terraform.Destroy(t, terraformOptions)

	// Test for "Init & Apply" phase
	// terraform.InitAndApply(t, terraformOptions)

	// Terraform Output
	vpnConnectionOutput := terraform.OutputJson(t, terraformOptions, "aws_vpn_connection_unittest")
	var tunnelStatusList []int

	var telemetries [][]Telemetry
	err := json.Unmarshal([]byte(vpnConnectionOutput), &telemetries)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	for _, telemetry := range telemetries {
		countUP := 0
		for _, t := range telemetry {
			if t.Status == "UP" {
				countUP++
			}
		}
		tunnelStatusList = append(tunnelStatusList, countUP)
	}

	fmt.Println(tunnelStatusList)

	////////////////////// UNIT TEST //////////////////////

	// Confirm all VPN connections have at least 1 tunnel UP
	for _, status := range tunnelStatusList {
		assert.True(t, status >= 1 && status <= 2, "At least 1 tunnel should be UP")
	}
}
