// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetInstancesInput struct {
	_ struct{} `type:"structure"`

	// A token used for advancing to the next page of results from your get instances
	// request.
	PageToken *string `locationName:"pageToken" type:"string"`
}

// String returns the string representation
func (s GetInstancesInput) String() string {
	return awsutil.Prettify(s)
}

type GetInstancesOutput struct {
	_ struct{} `type:"structure"`

	// An array of key-value pairs containing information about your instances.
	Instances []Instance `locationName:"instances" type:"list"`

	// A token used for advancing to the next page of results from your get instances
	// request.
	NextPageToken *string `locationName:"nextPageToken" type:"string"`
}

// String returns the string representation
func (s GetInstancesOutput) String() string {
	return awsutil.Prettify(s)
}