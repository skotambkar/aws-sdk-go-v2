// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/iot/enums"
)

type UpdateScheduledAuditInput struct {
	_ struct{} `type:"structure"`

	// The day of the month on which the scheduled audit takes place. Can be "1"
	// through "31" or "LAST". This field is required if the "frequency" parameter
	// is set to "MONTHLY". If days 29-31 are specified, and the month does not
	// have that many days, the audit takes place on the "LAST" day of the month.
	DayOfMonth *string `locationName:"dayOfMonth" type:"string"`

	// The day of the week on which the scheduled audit takes place. Can be one
	// of "SUN", "MON", "TUE", "WED", "THU", "FRI", or "SAT". This field is required
	// if the "frequency" parameter is set to "WEEKLY" or "BIWEEKLY".
	DayOfWeek enums.DayOfWeek `locationName:"dayOfWeek" type:"string" enum:"true"`

	// How often the scheduled audit takes place. Can be one of "DAILY", "WEEKLY",
	// "BIWEEKLY", or "MONTHLY". The start time of each audit is determined by the
	// system.
	Frequency enums.AuditFrequency `locationName:"frequency" type:"string" enum:"true"`

	// The name of the scheduled audit. (Max. 128 chars)
	//
	// ScheduledAuditName is a required field
	ScheduledAuditName *string `location:"uri" locationName:"scheduledAuditName" min:"1" type:"string" required:"true"`

	// Which checks are performed during the scheduled audit. Checks must be enabled
	// for your account. (Use DescribeAccountAuditConfiguration to see the list
	// of all checks, including those that are enabled or use UpdateAccountAuditConfiguration
	// to select which checks are enabled.)
	TargetCheckNames []string `locationName:"targetCheckNames" type:"list"`
}

// String returns the string representation
func (s UpdateScheduledAuditInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateScheduledAuditInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateScheduledAuditInput"}

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

type UpdateScheduledAuditOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the scheduled audit.
	ScheduledAuditArn *string `locationName:"scheduledAuditArn" type:"string"`
}

// String returns the string representation
func (s UpdateScheduledAuditOutput) String() string {
	return awsutil.Prettify(s)
}
