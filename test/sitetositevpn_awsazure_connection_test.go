package test

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"myterraform/utils"
	"os"
	"testing"

	// SSH Connection and Command Execution
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/ssh"
)

func Test_SiteToSiteVPN_AWSAzure_Connection_Apply(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../infra",
		VarFiles:     []string{"terraform.tfvars"},
		NoColor:      true,
	})

	// DESTROY
	//defer terraform.Destroy(t, terraformOptions)

	// CREATE
	terraform.InitAndApply(t, terraformOptions)
	// Wait for updating some configuration and figures in VPN Tunnel
	fmt.Println("WAITING FOR VPN TUNNEL TO BE UP AND RUNNING (15 MINUTES)...")
	//time.Sleep(20 * time.Minute)

	fmt.Println("+---------------------------------------------")
	fmt.Println("| [PREPARE] VPN Connection Status Testing")
	fmt.Println("+---------------------------------------------")

	vpnConnectionStatus := terraform.OutputJson(t, terraformOptions, "aws_vpn_connection_tunnel_status")

	var telemetryLst [][]utils.Telemetry
	errTelemetry := json.Unmarshal([]byte(vpnConnectionStatus), &telemetryLst)
	if errTelemetry != nil {
		log.Fatalf("Error parsing JSON for Telemetry: %v", errTelemetry)
	}

	var tunnelStatusList []int
	for _, telemetryObj := range telemetryLst {
		countUP := 0
		for _, t := range telemetryObj {
			if t.Status == "UP" {
				countUP++
			}
		}
		tunnelStatusList = append(tunnelStatusList, countUP)
	}

	var currentStatusCorrectList []int
	var targetStatusCorrectList []int
	for _, status := range tunnelStatusList {
		if status >= 1 && status <= 2 {
			currentStatusCorrectList = append(currentStatusCorrectList, 1)
		}
	}
	for i := 0; i < len(tunnelStatusList); i++ {
		targetStatusCorrectList = append(targetStatusCorrectList, 1)
	}

	fmt.Println("+----------------------------------------------------")
	fmt.Println("| [PREPRARE] VPN Static Routes Connections Testing")
	fmt.Println("+----------------------------------------------------")

	vpnConnection := terraform.OutputJson(t, terraformOptions, "aws_vpn_connection_info")
	azVnetGateway := terraform.OutputJson(t, terraformOptions, "azure_vnet_gateway_info")
	azVnet := terraform.OutputJson(t, terraformOptions, "azure_vnet_info")

	var vpnConnectionLst []utils.AwsVpnConnection
	errVpnConnection := json.Unmarshal([]byte(vpnConnection), &vpnConnectionLst)
	if errVpnConnection != nil {
		log.Fatalf("Error parsing JSON for VPN Connection: %v", errVpnConnection)
	}

	var azVnetGatewayObj []utils.AzureVnetGateway
	errAzVnetGateway := json.Unmarshal([]byte(azVnetGateway), &azVnetGatewayObj)
	if errAzVnetGateway != nil {
		log.Fatalf("Error parsing JSON for Azure Vnet Gateway: %v", errAzVnetGateway)
	}

	var azVnetObj []utils.AzureVnet
	errAzVnet := json.Unmarshal([]byte(azVnet), &azVnetObj)
	if errAzVnet != nil {
		log.Fatalf("Error parsing JSON for Azure Vnet: %v", errAzVnet)
	}

	var currentRouteCorrectList []int
	var targetRouteCorrectList []int
	for i := 0; i < len(vpnConnectionLst); i++ {
		targetRouteCorrectList = append(targetRouteCorrectList, 1)
	}

	// Confirm all VPN connections have at least 1 tunnel UP

	fmt.Println("+---------------------------------------------")
	fmt.Println("| [TEST] VPN Connection Status Testing")
	fmt.Println("+---------------------------------------------")

	fmt.Printf("Tunnel status list (UP: 1 - DOWN: 0): %v\n", tunnelStatusList)
	assert.Equal(t, targetStatusCorrectList, currentStatusCorrectList, "At least 1 tunnel UP for each VPN connection")

	// Confirm static routes in VPN connection that match wiuth Azure Vnet's address space

	fmt.Println("+----------------------------------------------------")
	fmt.Println("| [TEST] VPN Static Routes Connections Testing")
	fmt.Println("+----------------------------------------------------")

	for _, conn := range vpnConnectionLst {
		var vpnConnectionObj utils.VPNConnectionXML
		errVpnConnection = xml.Unmarshal([]byte(conn.CustomerGatewayConfiguration), &vpnConnectionObj)
		if errVpnConnection != nil {
			log.Fatalf("Error parsing XML for VPN Connection: %v", errVpnConnection)
		}

		for _, vgw := range azVnetGatewayObj {
			if vgw.PublicIpAddress == vpnConnectionObj.IpsecTunnel[0].CustomerGateway.TunnelOutsideAddress.IpAddress {
				for _, azVnet := range azVnetObj {
					for _, subnet := range azVnet.Subnets {
						if vgw.SubnetId == subnet.Id {
							if conn.Routes[0].DestinationCidrBlock == azVnet.AddressSpace[0] {
								currentRouteCorrectList = append(currentRouteCorrectList, 1)
							}
						}
					}
				}
			}
		}
	}

	fmt.Printf("Matched route list (MATCH: 1 - NOT-MATCH: 0): %v\n", currentRouteCorrectList)
	assert.Equal(t, currentRouteCorrectList, targetRouteCorrectList, "Static routes in VPN connection should match with Azure Vnet's address space")

	// Confirm 2 machines from different networks can communicate through PING

	fmt.Println("+----------------------------------------------------")
	fmt.Println("| [PREPARE] Ping test between 2 machines on AWS and Azure")
	fmt.Println("+----------------------------------------------------")

	awsInstancePublicIpsOutput := terraform.OutputList(t, terraformOptions, "aws_instance_public_ip_list")
	azureVmPrivateIpsOutput := terraform.OutputList(t, terraformOptions, "azure_vm_private_ip_list")

	var currentPingcorrectList []int
	var targetPingcorrectList []int
	for i := 0; i < len(azureVmPrivateIpsOutput); i++ {
		targetPingcorrectList = append(targetPingcorrectList, 1)
	}

	key, err := os.ReadFile("../infra/credentials/aws/private_key")
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: "ubuntu",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		// WARNING: InsecureIgnoreHostKey exposes the client to MITM attacks, use only for testing
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	fmt.Println("+----------------------------------------------------")
	fmt.Println("| [TEST] Ping test between 2 machines")
	fmt.Println("+----------------------------------------------------")

	for _, privIp := range azureVmPrivateIpsOutput {
		// Dial the server
		conn, err := ssh.Dial("tcp", awsInstancePublicIpsOutput[0]+":22", config)
		if err != nil {
			log.Fatalf("unable to connect: %v", err)
		}
		defer conn.Close()

		// Initalize a SSH connection
		session, err := conn.NewSession()
		if err != nil {
			log.Fatalf("unable to create session: %v", err)
		}
		defer session.Close()

		// Execute ping command (Default 5 packets are sent to test)
		output, err := session.CombinedOutput("ping -c 5 -W 1 " + privIp)
		_ = output
		if err != nil {
			fmt.Println("Error executing command:", err)
			currentPingcorrectList = append(currentPingcorrectList, 0)
		} else {
			fmt.Println("Ping test sucess to remote Azure host: " + privIp)
			currentPingcorrectList = append(currentPingcorrectList, 1)
		}
	}

	assert.Equal(t, currentPingcorrectList, targetPingcorrectList, "All hosts in AWS should be able to PING to all on Azure")
}
