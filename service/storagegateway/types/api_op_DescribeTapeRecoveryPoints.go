// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// DescribeTapeRecoveryPointsInput
type DescribeTapeRecoveryPointsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`

	// Specifies that the number of virtual tape recovery points that are described
	// be limited to the specified number.
	Limit *int64 `min:"1" type:"integer"`

	// An opaque string that indicates the position at which to begin describing
	// the virtual tape recovery points.
	Marker *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeTapeRecoveryPointsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTapeRecoveryPointsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeTapeRecoveryPointsInput"}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// DescribeTapeRecoveryPointsOutput
type DescribeTapeRecoveryPointsOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string `min:"50" type:"string"`

	// An opaque string that indicates the position at which the virtual tape recovery
	// points that were listed for description ended.
	//
	// Use this marker in your next request to list the next set of virtual tape
	// recovery points in the list. If there are no more recovery points to describe,
	// this field does not appear in the response.
	Marker *string `min:"1" type:"string"`

	// An array of TapeRecoveryPointInfos that are available for the specified gateway.
	TapeRecoveryPointInfos []TapeRecoveryPointInfo `type:"list"`
}

// String returns the string representation
func (s DescribeTapeRecoveryPointsOutput) String() string {
	return awsutil.Prettify(s)
}