// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetMinuteUsageInput struct {
	_ struct{} `type:"structure"`

	// Month is a required field
	Month *int64 `locationName:"month" type:"integer" required:"true"`

	// Year is a required field
	Year *int64 `locationName:"year" type:"integer" required:"true"`
}

// String returns the string representation
func (s GetMinuteUsageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetMinuteUsageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetMinuteUsageInput"}

	if s.Month == nil {
		invalidParams.Add(aws.NewErrParamRequired("Month"))
	}

	if s.Year == nil {
		invalidParams.Add(aws.NewErrParamRequired("Year"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetMinuteUsageOutput struct {
	_ struct{} `type:"structure"`

	EstimatedMinutesRemaining *int64 `locationName:"estimatedMinutesRemaining" type:"integer"`

	IsReservedMinutesCustomer *bool `locationName:"isReservedMinutesCustomer" type:"boolean"`

	TotalReservedMinuteAllocation *int64 `locationName:"totalReservedMinuteAllocation" type:"integer"`

	TotalScheduledMinutes *int64 `locationName:"totalScheduledMinutes" type:"integer"`

	UpcomingMinutesScheduled *int64 `locationName:"upcomingMinutesScheduled" type:"integer"`
}

// String returns the string representation
func (s GetMinuteUsageOutput) String() string {
	return awsutil.Prettify(s)
}
