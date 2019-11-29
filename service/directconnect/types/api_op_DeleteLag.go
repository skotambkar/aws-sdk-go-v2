// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/enums"
)

type DeleteLagInput struct {
	_ struct{} `type:"structure"`

	// The ID of the LAG.
	//
	// LagId is a required field
	LagId *string `locationName:"lagId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteLagInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteLagInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteLagInput"}

	if s.LagId == nil {
		invalidParams.Add(aws.NewErrParamRequired("LagId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Information about a link aggregation group (LAG).
type DeleteLagOutput struct {
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
func (s DeleteLagOutput) String() string {
	return awsutil.Prettify(s)
}