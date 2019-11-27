// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ec2/enums"
)

type RestoreAddressToClassicInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The Elastic IP address.
	//
	// PublicIp is a required field
	PublicIp *string `locationName:"publicIp" type:"string" required:"true"`
}

// String returns the string representation
func (s RestoreAddressToClassicInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RestoreAddressToClassicInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RestoreAddressToClassicInput"}

	if s.PublicIp == nil {
		invalidParams.Add(aws.NewErrParamRequired("PublicIp"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RestoreAddressToClassicOutput struct {
	_ struct{} `type:"structure"`

	// The Elastic IP address.
	PublicIp *string `locationName:"publicIp" type:"string"`

	// The move status for the IP address.
	Status enums.Status `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s RestoreAddressToClassicOutput) String() string {
	return awsutil.Prettify(s)
}
