// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/snowball/enums"
)

type UpdateClusterInput struct {
	_ struct{} `type:"structure"`

	// The ID of the updated Address object.
	AddressId *string `min:"40" type:"string"`

	// The cluster ID of the cluster that you want to update, for example CID123e4567-e89b-12d3-a456-426655440000.
	//
	// ClusterId is a required field
	ClusterId *string `min:"39" type:"string" required:"true"`

	// The updated description of this cluster.
	Description *string `min:"1" type:"string"`

	// The updated ID for the forwarding address for a cluster. This field is not
	// supported in most regions.
	ForwardingAddressId *string `min:"40" type:"string"`

	// The new or updated Notification object.
	Notification *Notification `type:"structure"`

	// The updated arrays of JobResource objects that can include updated S3Resource
	// objects or LambdaResource objects.
	Resources *JobResource `type:"structure"`

	// The new role Amazon Resource Name (ARN) that you want to associate with this
	// cluster. To create a role ARN, use the CreateRole (https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html)
	// API action in AWS Identity and Access Management (IAM).
	RoleARN *string `type:"string"`

	// The updated shipping option value of this cluster's ShippingDetails object.
	ShippingOption enums.ShippingOption `type:"string" enum:"true"`
}

// String returns the string representation
func (s UpdateClusterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateClusterInput"}
	if s.AddressId != nil && len(*s.AddressId) < 40 {
		invalidParams.Add(aws.NewErrParamMinLen("AddressId", 40))
	}

	if s.ClusterId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterId"))
	}
	if s.ClusterId != nil && len(*s.ClusterId) < 39 {
		invalidParams.Add(aws.NewErrParamMinLen("ClusterId", 39))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}
	if s.ForwardingAddressId != nil && len(*s.ForwardingAddressId) < 40 {
		invalidParams.Add(aws.NewErrParamMinLen("ForwardingAddressId", 40))
	}
	if s.Resources != nil {
		if err := s.Resources.Validate(); err != nil {
			invalidParams.AddNested("Resources", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateClusterOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateClusterOutput) String() string {
	return awsutil.Prettify(s)
}