// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/enums"
)

type ListTypeVersionsInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the type for which you want version summary
	// information.
	//
	// Conditional: You must specify TypeName or Arn.
	Arn *string `type:"string"`

	// The deprecation status of the type versions that you want to get summary
	// information about.
	//
	// Valid values include:
	//
	//    * LIVE: The type version is registered and can be used in CloudFormation
	//    operations, dependent on its provisioning behavior and visibility scope.
	//
	//    * DEPRECATED: The type version has been deregistered and can no longer
	//    be used in CloudFormation operations.
	DeprecatedStatus enums.DeprecatedStatus `type:"string" enum:"true"`

	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next
	// set of results.
	MaxResults *int64 `min:"1" type:"integer"`

	// If the previous paginated request didn't return all of the remaining results,
	// the response object's NextToken parameter value is set to a token. To retrieve
	// the next set of results, call this action again and assign that token to
	// the request object's NextToken parameter. If there are no remaining results,
	// the previous response object's NextToken parameter is set to null.
	NextToken *string `min:"1" type:"string"`

	// The kind of the type.
	//
	// Currently the only valid value is RESOURCE.
	Type enums.RegistryType `type:"string" enum:"true"`

	// The name of the type for which you want version summary information.
	//
	// Conditional: You must specify TypeName or Arn.
	TypeName *string `min:"10" type:"string"`
}

// String returns the string representation
func (s ListTypeVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTypeVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListTypeVersionsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.TypeName != nil && len(*s.TypeName) < 10 {
		invalidParams.Add(aws.NewErrParamMinLen("TypeName", 10))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListTypeVersionsOutput struct {
	_ struct{} `type:"structure"`

	// If the request doesn't return all of the remaining results, NextToken is
	// set to a token. To retrieve the next set of results, call this action again
	// and assign that token to the request object's NextToken parameter. If the
	// request returns all results, NextToken is set to null.
	NextToken *string `min:"1" type:"string"`

	// A list of TypeVersionSummary structures that contain information about the
	// specified type's versions.
	TypeVersionSummaries []TypeVersionSummary `type:"list"`
}

// String returns the string representation
func (s ListTypeVersionsOutput) String() string {
	return awsutil.Prettify(s)
}
