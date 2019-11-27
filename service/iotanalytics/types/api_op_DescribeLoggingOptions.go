// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeLoggingOptionsInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeLoggingOptionsInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeLoggingOptionsOutput struct {
	_ struct{} `type:"structure"`

	// The current settings of the AWS IoT Analytics logging options.
	LoggingOptions *LoggingOptions `locationName:"loggingOptions" type:"structure"`
}

// String returns the string representation
func (s DescribeLoggingOptionsOutput) String() string {
	return awsutil.Prettify(s)
}
