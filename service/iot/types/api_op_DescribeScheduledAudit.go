// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/iot/enums"
)

type DescribeScheduledAuditInput struct {
	_ struct{} `type:"structure"`

	// The name of the scheduled audit whose information you want to get.
	//
	// ScheduledAuditName is a required field
	ScheduledAuditName *string `location:"uri" locationName:"scheduledAuditName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeScheduledAuditInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeScheduledAuditInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeScheduledAuditInput"}

	if s.ScheduledAuditName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ScheduledAuditName"))
	}
	if s.ScheduledAuditName != nil && len(*s.ScheduledAuditName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ScheduledAuditName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeScheduledAuditOutput struct {
	_ struct{} `type:"structure"`

	// The day of the month on which the scheduled audit takes place. Will be "1"
	// through "31" or "LAST". If days 29-31 are specified, and the month does not
	// have that many days, the audit takes place on the "LAST" day of the month.
	DayOfMonth *string `locationName:"dayOfMonth" type:"string"`

	// The day of the week on which the scheduled audit takes place. One of "SUN",
	// "MON", "TUE", "WED", "THU", "FRI", or "SAT".
	DayOfWeek enums.DayOfWeek `locationName:"dayOfWeek" type:"string" enum:"true"`

	// How often the scheduled audit takes place. One of "DAILY", "WEEKLY", "BIWEEKLY",
	// or "MONTHLY". The start time of each audit is determined by the system.
	Frequency enums.AuditFrequency `locationName:"frequency" type:"string" enum:"true"`

	// The ARN of the scheduled audit.
	ScheduledAuditArn *string `locationName:"scheduledAuditArn" type:"string"`

	// The name of the scheduled audit.
	ScheduledAuditName *string `locationName:"scheduledAuditName" min:"1" type:"string"`

	// Which checks are performed during the scheduled audit. Checks must be enabled
	// for your account. (Use DescribeAccountAuditConfiguration to see the list
	// of all checks, including those that are enabled or use UpdateAccountAuditConfiguration
	// to select which checks are enabled.)
	TargetCheckNames []string `locationName:"targetCheckNames" type:"list"`
}

// String returns the string representation
func (s DescribeScheduledAuditOutput) String() string {
	return awsutil.Prettify(s)
}