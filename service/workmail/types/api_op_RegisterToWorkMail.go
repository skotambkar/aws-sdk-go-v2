// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RegisterToWorkMailInput struct {
	_ struct{} `type:"structure"`

	// The email for the user, group, or resource to be updated.
	//
	// Email is a required field
	Email *string `min:"1" type:"string" required:"true"`

	// The identifier for the user, group, or resource to be updated.
	//
	// EntityId is a required field
	EntityId *string `min:"12" type:"string" required:"true"`

	// The identifier for the organization under which the user, group, or resource
	// exists.
	//
	// OrganizationId is a required field
	OrganizationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RegisterToWorkMailInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterToWorkMailInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterToWorkMailInput"}

	if s.Email == nil {
		invalidParams.Add(aws.NewErrParamRequired("Email"))
	}
	if s.Email != nil && len(*s.Email) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Email", 1))
	}

	if s.EntityId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EntityId"))
	}
	if s.EntityId != nil && len(*s.EntityId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("EntityId", 12))
	}

	if s.OrganizationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OrganizationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RegisterToWorkMailOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RegisterToWorkMailOutput) String() string {
	return awsutil.Prettify(s)
}