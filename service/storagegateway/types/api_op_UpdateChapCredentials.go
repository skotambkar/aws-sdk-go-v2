// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A JSON object containing one or more of the following fields:
//
//    * UpdateChapCredentialsInput$InitiatorName
//
//    * UpdateChapCredentialsInput$SecretToAuthenticateInitiator
//
//    * UpdateChapCredentialsInput$SecretToAuthenticateTarget
//
//    * UpdateChapCredentialsInput$TargetARN
type UpdateChapCredentialsInput struct {
	_ struct{} `type:"structure"`

	// The iSCSI initiator that connects to the target.
	//
	// InitiatorName is a required field
	InitiatorName *string `min:"1" type:"string" required:"true"`

	// The secret key that the initiator (for example, the Windows client) must
	// provide to participate in mutual CHAP with the target.
	//
	// The secret key must be between 12 and 16 bytes when encoded in UTF-8.
	//
	// SecretToAuthenticateInitiator is a required field
	SecretToAuthenticateInitiator *string `min:"1" type:"string" required:"true" sensitive:"true"`

	// The secret key that the target must provide to participate in mutual CHAP
	// with the initiator (e.g. Windows client).
	//
	// Byte constraints: Minimum bytes of 12. Maximum bytes of 16.
	//
	// The secret key must be between 12 and 16 bytes when encoded in UTF-8.
	SecretToAuthenticateTarget *string `min:"1" type:"string" sensitive:"true"`

	// The Amazon Resource Name (ARN) of the iSCSI volume target. Use the DescribeStorediSCSIVolumes
	// operation to return the TargetARN for specified VolumeARN.
	//
	// TargetARN is a required field
	TargetARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateChapCredentialsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateChapCredentialsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateChapCredentialsInput"}

	if s.InitiatorName == nil {
		invalidParams.Add(aws.NewErrParamRequired("InitiatorName"))
	}
	if s.InitiatorName != nil && len(*s.InitiatorName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InitiatorName", 1))
	}

	if s.SecretToAuthenticateInitiator == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecretToAuthenticateInitiator"))
	}
	if s.SecretToAuthenticateInitiator != nil && len(*s.SecretToAuthenticateInitiator) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SecretToAuthenticateInitiator", 1))
	}
	if s.SecretToAuthenticateTarget != nil && len(*s.SecretToAuthenticateTarget) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SecretToAuthenticateTarget", 1))
	}

	if s.TargetARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetARN"))
	}
	if s.TargetARN != nil && len(*s.TargetARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("TargetARN", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A JSON object containing the following fields:
type UpdateChapCredentialsOutput struct {
	_ struct{} `type:"structure"`

	// The iSCSI initiator that connects to the target. This is the same initiator
	// name specified in the request.
	InitiatorName *string `min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the target. This is the same target specified
	// in the request.
	TargetARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s UpdateChapCredentialsOutput) String() string {
	return awsutil.Prettify(s)
}