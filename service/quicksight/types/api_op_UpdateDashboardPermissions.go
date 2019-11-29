// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateDashboardPermissionsInput struct {
	_ struct{} `type:"structure"`

	// AWS account ID that contains the dashboard you are updating.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The ID for the dashboard.
	//
	// DashboardId is a required field
	DashboardId *string `location:"uri" locationName:"DashboardId" min:"1" type:"string" required:"true"`

	// The permissions that you want to grant on this resource.
	GrantPermissions []ResourcePermission `min:"1" type:"list"`

	// The permissions that you want to revoke from this resource.
	RevokePermissions []ResourcePermission `min:"1" type:"list"`
}

// String returns the string representation
func (s UpdateDashboardPermissionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDashboardPermissionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDashboardPermissionsInput"}

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
	if s.GrantPermissions != nil && len(s.GrantPermissions) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GrantPermissions", 1))
	}
	if s.RevokePermissions != nil && len(s.RevokePermissions) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RevokePermissions", 1))
	}
	if s.GrantPermissions != nil {
		for i, v := range s.GrantPermissions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "GrantPermissions", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.RevokePermissions != nil {
		for i, v := range s.RevokePermissions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "RevokePermissions", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateDashboardPermissionsOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the dashboard.
	DashboardArn *string `type:"string"`

	// The ID for the dashboard.
	DashboardId *string `min:"1" type:"string"`

	// Information about the permissions on the dashboard.
	Permissions []ResourcePermission `min:"1" type:"list"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The http status of the request.
	Status *int64 `location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s UpdateDashboardPermissionsOutput) String() string {
	return awsutil.Prettify(s)
}
