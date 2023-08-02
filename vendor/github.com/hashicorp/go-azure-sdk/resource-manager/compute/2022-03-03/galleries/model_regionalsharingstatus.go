package galleries

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegionalSharingStatus struct {
	Details *string       `json:"details,omitempty"`
	Region  *string       `json:"region,omitempty"`
	State   *SharingState `json:"state,omitempty"`
}
