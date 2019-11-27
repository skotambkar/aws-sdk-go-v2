// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListSecurityProfilesInput struct {
	_ struct{} `type:"structure"`

	// The identifier for your Amazon Connect instance. To find the ID of your instance,
	// open the AWS console and select Amazon Connect. Select the alias of the instance
	// in the Instance alias column. The instance ID is displayed in the Overview
	// section of your instance settings. For example, the instance ID is the set
	// of characters at the end of the instance ARN, after instance/, such as 10a4c4eb-f57e-4d4c-b602-bf39176ced07.
	//
	// InstanceId is a required field
	InstanceId *string `location:"uri" locationName:"InstanceId" min:"1" type:"string" required:"true"`

	// The maximum number of security profiles to return.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListSecurityProfilesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListSecurityProfilesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListSecurityProfilesInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListSecurityProfilesOutput struct {
	_ struct{} `type:"structure"`

	// A string returned in the response. Use the value returned in the response
	// as the value of the NextToken in a subsequent request to retrieve the next
	// set of results.
	NextToken *string `type:"string"`

	// An array of SecurityProfileSummary objects.
	SecurityProfileSummaryList []SecurityProfileSummary `type:"list"`
}

// String returns the string representation
func (s ListSecurityProfilesOutput) String() string {
	return awsutil.Prettify(s)
}
