// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteApplicationReferenceDataSourceInput struct {
	_ struct{} `type:"structure"`

	// Name of an existing application.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// Version of the application. You can use the DescribeApplication (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html)
	// operation to get the current application version. If the version specified
	// is not the current version, the ConcurrentModificationException is returned.
	//
	// CurrentApplicationVersionId is a required field
	CurrentApplicationVersionId *int64 `min:"1" type:"long" required:"true"`

	// ID of the reference data source. When you add a reference data source to
	// your application using the AddApplicationReferenceDataSource (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_AddApplicationReferenceDataSource.html),
	// Amazon Kinesis Analytics assigns an ID. You can use the DescribeApplication
	// (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html)
	// operation to get the reference ID.
	//
	// ReferenceId is a required field
	ReferenceId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteApplicationReferenceDataSourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteApplicationReferenceDataSourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteApplicationReferenceDataSourceInput"}

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

	if s.ReferenceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReferenceId"))
	}
	if s.ReferenceId != nil && len(*s.ReferenceId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ReferenceId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteApplicationReferenceDataSourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteApplicationReferenceDataSourceOutput) String() string {
	return awsutil.Prettify(s)
}
