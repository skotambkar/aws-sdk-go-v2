// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListBusinessReportSchedulesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of schedules listed in the call.
	MaxResults *int64 `min:"1" type:"integer"`

	// The token used to list the remaining schedules from the previous API call.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListBusinessReportSchedulesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListBusinessReportSchedulesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListBusinessReportSchedulesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListBusinessReportSchedulesOutput struct {
	_ struct{} `type:"structure"`

	// The schedule of the reports.
	BusinessReportSchedules []BusinessReportSchedule `type:"list"`

	// The token used to list the remaining schedules from the previous API call.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListBusinessReportSchedulesOutput) String() string {
	return awsutil.Prettify(s)
}
