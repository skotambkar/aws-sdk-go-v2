// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteResourceInput struct {
	_ struct{} `type:"structure"`

	// The identifier associated with the organization from which the resource is
	// deleted.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`

	// The identifier of the resource to be deleted.
	//
	// ResourceId is a required field
	ResourceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteResourceInput"}

	if s.OrganizationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationId"))
	}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteResourceOutput) String() string {
	return awsutil.Prettify(s)
}