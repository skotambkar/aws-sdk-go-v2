// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateGroupInput struct {
	_ struct{} `type:"structure"`

	// A string containing the new description of the group.
	Description *string `type:"string"`

	// The name of the group.
	//
	// GroupName is a required field
	GroupName *string `min:"1" type:"string" required:"true"`

	// The new precedence value for the group. For more information about this parameter,
	// see .
	Precedence *int64 `type:"integer"`

	// The new role ARN for the group. This is used for setting the cognito:roles
	// and cognito:preferred_role claims in the token.
	RoleArn *string `min:"20" type:"string"`

	// The user pool ID for the user pool.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateGroupInput"}

	if s.GroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupName"))
	}
	if s.GroupName != nil && len(*s.GroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GroupName", 1))
	}
	if s.RoleArn != nil && len(*s.RoleArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("RoleArn", 20))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateGroupOutput struct {
	_ struct{} `type:"structure"`

	// The group object for the group.
	Group *GroupType `type:"structure"`
}

// String returns the string representation
func (s UpdateGroupOutput) String() string {
	return awsutil.Prettify(s)
}