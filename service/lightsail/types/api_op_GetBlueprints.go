// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetBlueprintsInput struct {
	_ struct{} `type:"structure"`

	// A Boolean value indicating whether to include inactive results in your request.
	IncludeInactive *bool `locationName:"includeInactive" type:"boolean"`

	// A token used for advancing to the next page of results from your get blueprints
	// request.
	PageToken *string `locationName:"pageToken" type:"string"`
}

// String returns the string representation
func (s GetBlueprintsInput) String() string {
	return awsutil.Prettify(s)
}

type GetBlueprintsOutput struct {
	_ struct{} `type:"structure"`

	// An array of key-value pairs that contains information about the available
	// blueprints.
	Blueprints []Blueprint `locationName:"blueprints" type:"list"`

	// A token used for advancing to the next page of results from your get blueprints
	// request.
	NextPageToken *string `locationName:"nextPageToken" type:"string"`
}

// String returns the string representation
func (s GetBlueprintsOutput) String() string {
	return awsutil.Prettify(s)
}
