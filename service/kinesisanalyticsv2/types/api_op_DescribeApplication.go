// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeApplicationInput struct {
	_ struct{} `type:"structure"`

	// The name of the application.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// Displays verbose information about a Kinesis Data Analytics application,
	// including the application's job plan.
	IncludeAdditionalDetails *bool `type:"boolean"`
}

// String returns the string representation
func (s DescribeApplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeApplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeApplicationInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeApplicationOutput struct {
	_ struct{} `type:"structure"`

	// Provides a description of the application, such as the application's Amazon
	// Resource Name (ARN), status, and latest version.
	//
	// ApplicationDetail is a required field
	ApplicationDetail *ApplicationDetail `type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeApplicationOutput) String() string {
	return awsutil.Prettify(s)
}