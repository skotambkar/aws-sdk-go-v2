// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type SendSSHPublicKeyInput struct {
	_ struct{} `type:"structure"`

	// The availability zone the EC2 instance was launched in.
	//
	// AvailabilityZone is a required field
	AvailabilityZone *string `min:"6" type:"string" required:"true"`

	// The EC2 instance you wish to publish the SSH key to.
	//
	// InstanceId is a required field
	InstanceId *string `min:"10" type:"string" required:"true"`

	// The OS user on the EC2 instance whom the key may be used to authenticate
	// as.
	//
	// InstanceOSUser is a required field
	InstanceOSUser *string `min:"1" type:"string" required:"true"`

	// The public key to be published to the instance. To use it after publication
	// you must have the matching private key.
	//
	// SSHPublicKey is a required field
	SSHPublicKey *string `min:"256" type:"string" required:"true"`
}

// String returns the string representation
func (s SendSSHPublicKeyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SendSSHPublicKeyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SendSSHPublicKeyInput"}

	if s.AvailabilityZone == nil {
		invalidParams.Add(aws.NewErrParamRequired("AvailabilityZone"))
	}
	if s.AvailabilityZone != nil && len(*s.AvailabilityZone) < 6 {
		invalidParams.Add(aws.NewErrParamMinLen("AvailabilityZone", 6))
	}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}
	if s.InstanceId != nil && len(*s.InstanceId) < 10 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceId", 10))
	}

	if s.InstanceOSUser == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceOSUser"))
	}
	if s.InstanceOSUser != nil && len(*s.InstanceOSUser) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InstanceOSUser", 1))
	}

	if s.SSHPublicKey == nil {
		invalidParams.Add(aws.NewErrParamRequired("SSHPublicKey"))
	}
	if s.SSHPublicKey != nil && len(*s.SSHPublicKey) < 256 {
		invalidParams.Add(aws.NewErrParamMinLen("SSHPublicKey", 256))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SendSSHPublicKeyOutput struct {
	_ struct{} `type:"structure"`

	// The request ID as logged by EC2 Connect. Please provide this when contacting
	// AWS Support.
	RequestId *string `type:"string"`

	// Indicates request success.
	Success *bool `type:"boolean"`
}

// String returns the string representation
func (s SendSSHPublicKeyOutput) String() string {
	return awsutil.Prettify(s)
}
