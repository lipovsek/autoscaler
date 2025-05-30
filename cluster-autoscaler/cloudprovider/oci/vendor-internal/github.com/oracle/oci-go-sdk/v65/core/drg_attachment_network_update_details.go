// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"encoding/json"
	"fmt"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/oci/vendor-internal/github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DrgAttachmentNetworkUpdateDetails The representation of DrgAttachmentNetworkUpdateDetails
type DrgAttachmentNetworkUpdateDetails interface {
}

type drgattachmentnetworkupdatedetails struct {
	JsonData []byte
	Type     string `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *drgattachmentnetworkupdatedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdrgattachmentnetworkupdatedetails drgattachmentnetworkupdatedetails
	s := struct {
		Model Unmarshalerdrgattachmentnetworkupdatedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *drgattachmentnetworkupdatedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "VCN":
		mm := VcnDrgAttachmentNetworkUpdateDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for DrgAttachmentNetworkUpdateDetails: %s.", m.Type)
		return *m, nil
	}
}

func (m drgattachmentnetworkupdatedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m drgattachmentnetworkupdatedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DrgAttachmentNetworkUpdateDetailsTypeEnum Enum with underlying type: string
type DrgAttachmentNetworkUpdateDetailsTypeEnum string

// Set of constants representing the allowable values for DrgAttachmentNetworkUpdateDetailsTypeEnum
const (
	DrgAttachmentNetworkUpdateDetailsTypeVcn DrgAttachmentNetworkUpdateDetailsTypeEnum = "VCN"
)

var mappingDrgAttachmentNetworkUpdateDetailsTypeEnum = map[string]DrgAttachmentNetworkUpdateDetailsTypeEnum{
	"VCN": DrgAttachmentNetworkUpdateDetailsTypeVcn,
}

var mappingDrgAttachmentNetworkUpdateDetailsTypeEnumLowerCase = map[string]DrgAttachmentNetworkUpdateDetailsTypeEnum{
	"vcn": DrgAttachmentNetworkUpdateDetailsTypeVcn,
}

// GetDrgAttachmentNetworkUpdateDetailsTypeEnumValues Enumerates the set of values for DrgAttachmentNetworkUpdateDetailsTypeEnum
func GetDrgAttachmentNetworkUpdateDetailsTypeEnumValues() []DrgAttachmentNetworkUpdateDetailsTypeEnum {
	values := make([]DrgAttachmentNetworkUpdateDetailsTypeEnum, 0)
	for _, v := range mappingDrgAttachmentNetworkUpdateDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDrgAttachmentNetworkUpdateDetailsTypeEnumStringValues Enumerates the set of values in String for DrgAttachmentNetworkUpdateDetailsTypeEnum
func GetDrgAttachmentNetworkUpdateDetailsTypeEnumStringValues() []string {
	return []string{
		"VCN",
	}
}

// GetMappingDrgAttachmentNetworkUpdateDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDrgAttachmentNetworkUpdateDetailsTypeEnum(val string) (DrgAttachmentNetworkUpdateDetailsTypeEnum, bool) {
	enum, ok := mappingDrgAttachmentNetworkUpdateDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
