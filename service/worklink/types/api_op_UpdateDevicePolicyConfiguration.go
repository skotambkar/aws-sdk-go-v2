// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateDevicePolicyConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The certificate chain, including intermediate certificates and the root certificate
	// authority certificate used to issue device certificates.
	DeviceCaCertificate *string `min:"1" type:"string"`

	// The ARN of the fleet.
	//
	// FleetArn is a required field
	FleetArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateDevicePolicyConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDevicePolicyConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDevicePolicyConfigurationInput"}
	if s.DeviceCaCertificate != nil && len(*s.DeviceCaCertificate) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DeviceCaCertificate", 1))
	}

	if s.FleetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FleetArn"))
	}
	if s.FleetArn != nil && len(*s.FleetArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("FleetArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateDevicePolicyConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateDevicePolicyConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}
