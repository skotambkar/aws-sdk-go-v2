// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type NotifyMigrationTaskStateInput struct {
	_ struct{} `type:"structure"`

	// Optional boolean flag to indicate whether any effect should take place. Used
	// to test if the caller has permission to make the call.
	DryRun *bool `type:"boolean"`

	// Unique identifier that references the migration task. Do not store personal
	// data in this field.
	//
	// MigrationTaskName is a required field
	MigrationTaskName *string `min:"1" type:"string" required:"true"`

	// Number of seconds after the UpdateDateTime within which the Migration Hub
	// can expect an update. If Migration Hub does not receive an update within
	// the specified interval, then the migration task will be considered stale.
	//
	// NextUpdateSeconds is a required field
	NextUpdateSeconds *int64 `type:"integer" required:"true"`

	// The name of the ProgressUpdateStream.
	//
	// ProgressUpdateStream is a required field
	ProgressUpdateStream *string `min:"1" type:"string" required:"true"`

	// Information about the task's progress and status.
	//
	// Task is a required field
	Task *Task `type:"structure" required:"true"`

	// The timestamp when the task was gathered.
	//
	// UpdateDateTime is a required field
	UpdateDateTime *time.Time `type:"timestamp" required:"true"`
}

// String returns the string representation
func (s NotifyMigrationTaskStateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *NotifyMigrationTaskStateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "NotifyMigrationTaskStateInput"}

	if s.MigrationTaskName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MigrationTaskName"))
	}
	if s.MigrationTaskName != nil && len(*s.MigrationTaskName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MigrationTaskName", 1))
	}

	if s.NextUpdateSeconds == nil {
		invalidParams.Add(aws.NewErrParamRequired("NextUpdateSeconds"))
	}

	if s.ProgressUpdateStream == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProgressUpdateStream"))
	}
	if s.ProgressUpdateStream != nil && len(*s.ProgressUpdateStream) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProgressUpdateStream", 1))
	}

	if s.Task == nil {
		invalidParams.Add(aws.NewErrParamRequired("Task"))
	}

	if s.UpdateDateTime == nil {
		invalidParams.Add(aws.NewErrParamRequired("UpdateDateTime"))
	}
	if s.Task != nil {
		if err := s.Task.Validate(); err != nil {
			invalidParams.AddNested("Task", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type NotifyMigrationTaskStateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s NotifyMigrationTaskStateOutput) String() string {
	return awsutil.Prettify(s)
}
