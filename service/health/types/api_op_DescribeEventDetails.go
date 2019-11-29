// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeEventDetailsInput struct {
	_ struct{} `type:"structure"`

	// A list of event ARNs (unique identifiers). For example: "arn:aws:health:us-east-1::event/EC2/EC2_INSTANCE_RETIREMENT_SCHEDULED/EC2_INSTANCE_RETIREMENT_SCHEDULED_ABC123-CDE456",
	// "arn:aws:health:us-west-1::event/EBS/AWS_EBS_LOST_VOLUME/AWS_EBS_LOST_VOLUME_CHI789_JKL101"
	//
	// EventArns is a required field
	EventArns []string `locationName:"eventArns" min:"1" type:"list" required:"true"`

	// The locale (language) to return information in. English (en) is the default
	// and the only supported value at this time.
	Locale *string `locationName:"locale" min:"2" type:"string"`
}

// String returns the string representation
func (s DescribeEventDetailsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEventDetailsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEventDetailsInput"}

	if s.EventArns == nil {
		invalidParams.Add(aws.NewErrParamRequired("EventArns"))
	}
	if s.EventArns != nil && len(s.EventArns) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EventArns", 1))
	}
	if s.Locale != nil && len(*s.Locale) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("Locale", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeEventDetailsOutput struct {
	_ struct{} `type:"structure"`

	// Error messages for any events that could not be retrieved.
	FailedSet []EventDetailsErrorItem `locationName:"failedSet" type:"list"`

	// Information about the events that could be retrieved.
	SuccessfulSet []EventDetails `locationName:"successfulSet" type:"list"`
}

// String returns the string representation
func (s DescribeEventDetailsOutput) String() string {
	return awsutil.Prettify(s)
}