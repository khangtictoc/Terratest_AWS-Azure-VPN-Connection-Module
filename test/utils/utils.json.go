package utils

//+-------------------------------
//+ AWS VPN Connection
//+-------------------------------

type Telemetry struct {
	Status string `json:"Status"`
}

type Route struct {
	DestinationCidrBlock string `json:"destination_cidr_block"`
	SourceCidrBlock      string `json:"source"`
	State                string `json:"state"`
}

type AwsVpnConnection struct {
	Routes                       []Route `json:"routes"`
	CustomerGatewayConfiguration string  `json:"customer_gateway_configuration"`
}

//+-------------------------------
//+ Azure Virtual Network
//+-------------------------------

type AzureVnet struct {
	Id           string   `json:"id"`
	AddressSpace []string `json:"address_space"`
	Subnets      []Subnet `json:"subnets"`
}

type Subnet struct {
	Id            string `json:"id"`
	AddressPrefix string `json:"address_prefix"`
}

//+-------------------------------
//+ Azure Virtual Network Gateway
//+-------------------------------

type AzureVnetGateway struct {
	Id              string `json:"id"`
	PublicIpAddress string `json:"public_ip_address"`
	SubnetId        string `json:"subnet_id"`
}