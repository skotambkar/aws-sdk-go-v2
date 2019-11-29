// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/enums"
)

type UpdateNetworkProfileInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the project for which you want to update
	// network profile settings.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`

	// The description of the network profile about which you are returning information.
	Description *string `locationName:"description" type:"string"`

	// The data throughput rate in bits per second, as an integer from 0 to 104857600.
	DownlinkBandwidthBits *int64 `locationName:"downlinkBandwidthBits" type:"long"`

	// Delay time for all packets to destination in milliseconds as an integer from
	// 0 to 2000.
	DownlinkDelayMs *int64 `locationName:"downlinkDelayMs" type:"long"`

	// Time variation in the delay of received packets in milliseconds as an integer
	// from 0 to 2000.
	DownlinkJitterMs *int64 `locationName:"downlinkJitterMs" type:"long"`

	// Proportion of received packets that fail to arrive from 0 to 100 percent.
	DownlinkLossPercent *int64 `locationName:"downlinkLossPercent" type:"integer"`

	// The name of the network profile about which you are returning information.
	Name *string `locationName:"name" type:"string"`

	// The type of network profile you wish to return information about. Valid values
	// are listed below.
	Type enums.NetworkProfileType `locationName:"type" type:"string" enum:"true"`

	// The data throughput rate in bits per second, as an integer from 0 to 104857600.
	UplinkBandwidthBits *int64 `locationName:"uplinkBandwidthBits" type:"long"`

	// Delay time for all packets to destination in milliseconds as an integer from
	// 0 to 2000.
	UplinkDelayMs *int64 `locationName:"uplinkDelayMs" type:"long"`

	// Time variation in the delay of received packets in milliseconds as an integer
	// from 0 to 2000.
	UplinkJitterMs *int64 `locationName:"uplinkJitterMs" type:"long"`

	// Proportion of transmitted packets that fail to arrive from 0 to 100 percent.
	UplinkLossPercent *int64 `locationName:"uplinkLossPercent" type:"integer"`
}

// String returns the string representation
func (s UpdateNetworkProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateNetworkProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateNetworkProfileInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateNetworkProfileOutput struct {
	_ struct{} `type:"structure"`

	// A list of the available network profiles.
	NetworkProfile *NetworkProfile `locationName:"networkProfile" type:"structure"`
}

// String returns the string representation
func (s UpdateNetworkProfileOutput) String() string {
	return awsutil.Prettify(s)
}