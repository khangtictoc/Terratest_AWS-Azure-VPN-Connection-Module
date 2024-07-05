package utils

import "encoding/xml"

//+-------------------------------
//+ AWS VPN Connection
//+-------------------------------

type VPNConnectionXML struct {
	XMLName     xml.Name         `xml:"vpn_connection"`
	IpsecTunnel []IpsecTunnelXML `xml:"ipsec_tunnel"`
}

type IpsecTunnelXML struct {
	XMLName         xml.Name           `xml:"ipsec_tunnel"`
	CustomerGateway CustomerGatewayXML `xml:"customer_gateway"`
}

type CustomerGatewayXML struct {
	XMLName              xml.Name                `xml:"customer_gateway"`
	TunnelOutsideAddress TunnelOutsideAddressXML `xml:"tunnel_outside_address"`
}

type TunnelOutsideAddressXML struct {
	XMLName   xml.Name `xml:"tunnel_outside_address"`
	IpAddress string   `xml:"ip_address"`
}
