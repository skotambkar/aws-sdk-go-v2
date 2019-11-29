// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutTelemetryRecordsInput struct {
	_ struct{} `type:"structure"`

	EC2InstanceId *string `type:"string"`

	Hostname *string `type:"string"`

	ResourceARN *string `type:"string"`

	// TelemetryRecords is a required field
	TelemetryRecords []TelemetryRecord `type:"list" required:"true"`
}

// String returns the string representation
func (s PutTelemetryRecordsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutTelemetryRecordsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutTelemetryRecordsInput"}

	if s.TelemetryRecords == nil {
		invalidParams.Add(aws.NewErrParamRequired("TelemetryRecords"))
	}
	if s.TelemetryRecords != nil {
		for i, v := range s.TelemetryRecords {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "TelemetryRecords", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutTelemetryRecordsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutTelemetryRecordsOutput) String() string {
	return awsutil.Prettify(s)
}
