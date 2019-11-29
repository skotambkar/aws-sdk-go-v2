// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input for the DescribeDeliveryChannels action.
type DescribeDeliveryChannelsInput struct {
	_ struct{} `type:"structure"`

	// A list of delivery channel names.
	DeliveryChannelNames []string `type:"list"`
}

// String returns the string representation
func (s DescribeDeliveryChannelsInput) String() string {
	return awsutil.Prettify(s)
}

// The output for the DescribeDeliveryChannels action.
type DescribeDeliveryChannelsOutput struct {
	_ struct{} `type:"structure"`

	// A list that contains the descriptions of the specified delivery channel.
	DeliveryChannels []DeliveryChannel `type:"list"`
}

// String returns the string representation
func (s DescribeDeliveryChannelsOutput) String() string {
	return awsutil.Prettify(s)
}
