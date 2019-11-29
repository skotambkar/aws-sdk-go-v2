// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AddApplicationInputInput struct {
	_ struct{} `type:"structure"`

	// Name of your existing Amazon Kinesis Analytics application to which you want
	// to add the streaming source.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// Current version of your Amazon Kinesis Analytics application. You can use
	// the DescribeApplication (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html)
	// operation to find the current application version.
	//
	// CurrentApplicationVersionId is a required field
	CurrentApplicationVersionId *int64 `min:"1" type:"long" required:"true"`

	// The Input (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_Input.html)
	// to add.
	//
	// Input is a required field
	Input *Input `type:"structure" required:"true"`
}

// String returns the string representation
func (s AddApplicationInputInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddApplicationInputInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddApplicationInputInput"}

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

	if s.Input == nil {
		invalidParams.Add(aws.NewErrParamRequired("Input"))
	}
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			invalidParams.AddNested("Input", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AddApplicationInputOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AddApplicationInputOutput) String() string {
	return awsutil.Prettify(s)
}