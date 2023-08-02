package virtualnetworkgatewayconnections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VngClientConnectionConfigurationProperties struct {
	ProvisioningState                 *ProvisioningState `json:"provisioningState,omitempty"`
	VirtualNetworkGatewayPolicyGroups []SubResource      `json:"virtualNetworkGatewayPolicyGroups"`
	VpnClientAddressPool              AddressSpace       `json:"vpnClientAddressPool"`
}
