// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeVpcEndpointServicePermissionsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// One or more filters.
	//
	//    * principal - The ARN of the principal.
	//
	//    * principal-type - The principal type (All | Service | OrganizationUnit
	//    | Account | User | Role).
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The maximum number of results to return for the request in a single page.
	// The remaining results of the initial request can be seen by sending another
	// request with the returned NextToken value. This value can be between 5 and
	// 1000; if MaxResults is given a value larger than 1000, only 1000 results
	// are returned.
	MaxResults *int64 `type:"integer"`

	// The token to retrieve the next page of results.
	NextToken *string `type:"string"`

	// The ID of the service.
	//
	// ServiceId is a required field
	ServiceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeVpcEndpointServicePermissionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeVpcEndpointServicePermissionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeVpcEndpointServicePermissionsInput"}

	if s.ServiceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServiceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeVpcEndpointServicePermissionsOutput struct {
	_ struct{} `type:"structure"`

	// Information about one or more allowed principals.
	AllowedPrincipals []AllowedPrincipal `locationName:"allowedPrincipals" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeVpcEndpointServicePermissionsOutput) String() string {
	return awsutil.Prettify(s)
}
