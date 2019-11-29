// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeFleetInstancesInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The filters.
	//
	//    * instance-type - The instance type.
	Filters []Filter `locationName:"Filter" locationNameList:"Filter" type:"list"`

	// The ID of the EC2 Fleet.
	//
	// FleetId is a required field
	FleetId *string `type:"string" required:"true"`

	// The maximum number of results to return in a single call. Specify a value
	// between 1 and 1000. The default value is 1000. To retrieve the remaining
	// results, make another call with the returned NextToken value.
	MaxResults *int64 `type:"integer"`

	// The token for the next set of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeFleetInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeFleetInstancesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeFleetInstancesInput"}

	if s.FleetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeFleetInstancesOutput struct {
	_ struct{} `type:"structure"`

	// The running instances. This list is refreshed periodically and might be out
	// of date.
	ActiveInstances []ActiveInstance `locationName:"activeInstanceSet" locationNameList:"item" type:"list"`

	// The ID of the EC2 Fleet.
	FleetId *string `locationName:"fleetId" type:"string"`

	// The token for the next set of results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeFleetInstancesOutput) String() string {
	return awsutil.Prettify(s)
}
