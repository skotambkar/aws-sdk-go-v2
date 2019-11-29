// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input for the DescribeConfigurationRecorderStatus action.
type DescribeConfigurationRecorderStatusInput struct {
	_ struct{} `type:"structure"`

	// The name(s) of the configuration recorder. If the name is not specified,
	// the action returns the current status of all the configuration recorders
	// associated with the account.
	ConfigurationRecorderNames []string `type:"list"`
}

// String returns the string representation
func (s DescribeConfigurationRecorderStatusInput) String() string {
	return awsutil.Prettify(s)
}

// The output for the DescribeConfigurationRecorderStatus action, in JSON format.
type DescribeConfigurationRecorderStatusOutput struct {
	_ struct{} `type:"structure"`

	// A list that contains status of the specified recorders.
	ConfigurationRecordersStatus []ConfigurationRecorderStatus `type:"list"`
}

// String returns the string representation
func (s DescribeConfigurationRecorderStatusOutput) String() string {
	return awsutil.Prettify(s)
}
