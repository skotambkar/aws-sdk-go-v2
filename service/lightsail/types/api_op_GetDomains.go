// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetDomainsInput struct {
	_ struct{} `type:"structure"`

	// A token used for advancing to the next page of results from your get domains
	// request.
	PageToken *string `locationName:"pageToken" type:"string"`
}

// String returns the string representation
func (s GetDomainsInput) String() string {
	return awsutil.Prettify(s)
}

type GetDomainsOutput struct {
	_ struct{} `type:"structure"`

	// An array of key-value pairs containing information about each of the domain
	// entries in the user's account.
	Domains []Domain `locationName:"domains" type:"list"`

	// A token used for advancing to the next page of results from your get active
	// names request.
	NextPageToken *string `locationName:"nextPageToken" type:"string"`
}

// String returns the string representation
func (s GetDomainsOutput) String() string {
	return awsutil.Prettify(s)
}