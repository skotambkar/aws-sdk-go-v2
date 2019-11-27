// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateHsmInput struct {
	_ struct{} `type:"structure"`

	// The Availability Zone where you are creating the HSM. To find the cluster's
	// Availability Zones, use DescribeClusters.
	//
	// AvailabilityZone is a required field
	AvailabilityZone *string `type:"string" required:"true"`

	// The identifier (ID) of the HSM's cluster. To find the cluster ID, use DescribeClusters.
	//
	// ClusterId is a required field
	ClusterId *string `type:"string" required:"true"`

	// The HSM's IP address. If you specify an IP address, use an available address
	// from the subnet that maps to the Availability Zone where you are creating
	// the HSM. If you don't specify an IP address, one is chosen for you from that
	// subnet.
	IpAddress *string `type:"string"`
}

// String returns the string representation
func (s CreateHsmInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateHsmInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateHsmInput"}

	if s.AvailabilityZone == nil {
		invalidParams.Add(aws.NewErrParamRequired("AvailabilityZone"))
	}

	if s.ClusterId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateHsmOutput struct {
	_ struct{} `type:"structure"`

	// Information about the HSM that was created.
	Hsm *Hsm `type:"structure"`
}

// String returns the string representation
func (s CreateHsmOutput) String() string {
	return awsutil.Prettify(s)
}
