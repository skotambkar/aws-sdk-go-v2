// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/enums"
)

type DescribeStackResourceDriftsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to be returned with a single call. If the number
	// of available results exceeds this maximum, the response includes a NextToken
	// value that you can assign to the NextToken request parameter to get the next
	// set of results.
	MaxResults *int64 `min:"1" type:"integer"`

	// A string that identifies the next page of stack resource drift results.
	NextToken *string `min:"1" type:"string"`

	// The name of the stack for which you want drift information.
	//
	// StackName is a required field
	StackName *string `min:"1" type:"string" required:"true"`

	// The resource drift status values to use as filters for the resource drift
	// results returned.
	//
	//    * DELETED: The resource differs from its expected template configuration
	//    in that the resource has been deleted.
	//
	//    * MODIFIED: One or more resource properties differ from their expected
	//    template values.
	//
	//    * IN_SYNC: The resources's actual configuration matches its expected template
	//    configuration.
	//
	//    * NOT_CHECKED: AWS CloudFormation does not currently return this value.
	StackResourceDriftStatusFilters []enums.StackResourceDriftStatus `min:"1" type:"list"`
}

// String returns the string representation
func (s DescribeStackResourceDriftsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeStackResourceDriftsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeStackResourceDriftsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.StackName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackName"))
	}
	if s.StackName != nil && len(*s.StackName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StackName", 1))
	}
	if s.StackResourceDriftStatusFilters != nil && len(s.StackResourceDriftStatusFilters) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StackResourceDriftStatusFilters", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeStackResourceDriftsOutput struct {
	_ struct{} `type:"structure"`

	// If the request doesn't return all of the remaining results, NextToken is
	// set to a token. To retrieve the next set of results, call DescribeStackResourceDrifts
	// again and assign that token to the request object's NextToken parameter.
	// If the request returns all results, NextToken is set to null.
	NextToken *string `min:"1" type:"string"`

	// Drift information for the resources that have been checked for drift in the
	// specified stack. This includes actual and expected configuration values for
	// resources where AWS CloudFormation detects drift.
	//
	// For a given stack, there will be one StackResourceDrift for each stack resource
	// that has been checked for drift. Resources that have not yet been checked
	// for drift are not included. Resources that do not currently support drift
	// detection are not checked, and so not included. For a list of resources that
	// support drift detection, see Resources that Support Drift Detection (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-stack-drift-resource-list.html).
	//
	// StackResourceDrifts is a required field
	StackResourceDrifts []StackResourceDrift `type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeStackResourceDriftsOutput) String() string {
	return awsutil.Prettify(s)
}