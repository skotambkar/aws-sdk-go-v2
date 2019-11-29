// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/enums"
)

type CreateLagInput struct {
	_ struct{} `type:"structure"`

	// The tags to associate with the automtically created LAGs.
	ChildConnectionTags []Tag `locationName:"childConnectionTags" min:"1" type:"list"`

	// The ID of an existing connection to migrate to the LAG.
	ConnectionId *string `locationName:"connectionId" type:"string"`

	// The bandwidth of the individual physical connections bundled by the LAG.
	// The possible values are 50Mbps, 100Mbps, 200Mbps, 300Mbps, 400Mbps, 500Mbps,
	// 1Gbps, 2Gbps, 5Gbps, and 10Gbps.
	//
	// ConnectionsBandwidth is a required field
	ConnectionsBandwidth *string `locationName:"connectionsBandwidth" type:"string" required:"true"`

	// The name of the LAG.
	//
	// LagName is a required field
	LagName *string `locationName:"lagName" type:"string" required:"true"`

	// The location for the LAG.
	//
	// Location is a required field
	Location *string `locationName:"location" type:"string" required:"true"`

	// The number of physical connections initially provisioned and bundled by the
	// LAG.
	//
	// NumberOfConnections is a required field
	NumberOfConnections *int64 `locationName:"numberOfConnections" type:"integer" required:"true"`

	// The name of the service provider associated with the LAG.
	ProviderName *string `locationName:"providerName" type:"string"`

	// The tags to associate with the LAG.
	Tags []Tag `locationName:"tags" min:"1" type:"list"`
}

// String returns the string representation
func (s CreateLagInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateLagInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateLagInput"}
	if s.ChildConnectionTags != nil && len(s.ChildConnectionTags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChildConnectionTags", 1))
	}

	if s.ConnectionsBandwidth == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectionsBandwidth"))
	}

	if s.LagName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LagName"))
	}

	if s.Location == nil {
		invalidParams.Add(aws.NewErrParamRequired("Location"))
	}

	if s.NumberOfConnections == nil {
		invalidParams.Add(aws.NewErrParamRequired("NumberOfConnections"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.ChildConnectionTags != nil {
		for i, v := range s.ChildConnectionTags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ChildConnectionTags", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about a link aggregation group (LAG).
type CreateLagOutput struct {
	_ struct{} `type:"structure"`

	// Indicates whether the LAG can host other connections.
	AllowsHostedConnections *bool `locationName:"allowsHostedConnections" type:"boolean"`

	// The AWS Direct Connect endpoint that hosts the LAG.
	AwsDevice *string `locationName:"awsDevice" deprecated:"true" type:"string"`

	// The AWS Direct Connect endpoint that hosts the LAG.
	AwsDeviceV2 *string `locationName:"awsDeviceV2" type:"string"`

	// The connections bundled by the LAG.
	Connections []Connection `locationName:"connections" type:"list"`

	// The individual bandwidth of the physical connections bundled by the LAG.
	// The possible values are 1Gbps and 10Gbps.
	ConnectionsBandwidth *string `locationName:"connectionsBandwidth" type:"string"`

	// Indicates whether the LAG supports a secondary BGP peer in the same address
	// family (IPv4/IPv6).
	HasLogicalRedundancy enums.HasLogicalRedundancy `locationName:"hasLogicalRedundancy" type:"string" enum:"true"`

	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable *bool `locationName:"jumboFrameCapable" type:"boolean"`

	// The ID of the LAG.
	LagId *string `locationName:"lagId" type:"string"`

	// The name of the LAG.
	LagName *string `locationName:"lagName" type:"string"`

	// The state of the LAG. The following are the possible values:
	//
	//    * requested: The initial state of a LAG. The LAG stays in the requested
	//    state until the Letter of Authorization (LOA) is available.
	//
	//    * pending: The LAG has been approved and is being initialized.
	//
	//    * available: The network link is established and the LAG is ready for
	//    use.
	//
	//    * down: The network link is down.
	//
	//    * deleting: The LAG is being deleted.
	//
	//    * deleted: The LAG is deleted.
	//
	//    * unknown: The state of the LAG is not available.
	LagState enums.LagState `locationName:"lagState" type:"string" enum:"true"`

	// The location of the LAG.
	Location *string `locationName:"location" type:"string"`

	// The minimum number of physical connections that must be operational for the
	// LAG itself to be operational.
	MinimumLinks *int64 `locationName:"minimumLinks" type:"integer"`

	// The number of physical connections bundled by the LAG, up to a maximum of
	// 10.
	NumberOfConnections *int64 `locationName:"numberOfConnections" type:"integer"`

	// The ID of the AWS account that owns the LAG.
	OwnerAccount *string `locationName:"ownerAccount" type:"string"`

	// The name of the service provider associated with the LAG.
	ProviderName *string `locationName:"providerName" type:"string"`

	// The AWS Region where the connection is located.
	Region *string `locationName:"region" type:"string"`

	// The tags associated with the LAG.
	Tags []Tag `locationName:"tags" min:"1" type:"list"`
}

// String returns the string representation
func (s CreateLagOutput) String() string {
	return awsutil.Prettify(s)
}
