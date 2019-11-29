// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteApplicationInputProcessingConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The Kinesis Analytics application name.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// The version ID of the Kinesis Analytics application.
	//
	// CurrentApplicationVersionId is a required field
	CurrentApplicationVersionId *int64 `min:"1" type:"long" required:"true"`

	// The ID of the input configuration from which to delete the input processing
	// configuration. You can get a list of the input IDs for an application by
	// using the DescribeApplication (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html)
	// operation.
	//
	// InputId is a required field
	InputId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteApplicationInputProcessingConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteApplicationInputProcessingConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteApplicationInputProcessingConfigurationInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}

	if s.CurrentApplicationVersionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CurrentApplicationVersionId"))
	}
	if s.CurrentApplicationVersionId != nil && *s.CurrentApplicationVersionId < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("CurrentApplicationVersionId", 1))
	}

	if s.InputId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputId"))
	}
	if s.InputId != nil && len(*s.InputId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InputId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteApplicationInputProcessingConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteApplicationInputProcessingConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}