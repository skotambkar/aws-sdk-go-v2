// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeParametersInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to include in the response. If more results
	// exist than the specified MaxResults value, a token is included in the response
	// so that the remaining results can be retrieved.
	//
	// The value for MaxResults must be between 20 and 100.
	MaxResults *int64 `type:"integer"`

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	NextToken *string `type:"string"`

	// The name of the parameter group.
	//
	// ParameterGroupName is a required field
	ParameterGroupName *string `type:"string" required:"true"`

	// How the parameter is defined. For example, system denotes a system-defined
	// parameter.
	Source *string `type:"string"`
}

// String returns the string representation
func (s DescribeParametersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeParametersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeParametersInput"}

	if s.ParameterGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParameterGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeParametersOutput struct {
	_ struct{} `type:"structure"`

	// Provides an identifier to allow retrieval of paginated results.
	NextToken *string `type:"string"`

	// A list of parameters within a parameter group. Each element in the list represents
	// one parameter.
	Parameters []Parameter `type:"list"`
}

// String returns the string representation
func (s DescribeParametersOutput) String() string {
	return awsutil.Prettify(s)
}
