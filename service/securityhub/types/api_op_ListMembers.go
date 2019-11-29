// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListMembersInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of items that you want in the response.
	MaxResults *int64 `location:"querystring" locationName:"MaxResults" min:"1" type:"integer"`

	// Paginates results. Set the value of this parameter to NULL on your first
	// call to the ListMembers operation. For subsequent calls to the operation,
	// fill nextToken in the request with the value of nextToken from the previous
	// response to continue listing data.
	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`

	// Specifies which member accounts the response includes based on their relationship
	// status with the master account. The default value is TRUE. If onlyAssociated
	// is set to TRUE, the response includes member accounts whose relationship
	// status with the master is set to ENABLED or DISABLED. If onlyAssociated is
	// set to FALSE, the response includes all existing member accounts.
	OnlyAssociated *bool `location:"querystring" locationName:"OnlyAssociated" type:"boolean"`
}

// String returns the string representation
func (s ListMembersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListMembersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListMembersInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListMembersOutput struct {
	_ struct{} `type:"structure"`

	// Member details returned by the operation.
	Members []Member `type:"list"`

	// The token that is required for pagination.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListMembersOutput) String() string {
	return awsutil.Prettify(s)
}
