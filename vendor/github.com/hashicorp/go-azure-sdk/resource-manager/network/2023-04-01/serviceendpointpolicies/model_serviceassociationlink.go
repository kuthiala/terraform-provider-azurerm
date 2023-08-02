package serviceendpointpolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceAssociationLink struct {
	Etag       *string                                 `json:"etag,omitempty"`
	Id         *string                                 `json:"id,omitempty"`
	Name       *string                                 `json:"name,omitempty"`
	Properties *ServiceAssociationLinkPropertiesFormat `json:"properties,omitempty"`
	Type       *string                                 `json:"type,omitempty"`
}
