// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type FilterLogEventsInput struct {
	_ struct{} `type:"structure"`

	// The end of the time range, expressed as the number of milliseconds after
	// Jan 1, 1970 00:00:00 UTC. Events with a timestamp later than this time are
	// not returned.
	EndTime *int64 `locationName:"endTime" type:"long"`

	// The filter pattern to use. For more information, see Filter and Pattern Syntax
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/FilterAndPatternSyntax.html).
	//
	// If not provided, all the events are matched.
	FilterPattern *string `locationName:"filterPattern" type:"string"`

	// If the value is true, the operation makes a best effort to provide responses
	// that contain events from multiple log streams within the log group, interleaved
	// in a single response. If the value is false, all the matched log events in
	// the first log stream are searched first, then those in the next log stream,
	// and so on. The default is false.
	//
	// IMPORTANT: Starting on June 17, 2019, this parameter will be ignored and
	// the value will be assumed to be true. The response from this operation will
	// always interleave events from multiple log streams within a log group.
	Interleaved *bool `locationName:"interleaved" deprecated:"true" type:"boolean"`

	// The maximum number of events to return. The default is 10,000 events.
	Limit *int64 `locationName:"limit" min:"1" type:"integer"`

	// The name of the log group to search.
	//
	// LogGroupName is a required field
	LogGroupName *string `locationName:"logGroupName" min:"1" type:"string" required:"true"`

	// Filters the results to include only events from log streams that have names
	// starting with this prefix.
	//
	// If you specify a value for both logStreamNamePrefix and logStreamNames, but
	// the value for logStreamNamePrefix does not match any log stream names specified
	// in logStreamNames, the action returns an InvalidParameterException error.
	LogStreamNamePrefix *string `locationName:"logStreamNamePrefix" min:"1" type:"string"`

	// Filters the results to only logs from the log streams in this list.
	//
	// If you specify a value for both logStreamNamePrefix and logStreamNames, the
	// action returns an InvalidParameterException error.
	LogStreamNames []string `locationName:"logStreamNames" min:"1" type:"list"`

	// The token for the next set of events to return. (You received this token
	// from a previous call.)
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`

	// The start of the time range, expressed as the number of milliseconds after
	// Jan 1, 1970 00:00:00 UTC. Events with a timestamp before this time are not
	// returned.
	StartTime *int64 `locationName:"startTime" type:"long"`
}

// String returns the string representation
func (s FilterLogEventsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *FilterLogEventsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "FilterLogEventsInput"}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}

	if s.LogGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LogGroupName"))
	}
	if s.LogGroupName != nil && len(*s.LogGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogGroupName", 1))
	}
	if s.LogStreamNamePrefix != nil && len(*s.LogStreamNamePrefix) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogStreamNamePrefix", 1))
	}
	if s.LogStreamNames != nil && len(s.LogStreamNames) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogStreamNames", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type FilterLogEventsOutput struct {
	_ struct{} `type:"structure"`

	// The matched events.
	Events []FilteredLogEvent `locationName:"events" type:"list"`

	// The token to use when requesting the next set of items. The token expires
	// after 24 hours.
	NextToken *string `locationName:"nextToken" min:"1" type:"string"`

	// Indicates which log streams have been searched and whether each has been
	// searched completely.
	SearchedLogStreams []SearchedLogStream `locationName:"searchedLogStreams" type:"list"`
}

// String returns the string representation
func (s FilterLogEventsOutput) String() string {
	return awsutil.Prettify(s)
}