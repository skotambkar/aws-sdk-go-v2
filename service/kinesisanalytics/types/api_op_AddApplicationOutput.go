// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AddApplicationOutputInput struct {
	_ struct{} `type:"structure"`

	// Name of the application to which you want to add the output configuration.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// Version of the application to which you want to add the output configuration.
	// You can use the DescribeApplication (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html)
	// operation to get the current application version. If the version specified
	// is not the current version, the ConcurrentModificationException is returned.
	//
	// CurrentApplicationVersionId is a required field
	CurrentApplicationVersionId *int64 `min:"1" type:"long" required:"true"`

	// An array of objects, each describing one output configuration. In the output
	// configuration, you specify the name of an in-application stream, a destination
	// (that is, an Amazon Kinesis stream, an Amazon Kinesis Firehose delivery stream,
	// or an AWS Lambda function), and record the formation to use when writing
	// to the destination.
	//
	// Output is a required field
	Output *Output `type:"structure" required:"true"`
}

// String returns the string representation
func (s AddApplicationOutputInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddApplicationOutputInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddApplicationOutputInput"}

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

	if s.Output == nil {
		invalidParams.Add(aws.NewErrParamRequired("Output"))
	}
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
			invalidParams.AddNested("Output", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AddApplicationOutputOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AddApplicationOutputOutput) String() string {
	return awsutil.Prettify(s)
}
