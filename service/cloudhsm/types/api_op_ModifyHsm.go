// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the inputs for the ModifyHsm operation.
type ModifyHsmInput struct {
	_ struct{} `locationName:"ModifyHsmRequest" type:"structure"`

	// The new IP address for the elastic network interface (ENI) attached to the
	// HSM.
	//
	// If the HSM is moved to a different subnet, and an IP address is not specified,
	// an IP address will be randomly chosen from the CIDR range of the new subnet.
	EniIp *string `locationName:"EniIp" type:"string"`

	// The new external ID.
	ExternalId *string `locationName:"ExternalId" type:"string"`

	// The ARN of the HSM to modify.
	//
	// HsmArn is a required field
	HsmArn *string `locationName:"HsmArn" type:"string" required:"true"`

	// The new IAM role ARN.
	IamRoleArn *string `locationName:"IamRoleArn" type:"string"`

	// The new identifier of the subnet that the HSM is in. The new subnet must
	// be in the same Availability Zone as the current subnet.
	SubnetId *string `locationName:"SubnetId" type:"string"`

	// The new IP address for the syslog monitoring server. The AWS CloudHSM service
	// only supports one syslog monitoring server.
	SyslogIp *string `locationName:"SyslogIp" type:"string"`
}

// String returns the string representation
func (s ModifyHsmInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyHsmInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyHsmInput"}

	if s.HsmArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("HsmArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of the ModifyHsm operation.
type ModifyHsmOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the HSM.
	HsmArn *string `type:"string"`
}

// String returns the string representation
func (s ModifyHsmOutput) String() string {
	return awsutil.Prettify(s)
}