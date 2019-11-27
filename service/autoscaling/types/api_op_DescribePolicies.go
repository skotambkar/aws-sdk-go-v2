// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribePoliciesInput struct {
	_ struct{} `type:"structure"`

	// The name of the Auto Scaling group.
	AutoScalingGroupName *string `min:"1" type:"string"`

	// The maximum number of items to be returned with each call. The default value
	// is 50 and the maximum value is 100.
	MaxRecords *int64 `type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`

	// The names of one or more policies. If you omit this parameter, all policies
	// are described. If a group name is provided, the results are limited to that
	// group. This list is limited to 50 items. If you specify an unknown policy
	// name, it is ignored with no error.
	PolicyNames []string `type:"list"`

	// One or more policy types. The valid values are SimpleScaling, StepScaling,
	// and TargetTrackingScaling.
	PolicyTypes []string `type:"list"`
}

// String returns the string representation
func (s DescribePoliciesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribePoliciesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribePoliciesInput"}
	if s.AutoScalingGroupName != nil && len(*s.AutoScalingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AutoScalingGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribePoliciesOutput struct {
	_ struct{} `type:"structure"`

	// A string that indicates that the response contains more items than can be
	// returned in a single response. To receive additional items, specify this
	// string for the NextToken value when requesting the next set of items. This
	// value is null when there are no more items to return.
	NextToken *string `type:"string"`

	// The scaling policies.
	ScalingPolicies []ScalingPolicy `type:"list"`
}

// String returns the string representation
func (s DescribePoliciesOutput) String() string {
	return awsutil.Prettify(s)
}
