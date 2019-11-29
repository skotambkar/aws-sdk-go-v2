// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A JSON object containing the following fields:
//
//    * DeleteBandwidthRateLimitInput$BandwidthType
type DeleteBandwidthRateLimitInput struct {
	_ struct{} `type:"structure"`

	// One of the BandwidthType values that indicates the gateway bandwidth rate
	// limit to delete.
	//
	// Valid Values: Upload, Download, All.
	//
	// BandwidthType is a required field
	BandwidthType *string `min:"3" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBandwidthRateLimitInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBandwidthRateLimitInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBandwidthRateLimitInput"}

	if s.BandwidthType == nil {
		invalidParams.Add(aws.NewErrParamRequired("BandwidthType"))
	}
	if s.BandwidthType != nil && len(*s.BandwidthType) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("BandwidthType", 3))
	}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A JSON object containing the of the gateway whose bandwidth rate information
// was deleted.
type DeleteBandwidthRateLimitOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s DeleteBandwidthRateLimitOutput) String() string {
	return awsutil.Prettify(s)
}