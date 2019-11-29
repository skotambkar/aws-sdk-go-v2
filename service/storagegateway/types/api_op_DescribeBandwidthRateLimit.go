// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A JSON object containing the of the gateway.
type DescribeBandwidthRateLimitInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeBandwidthRateLimitInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeBandwidthRateLimitInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeBandwidthRateLimitInput"}

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

// A JSON object containing the following fields:
type DescribeBandwidthRateLimitOutput struct {
	_ struct{} `type:"structure"`

	// The average download bandwidth rate limit in bits per second. This field
	// does not appear in the response if the download rate limit is not set.
	AverageDownloadRateLimitInBitsPerSec *int64 `min:"102400" type:"long"`

	// The average upload bandwidth rate limit in bits per second. This field does
	// not appear in the response if the upload rate limit is not set.
	AverageUploadRateLimitInBitsPerSec *int64 `min:"51200" type:"long"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s DescribeBandwidthRateLimitOutput) String() string {
	return awsutil.Prettify(s)
}