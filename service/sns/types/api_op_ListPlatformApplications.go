// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Input for ListPlatformApplications action.
type ListPlatformApplicationsInput struct {
	_ struct{} `type:"structure"`

	// NextToken string is used when calling ListPlatformApplications action to
	// retrieve additional records that are available after the first page results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListPlatformApplicationsInput) String() string {
	return awsutil.Prettify(s)
}

// Response for ListPlatformApplications action.
type ListPlatformApplicationsOutput struct {
	_ struct{} `type:"structure"`

	// NextToken string is returned when calling ListPlatformApplications action
	// if additional records are available after the first page results.
	NextToken *string `type:"string"`

	// Platform applications returned when calling ListPlatformApplications action.
	PlatformApplications []PlatformApplication `type:"list"`
}

// String returns the string representation
func (s ListPlatformApplicationsOutput) String() string {
	return awsutil.Prettify(s)
}