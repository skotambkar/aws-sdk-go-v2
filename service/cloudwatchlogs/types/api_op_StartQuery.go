// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartQueryInput struct {
	_ struct{} `type:"structure"`

	// The end of the time range to query. The range is inclusive, so the specified
	// end time is included in the query. Specified as epoch time, the number of
	// seconds since January 1, 1970, 00:00:00 UTC.
	//
	// EndTime is a required field
	EndTime *int64 `locationName:"endTime" type:"long" required:"true"`

	// The maximum number of log events to return in the query. If the query string
	// uses the fields command, only the specified fields and their values are returned.
	// The default is 1000.
	Limit *int64 `locationName:"limit" min:"1" type:"integer"`

	// The log group on which to perform the query.
	//
	// A StartQuery operation must include a logGroupNames or a logGroupName parameter,
	// but not both.
	LogGroupName *string `locationName:"logGroupName" min:"1" type:"string"`

	// The list of log groups to be queried. You can include up to 20 log groups.
	//
	// A StartQuery operation must include a logGroupNames or a logGroupName parameter,
	// but not both.
	LogGroupNames []string `locationName:"logGroupNames" type:"list"`

	// The query string to use. For more information, see CloudWatch Logs Insights
	// Query Syntax (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html).
	//
	// QueryString is a required field
	QueryString *string `locationName:"queryString" type:"string" required:"true"`

	// The beginning of the time range to query. The range is inclusive, so the
	// specified start time is included in the query. Specified as epoch time, the
	// number of seconds since January 1, 1970, 00:00:00 UTC.
	//
	// StartTime is a required field
	StartTime *int64 `locationName:"startTime" type:"long" required:"true"`
}

// String returns the string representation
func (s StartQueryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartQueryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartQueryInput"}

	if s.EndTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndTime"))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.LogGroupName != nil && len(*s.LogGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogGroupName", 1))
	}

	if s.QueryString == nil {
		invalidParams.Add(aws.NewErrParamRequired("QueryString"))
	}

	if s.StartTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("StartTime"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartQueryOutput struct {
	_ struct{} `type:"structure"`

	// The unique ID of the query.
	QueryId *string `locationName:"queryId" type:"string"`
}

// String returns the string representation
func (s StartQueryOutput) String() string {
	return awsutil.Prettify(s)
}
