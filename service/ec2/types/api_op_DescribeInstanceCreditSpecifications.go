// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeInstanceCreditSpecificationsInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The filters.
	//
	//    * instance-id - The ID of the instance.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The instance IDs.
	//
	// Default: Describes all your instances.
	//
	// Constraints: Maximum 1000 explicitly specified instance IDs.
	InstanceIds []string `locationName:"InstanceId" locationNameList:"InstanceId" type:"list"`

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value. This
	// value can be between 5 and 1000. You cannot specify this parameter and the
	// instance IDs parameter in the same call.
	MaxResults *int64 `min:"5" type:"integer"`

	// The token to retrieve the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeInstanceCreditSpecificationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeInstanceCreditSpecificationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeInstanceCreditSpecificationsInput"}
	if s.MaxResults != nil && *s.MaxResults < 5 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 5))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeInstanceCreditSpecificationsOutput struct {
	_ struct{} `type:"structure"`

	// Information about the credit option for CPU usage of an instance.
	InstanceCreditSpecifications []InstanceCreditSpecification `locationName:"instanceCreditSpecificationSet" locationNameList:"item" type:"list"`

	// The token to use to retrieve the next page of results. This value is null
	// when there are no more results to return.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeInstanceCreditSpecificationsOutput) String() string {
	return awsutil.Prettify(s)
}
