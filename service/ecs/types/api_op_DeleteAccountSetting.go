// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ecs/enums"
)

type DeleteAccountSettingInput struct {
	_ struct{} `type:"structure"`

	// The resource name for which to disable the account setting. If serviceLongArnFormat
	// is specified, the ARN for your Amazon ECS services is affected. If taskLongArnFormat
	// is specified, the ARN and resource ID for your Amazon ECS tasks is affected.
	// If containerInstanceLongArnFormat is specified, the ARN and resource ID for
	// your Amazon ECS container instances is affected. If awsvpcTrunking is specified,
	// the ENI limit for your Amazon ECS container instances is affected.
	//
	// Name is a required field
	Name enums.SettingName `locationName:"name" type:"string" required:"true" enum:"true"`

	// The ARN of the principal, which can be an IAM user, IAM role, or the root
	// user. If you specify the root user, it disables the account setting for all
	// IAM users, IAM roles, and the root user of the account unless an IAM user
	// or role explicitly overrides these settings. If this field is omitted, the
	// setting is changed only for the authenticated user.
	PrincipalArn *string `locationName:"principalArn" type:"string"`
}

// String returns the string representation
func (s DeleteAccountSettingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAccountSettingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAccountSettingInput"}
	if len(s.Name) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteAccountSettingOutput struct {
	_ struct{} `type:"structure"`

	// The account setting for the specified principal ARN.
	Setting *Setting `locationName:"setting" type:"structure"`
}

// String returns the string representation
func (s DeleteAccountSettingOutput) String() string {
	return awsutil.Prettify(s)
}
