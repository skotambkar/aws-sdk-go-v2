// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeLoadBasedAutoScalingInput struct {
	_ struct{} `type:"structure"`

	// An array of layer IDs.
	//
	// LayerIds is a required field
	LayerIds []string `type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeLoadBasedAutoScalingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeLoadBasedAutoScalingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeLoadBasedAutoScalingInput"}

	if s.LayerIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("LayerIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a DescribeLoadBasedAutoScaling request.
type DescribeLoadBasedAutoScalingOutput struct {
	_ struct{} `type:"structure"`

	// An array of LoadBasedAutoScalingConfiguration objects that describe each
	// layer's configuration.
	LoadBasedAutoScalingConfigurations []LoadBasedAutoScalingConfiguration `type:"list"`
}

// String returns the string representation
func (s DescribeLoadBasedAutoScalingOutput) String() string {
	return awsutil.Prettify(s)
}
