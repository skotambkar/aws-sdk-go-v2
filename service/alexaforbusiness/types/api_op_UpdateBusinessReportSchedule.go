// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/enums"
)

type UpdateBusinessReportScheduleInput struct {
	_ struct{} `type:"structure"`

	// The format of the generated report (individual CSV files or zipped files
	// of individual files).
	Format enums.BusinessReportFormat `type:"string" enum:"true"`

	// The recurrence of the reports.
	Recurrence *BusinessReportRecurrence `type:"structure"`

	// The S3 location of the output reports.
	S3BucketName *string `type:"string"`

	// The S3 key where the report is delivered.
	S3KeyPrefix *string `type:"string"`

	// The ARN of the business report schedule.
	//
	// ScheduleArn is a required field
	ScheduleArn *string `type:"string" required:"true"`

	// The name identifier of the schedule.
	ScheduleName *string `type:"string"`
}

// String returns the string representation
func (s UpdateBusinessReportScheduleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateBusinessReportScheduleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateBusinessReportScheduleInput"}

	if s.ScheduleArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ScheduleArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateBusinessReportScheduleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateBusinessReportScheduleOutput) String() string {
	return awsutil.Prettify(s)
}