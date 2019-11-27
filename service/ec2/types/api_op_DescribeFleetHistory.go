// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ec2/enums"
)

type DescribeFleetHistoryInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The type of events to describe. By default, all events are described.
	EventType enums.FleetEventType `type:"string" enum:"true"`

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

	// The start date and time for the events, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ).
	//
	// StartTime is a required field
	StartTime *time.Time `type:"timestamp" required:"true"`
}

// String returns the string representation
func (s DescribeFleetHistoryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeFleetHistoryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeFleetHistoryInput"}

	if s.FleetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetId"))
	}

	if s.StartTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("StartTime"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeFleetHistoryOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the EC Fleet.
	FleetId *string `locationName:"fleetId" type:"string"`

	// Information about the events in the history of the EC2 Fleet.
	HistoryRecords []HistoryRecordEntry `locationName:"historyRecordSet" locationNameList:"item" type:"list"`

	// The last date and time for the events, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ).
	// All records up to this time were retrieved.
	//
	// If nextToken indicates that there are more results, this value is not present.
	LastEvaluatedTime *time.Time `locationName:"lastEvaluatedTime" type:"timestamp"`

	// The token for the next set of results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The start date and time for the events, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ).
	StartTime *time.Time `locationName:"startTime" type:"timestamp"`
}

// String returns the string representation
func (s DescribeFleetHistoryOutput) String() string {
	return awsutil.Prettify(s)
}
