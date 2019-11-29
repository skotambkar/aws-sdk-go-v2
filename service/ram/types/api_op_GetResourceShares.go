// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ram/enums"
)

type GetResourceSharesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// The name of the resource share.
	Name *string `locationName:"name" type:"string"`

	// The token for the next page of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The type of owner.
	//
	// ResourceOwner is a required field
	ResourceOwner enums.ResourceOwner `locationName:"resourceOwner" type:"string" required:"true" enum:"true"`

	// The Amazon Resource Names (ARN) of the resource shares.
	ResourceShareArns []string `locationName:"resourceShareArns" type:"list"`

	// The status of the resource share.
	ResourceShareStatus enums.ResourceShareStatus `locationName:"resourceShareStatus" type:"string" enum:"true"`

	// One or more tag filters.
	TagFilters []TagFilter `locationName:"tagFilters" type:"list"`
}

// String returns the string representation
func (s GetResourceSharesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetResourceSharesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetResourceSharesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if len(s.ResourceOwner) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ResourceOwner"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetResourceSharesOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// Information about the resource shares.
	ResourceShares []ResourceShare `locationName:"resourceShares" type:"list"`
}

// String returns the string representation
func (s GetResourceSharesOutput) String() string {
	return awsutil.Prettify(s)
}
