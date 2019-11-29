// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateDashboardPublishedVersionInput struct {
	_ struct{} `type:"structure"`

	// AWS account ID that contains the dashboard you are updating.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The ID for the dashboard.
	//
	// DashboardId is a required field
	DashboardId *string `location:"uri" locationName:"DashboardId" min:"1" type:"string" required:"true"`

	// The version number of the dashboard.
	//
	// VersionNumber is a required field
	VersionNumber *int64 `location:"uri" locationName:"VersionNumber" min:"1" type:"long" required:"true"`
}

// String returns the string representation
func (s UpdateDashboardPublishedVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDashboardPublishedVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDashboardPublishedVersionInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.DashboardId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DashboardId"))
	}
	if s.DashboardId != nil && len(*s.DashboardId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DashboardId", 1))
	}

	if s.VersionNumber == nil {
		invalidParams.Add(aws.NewErrParamRequired("VersionNumber"))
	}
	if s.VersionNumber != nil && *s.VersionNumber < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("VersionNumber", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateDashboardPublishedVersionOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the dashboard.
	DashboardArn *string `type:"string"`

	// The ID for the dashboard.
	DashboardId *string `min:"1" type:"string"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The http status of the request.
	Status *int64 `location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s UpdateDashboardPublishedVersionOutput) String() string {
	return awsutil.Prettify(s)
}
