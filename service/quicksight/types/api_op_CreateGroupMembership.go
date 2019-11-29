// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateGroupMembershipInput struct {
	_ struct{} `type:"structure"`

	// The ID for the AWS account that the group is in. Currently, you use the ID
	// for the AWS account that contains your Amazon QuickSight account.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The name of the group that you want to add the user to.
	//
	// GroupName is a required field
	GroupName *string `location:"uri" locationName:"GroupName" min:"1" type:"string" required:"true"`

	// The name of the user that you want to add to the group membership.
	//
	// MemberName is a required field
	MemberName *string `location:"uri" locationName:"MemberName" min:"1" type:"string" required:"true"`

	// The namespace. Currently, you should set this to default.
	//
	// Namespace is a required field
	Namespace *string `location:"uri" locationName:"Namespace" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateGroupMembershipInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateGroupMembershipInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateGroupMembershipInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.GroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupName"))
	}
	if s.GroupName != nil && len(*s.GroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GroupName", 1))
	}

	if s.MemberName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MemberName"))
	}
	if s.MemberName != nil && len(*s.MemberName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MemberName", 1))
	}

	if s.Namespace == nil {
		invalidParams.Add(aws.NewErrParamRequired("Namespace"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateGroupMembershipOutput struct {
	_ struct{} `type:"structure"`

	// The group member.
	GroupMember *GroupMember `type:"structure"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The http status of the request.
	Status *int64 `location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s CreateGroupMembershipOutput) String() string {
	return awsutil.Prettify(s)
}
