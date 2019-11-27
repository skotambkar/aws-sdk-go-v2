// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyMountTargetSecurityGroupsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the mount target whose security groups you want to modify.
	//
	// MountTargetId is a required field
	MountTargetId *string `location:"uri" locationName:"MountTargetId" type:"string" required:"true"`

	// An array of up to five VPC security group IDs.
	SecurityGroups []string `type:"list"`
}

// String returns the string representation
func (s ModifyMountTargetSecurityGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyMountTargetSecurityGroupsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyMountTargetSecurityGroupsInput"}

	if s.MountTargetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MountTargetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyMountTargetSecurityGroupsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ModifyMountTargetSecurityGroupsOutput) String() string {
	return awsutil.Prettify(s)
}
