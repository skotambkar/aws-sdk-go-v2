// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetActiveNamesInput struct {
	_ struct{} `type:"structure"`

	// A token used for paginating results from your get active names request.
	PageToken *string `locationName:"pageToken" type:"string"`
}

// String returns the string representation
func (s GetActiveNamesInput) String() string {
	return awsutil.Prettify(s)
}

type GetActiveNamesOutput struct {
	_ struct{} `type:"structure"`

	// The list of active names returned by the get active names request.
	ActiveNames []string `locationName:"activeNames" type:"list"`

	// A token used for advancing to the next page of results from your get active
	// names request.
	NextPageToken *string `locationName:"nextPageToken" type:"string"`
}

// String returns the string representation
func (s GetActiveNamesOutput) String() string {
	return awsutil.Prettify(s)
}