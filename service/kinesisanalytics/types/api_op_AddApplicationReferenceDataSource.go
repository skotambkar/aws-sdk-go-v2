// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AddApplicationReferenceDataSourceInput struct {
	_ struct{} `type:"structure"`

	// Name of an existing application.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// Version of the application for which you are adding the reference data source.
	// You can use the DescribeApplication (https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_DescribeApplication.html)
	// operation to get the current application version. If the version specified
	// is not the current version, the ConcurrentModificationException is returned.
	//
	// CurrentApplicationVersionId is a required field
	CurrentApplicationVersionId *int64 `min:"1" type:"long" required:"true"`

	// The reference data source can be an object in your Amazon S3 bucket. Amazon
	// Kinesis Analytics reads the object and copies the data into the in-application
	// table that is created. You provide an S3 bucket, object key name, and the
	// resulting in-application table that is created. You must also provide an
	// IAM role with the necessary permissions that Amazon Kinesis Analytics can
	// assume to read the object from your S3 bucket on your behalf.
	//
	// ReferenceDataSource is a required field
	ReferenceDataSource *ReferenceDataSource `type:"structure" required:"true"`
}

// String returns the string representation
func (s AddApplicationReferenceDataSourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddApplicationReferenceDataSourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddApplicationReferenceDataSourceInput"}

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

	if s.ReferenceDataSource == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReferenceDataSource"))
	}
	if s.ReferenceDataSource != nil {
		if err := s.ReferenceDataSource.Validate(); err != nil {
			invalidParams.AddNested("ReferenceDataSource", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AddApplicationReferenceDataSourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AddApplicationReferenceDataSourceOutput) String() string {
	return awsutil.Prettify(s)
}
