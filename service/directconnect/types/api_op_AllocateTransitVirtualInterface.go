// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AllocateTransitVirtualInterfaceInput struct {
	_ struct{} `type:"structure"`

	// The ID of the connection on which the transit virtual interface is provisioned.
	//
	// ConnectionId is a required field
	ConnectionId *string `locationName:"connectionId" type:"string" required:"true"`

	// Information about the transit virtual interface.
	//
	// NewTransitVirtualInterfaceAllocation is a required field
	NewTransitVirtualInterfaceAllocation *NewTransitVirtualInterfaceAllocation `locationName:"newTransitVirtualInterfaceAllocation" type:"structure" required:"true"`

	// The ID of the AWS account that owns the transit virtual interface.
	//
	// OwnerAccount is a required field
	OwnerAccount *string `locationName:"ownerAccount" type:"string" required:"true"`
}

// String returns the string representation
func (s AllocateTransitVirtualInterfaceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AllocateTransitVirtualInterfaceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AllocateTransitVirtualInterfaceInput"}

	if s.ConnectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectionId"))
	}

	if s.NewTransitVirtualInterfaceAllocation == nil {
		invalidParams.Add(aws.NewErrParamRequired("NewTransitVirtualInterfaceAllocation"))
	}

	if s.OwnerAccount == nil {
		invalidParams.Add(aws.NewErrParamRequired("OwnerAccount"))
	}
	if s.NewTransitVirtualInterfaceAllocation != nil {
		if err := s.NewTransitVirtualInterfaceAllocation.Validate(); err != nil {
			invalidParams.AddNested("NewTransitVirtualInterfaceAllocation", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AllocateTransitVirtualInterfaceOutput struct {
	_ struct{} `type:"structure"`

	// Information about a virtual interface.
	VirtualInterface *VirtualInterface `locationName:"virtualInterface" type:"structure"`
}

// String returns the string representation
func (s AllocateTransitVirtualInterfaceOutput) String() string {
	return awsutil.Prettify(s)
}