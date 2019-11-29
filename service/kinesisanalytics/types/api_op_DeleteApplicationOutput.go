// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteApplicationOutputInput struct {
	_ struct{} `type:"structure"`

	// Amazon Kinesis Analytics application name.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// Amazon Kinesis Analytics application version. You can use the DescribeApplication
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html)
	// operation to get the current application version. If the version specified
	// is not the current version, the ConcurrentModificationException is returned.
	//
	// CurrentApplicationVersionId is a required field
	CurrentApplicationVersionId *int64 `min:"1" type:"long" required:"true"`

	// The ID of the configuration to delete. Each output configuration that is
	// added to the application, either when the application is created or later
	// using the AddApplicationOutput (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_AddApplicationOutput.html)
	// operation, has a unique ID. You need to provide the ID to uniquely identify
	// the output configuration that you want to delete from the application configuration.
	// You can use the DescribeApplication (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html)
	// operation to get the specific OutputId.
	//
	// OutputId is a required field
	OutputId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteApplicationOutputInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteApplicationOutputInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteApplicationOutputInput"}

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

	if s.OutputId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OutputId"))
	}
	if s.OutputId != nil && len(*s.OutputId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OutputId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteApplicationOutputOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteApplicationOutputOutput) String() string {
	return awsutil.Prettify(s)
}