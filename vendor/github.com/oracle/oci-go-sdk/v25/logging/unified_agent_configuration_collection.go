// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// loggingManagementControlplane API
//
// loggingManagementControlplane API specification
//

package logging

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// UnifiedAgentConfigurationCollection Results of a UnifiedAgentConfiguration search. Contains UnifiedAgentConfigurationSummary items
type UnifiedAgentConfigurationCollection struct {

	// List of UnifiedAgentConfigurationSummary.
	Items []UnifiedAgentConfigurationSummary `mandatory:"true" json:"items"`
}

func (m UnifiedAgentConfigurationCollection) String() string {
	return common.PointerString(m)
}
