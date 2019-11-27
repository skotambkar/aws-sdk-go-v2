// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the inputs for the GetDirectoryLimits operation.
type GetDirectoryLimitsInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetDirectoryLimitsInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the results of the GetDirectoryLimits operation.
type GetDirectoryLimitsOutput struct {
	_ struct{} `type:"structure"`

	// A DirectoryLimits object that contains the directory limits for the current
	// region.
	DirectoryLimits *DirectoryLimits `type:"structure"`
}

// String returns the string representation
func (s GetDirectoryLimitsOutput) String() string {
	return awsutil.Prettify(s)
}
